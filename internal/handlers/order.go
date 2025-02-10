package handlers

import (
	"fmt"
	"strings"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx, a *types.AuthData) error
	GetOrderByCode(c *fiber.Ctx, a *types.AuthData) error
	UpdateOrderCart(c *fiber.Ctx, a *types.AuthData) error
	UpdateOrder(c *fiber.Ctx, a *types.AuthData) error
}

type orderHandler struct {
	db        data.OrderStore
	validator *types.XValidator
}

func NewOrderHandler(db data.OrderStore, v *types.XValidator) OrderHandler {
	return &orderHandler{
		db:        db,
		validator: v,
	}
}

func (h *orderHandler) GetOrders(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	orders, err := h.db.GetOrders(a)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(orders, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *orderHandler) GetOrderByCode(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)

	order, err := h.db.GetOrderByCode(c.Params("code"), a)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(order, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *orderHandler) UpdateOrderCart(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)
	r := new(types.UpdateOrderCartRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	order, err := h.db.UpdateOrderCart(r, a)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(order, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *orderHandler) UpdateOrder(c *fiber.Ctx, a *types.AuthData) error {
	res := new(types.APIResponse)
	r := new(types.UpdateOrderRequest)

	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if errs := h.validator.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		res = types.RespondBadRequest(nil, strings.Join(errMsgs, " and "))
		return c.Status(res.Status).JSON(res)
	}

	order, err := h.db.UpdateOrder(r, a)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(order, "Success")
	return c.Status(res.Status).JSON(res)
}
