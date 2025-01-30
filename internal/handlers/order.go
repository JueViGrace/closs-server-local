package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
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

	rows, err := h.db.MyStore.Queries.GetOrdersByUser(h.db.MyStore.Ctx, a.Username)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	orders, err = types.GroupOrdersByUserRow(rows)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(orders, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *orderHandler) GetOrderByCode(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)
	order := new(types.OrderWithLinesResponse)

	rows, err := h.db.MyStore.Queries.GetOrderByCode(h.db.MyStore.Ctx, database.GetOrderByCodeParams{
		Upickup:   a.Username,
		Documento: c.Params("code"),
	})
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	order, err = types.GroupOrderByCodeRow(rows)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(order, "Success")
	return c.Status(res.Status).JSON(res)
}
