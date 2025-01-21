package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx, a *types.AuthData) error
}

type userHandler struct {
	db *data.Storage
}

func NewUserHandler(db *data.Storage) UserHandler {
	return &userHandler{
		db: db,
	}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	users, err := h.db.Queries.GetUsers(h.db.Ctx)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(users, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) GetUserById(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	user, err := h.db.Queries.GetUserByUsername(h.db.Ctx, a.Jwt.Claims.Username)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}
