package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
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

func (h *userHandler) GetUserById(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	// user, err := h.db.Queries.GetUserById(h.db.Ctx, a.Jwt.Claims.UserId.String())
	// if err != nil {
	// 	res = types.RespondNotFound(err.Error(), "Failed")
	// 	return c.Status(res.Status).JSON(res)
	// }

	res = types.RespondOk(types.UserResponse{}, "Success")
	return c.Status(res.Status).JSON(res)
}
