package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ConfigHandler interface {
	GetConfigsByUser(c *fiber.Ctx, a *types.AuthData) error
}

type configHandler struct {
	db        *data.Storage
	validator *types.XValidator
}

func NewConfigHandler(db *data.Storage, v *types.XValidator) ConfigHandler {
	return &configHandler{
		db:        db,
		validator: v,
	}
}

func (h *configHandler) GetConfigsByUser(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	// config, err := h.db.Queries.GetConfigsByUser(h.db.Ctx, a.Jwt.Claims.Username)
	// if err != nil {
	//
	// 	res = types.RespondNotFound(err.Error(), "Failed")
	// 	return c.Status(res.Status).JSON(res)
	// }

	res = types.RespondOk(types.ConfigResponse{}, "Success")
	return c.Status(res.Status).JSON(res)
}
