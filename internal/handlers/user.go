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

	dbUser, err := h.db.MyStore.Queries.GetUserByUsername(h.db.MyStore.Ctx, a.Username)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	user := types.DbUserToUser(a.UserId, &dbUser)

	res = types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) GetUserInfoById(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	dbUser, err := h.db.MyStore.Queries.GetUserByUsername(h.db.MyStore.Ctx, a.Username)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(types.PickerInfoResponse{
		Name:    dbUser.Nombre,
		Almacen: dbUser.Almacen,
	}, "Success")
	return c.Status(res.Status).JSON(res)
}
