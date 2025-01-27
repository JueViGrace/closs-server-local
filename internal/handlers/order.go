package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx, a *types.AuthData) error
	GetOrderByCode(c *fiber.Ctx, a *types.AuthData) error
}

type orderHandler struct {
	db        *data.Storage
	validator *types.XValidator
}

func NewOrderHandler(db *data.Storage, v *types.XValidator) OrderHandler {
	return &orderHandler{
		db:        db,
		validator: v,
	}
}

func (h *orderHandler) GetOrders(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)
	orders := make([]types.OrderWithLinesResponse, 0)

	dbOrders, err := h.db.MyStore.Queries.GetOrdersByUser(h.db.MyStore.Ctx, a.Username)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	orders = types.GroupOrderByUserRow(dbOrders)

	res = types.RespondOk(orders, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *orderHandler) GetOrderByCode(c *fiber.Ctx, a *types.AuthData) error {
	return nil
}
