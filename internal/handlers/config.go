package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ConfigHandler interface {
	GetConfigsByUser(c *fiber.Ctx) error
}

type configHandler struct {
	db *data.Storage
}

func NewConfigHandler(db *data.Storage) ConfigHandler {
	return &configHandler{
		db: db,
	}
}

func (h *configHandler) GetConfigsByUser(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	config, err := h.db.Queries.GetConfigsByUser(c.Params("username"))
	if err != nil {

		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(config, "Success")
	return c.Status(res.Status).JSON(res)
}
