package handlers

import (
	"fmt"
	"strings"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/database"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
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

	user, err := h.db.Queries.GetUserByUsername(h.db.Ctx, r.Username)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if !util.ValidatePassword(r.Password, user.PasswordApp) {
		res = types.RespondBadRequest(nil, "invalid credentials")
		return c.Status(res.Status).JSON(res)
	}

	token, err := createTokens(&user)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.Cache.SessionStorage().CreateSession(&types.Session{
		Username:     user.Username,
		RefreshToken: token.RefreshToken,
	})
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) Refresh(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.RefreshRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	token, err := util.ValidateJWT(r.Token)
	if err != nil {
		h.db.Cache.SessionStorage().DeleteSession(r.Token)
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	claims, ok := token.Claims.(util.JWTClaims)
	if !ok {
		h.db.Cache.SessionStorage().DeleteSession(r.Token)
		res = types.RespondBadRequest(nil, "invalid request")
		return c.Status(res.Status).JSON(res)
	}

	user, err := h.db.Queries.GetUserByUsername(h.db.Ctx, claims.Username)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	newTokens, err := createTokens(&user)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.Cache.SessionStorage().UpdateSession(&types.Session{
		Username:     user.Username,
		RefreshToken: newTokens.RefreshToken,
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

func createTokens(user *database.KeWusuario) (*types.AuthResponse, error) {
	accessToken, err := util.CreateAccessToken(user.Username, user.Vendedor)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.CreateRefreshToken(user.Username, user.Vendedor)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
