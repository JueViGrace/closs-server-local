package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ConfigHandler interface {
	GetConfigs(c *fiber.Ctx) error
	GetConfigsByUser(c *fiber.Ctx) error
	CreateConfig(c *fiber.Ctx) error
	UpdateConfig(c *fiber.Ctx) error
	DeleteConfig(c *fiber.Ctx) error
}

type configHandler struct {
	db *data.Storage
}

func NewConfigHandler(db *data.Storage) ConfigHandler {
	return &configHandler{
		db: db,
	}
}

func (h *configHandler) GetConfigs(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	configs, err := h.db.GetConfigs()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(configs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) GetConfigsByUser(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	config, err := h.db.GetConfigsByUser(c.Params("username"))
	if err != nil {

		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(config, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) CreateConfig(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateConfigRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateConfig(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) UpdateConfig(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateConfigRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateConfig(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *configHandler) DeleteConfig(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.DeleteConfigRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	err := h.db.DeleteConfig(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}
