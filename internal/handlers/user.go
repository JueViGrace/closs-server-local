package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUserById(c *fiber.Ctx, a *types.AuthData) error
	GetUserInfoById(c *fiber.Ctx, a *types.AuthData) error
}

type userHandler struct {
	db        data.UserStore
	validator *types.XValidator
}

func NewUserHandler(db data.UserStore, v *types.XValidator) UserHandler {
	return &userHandler{
		db:        db,
		validator: v,
	}
}

func (h *userHandler) GetUserById(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	user, err := h.db.GetUser(a)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) GetUserInfoById(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	info, err := h.db.GetUserInfo(a)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(info, "Success")
	return c.Status(res.Status).JSON(res)
}
