package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx) error
	GetOrdersWithLines(c *fiber.Ctx) error
	GetOrderByCode(c *fiber.Ctx) error
	GetOrderByCodeWithLines(c *fiber.Ctx) error
}

type orderHandler struct {
	db *data.Storage
}

func NewOrderHandler(db *data.Storage) OrderHandler {
	return &orderHandler{
		db: db,
	}
}

func (h *orderHandler) GetOrders(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetOrdersWithLines(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetOrderByCode(c *fiber.Ctx) error {
	return nil
}

func (h *orderHandler) GetOrderByCodeWithLines(c *fiber.Ctx) error {
	return nil
}
