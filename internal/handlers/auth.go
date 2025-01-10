package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
	RecoverPassword(c *fiber.Ctx) error
}

type authHandler struct {
	db data.AuthStore
}

func NewAuthHandler(db data.AuthStore) AuthHandler {
	return &authHandler{
		db: db,
	}
}

// todo: set token cookie or header
func (h *authHandler) SignIn(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.SignInRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.SignIn(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) Refresh(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.RefreshRequest)

	if err := c.BodyParser(r); err == nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	token, err := h.db.Refresh(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(token, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) RecoverPassword(c *fiber.Ctx) error {
	res := types.RespondOk("RecoverPassword handler", "Success")
	return c.Status(res.Status).JSON(res)
}
