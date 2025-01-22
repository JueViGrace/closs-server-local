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
	db        *data.Storage
	validator *types.XValidator
}

func NewUserHandler(db *data.Storage, v *types.XValidator) UserHandler {
	return &userHandler{
		db:        db,
		validator: v,
	}
}

func (h *userHandler) GetUserById(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	dbUser, err := h.db.Queries.GetUserByUsername(h.db.Ctx, a.Username)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	user := types.DbUserToUser(a.UserId, &dbUser)

	res = types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}
