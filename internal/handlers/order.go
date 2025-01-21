package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx, a *types.AuthData) error
	GetOrdersWithLines(c *fiber.Ctx, a *types.AuthData) error
	GetOrderByCode(c *fiber.Ctx, a *types.AuthData) error
	GetOrderByCodeWithLines(c *fiber.Ctx, a *types.AuthData) error
}

type orderHandler struct {
	db *data.Storage
}

func NewOrderHandler(db *data.Storage) OrderHandler {
	return &orderHandler{
		db: db,
	}
}

func (h *orderHandler) GetOrders(c *fiber.Ctx, a *types.AuthData) error {
	return nil
}

func (h *orderHandler) GetOrdersWithLines(c *fiber.Ctx, a *types.AuthData) error {
	return nil
}

func (h *orderHandler) GetOrderByCode(c *fiber.Ctx, a *types.AuthData) error {
	return nil
}

func (h *orderHandler) GetOrderByCodeWithLines(c *fiber.Ctx, a *types.AuthData) error {
	return nil
}
