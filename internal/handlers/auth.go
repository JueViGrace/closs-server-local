package handlers

import (
	"fmt"
	"strings"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx, a *types.AuthData) error
	RecoverPassword(c *fiber.Ctx) error
}

type authHandler struct {
	db        *data.Storage
	validator *types.XValidator
}

func NewAuthHandler(db *data.Storage, v *types.XValidator) AuthHandler {
	return &authHandler{
		db:        db,
		validator: v,
	}
}

func (h *authHandler) SignIn(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.SignInRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	dbSession, err := h.db.CacheStore.SessionStorage().GetSessionByUsername(r.Username)
	if err == nil {
		err = h.db.CacheStore.SessionStorage().DeleteSession(dbSession.UserId)
		if err != nil {
			res = types.RespondBadRequest(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}
	}

	dbUser, err := h.db.MyStore.Queries.GetUserByUsername(h.db.MyStore.Ctx, r.Username)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	id, err := uuid.NewV7()
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	user := types.DbUserToUser(id, &dbUser)

	if !util.ValidatePassword(r.Password, user.Password) {
		res = types.RespondBadRequest(nil, "invalid credentials")
		return c.Status(res.Status).JSON(res)
	}

	token, err := createTokens(user)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.CacheStore.SessionStorage().CreateSession(&types.Session{
		UserId:       user.ID,
		Username:     user.Username,
		RefreshToken: token.RefreshToken,
		AccessToken:  token.AccessToken,
	})
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) Refresh(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)
	r := new(types.RefreshRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	dbUser, err := h.db.MyStore.Queries.GetUserByUsername(h.db.MyStore.Ctx, a.Username)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	user := types.DbUserToUser(a.UserId, &dbUser)

	newTokens, err := createTokens(user)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.CacheStore.SessionStorage().UpdateSession(&types.Session{
		UserId:       user.ID,
		Username:     dbUser.Username,
		RefreshToken: newTokens.RefreshToken,
		AccessToken:  newTokens.AccessToken,
	})
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(newTokens, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) RecoverPassword(c *fiber.Ctx) error {
	res := types.RespondOk("RecoverPassword handler", "Success")
	return c.Status(res.Status).JSON(res)
}

func createTokens(user *types.UserResponse) (*types.AuthResponse, error) {
	accessToken, err := util.CreateAccessToken(user.ID.String(), user.Username, user.Code)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.CreateRefreshToken(user.ID.String(), user.Username, user.Code)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
