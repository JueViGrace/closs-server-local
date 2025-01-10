package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler interface {
	GetCustomers(c *fiber.Ctx) error
	GetCustomerByCode(c *fiber.Ctx) error
	CreateCustomer(c *fiber.Ctx) error
	UpdateCustomer(c *fiber.Ctx) error
}

type customerHandler struct {
	db data.CustomerStore
}

func NewCustomerHandler(db data.CustomerStore) CustomerHandler {
	return &customerHandler{
		db: db,
	}
}

func (h *customerHandler) GetCustomers(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	mq := new(types.ManagerCustomerQueries)
	if err := c.QueryParser(mq); err != nil {
		res = types.RespondBadRequest(nil, "unable to parse query parameters")
		return c.Status(res.Status).JSON(res)
	}

	sq := new(types.SalesmanCustomerQueries)
	if err := c.QueryParser(sq); err != nil {
		res = types.RespondBadRequest(nil, "unable to parse query parameters")
		return c.Status(res.Status).JSON(res)
	}

	switch {
	case mq != nil:
		customers, err := h.db.GetCustomersByManager(mq.Manager)
		if err != nil {
			res = types.RespondNotFound(err.Error(), "Failed")
			return c.Status(res.Status).JSON(res)
		}
		res = types.RespondOk(customers, "Success manager")

	case sq != nil:
		customers, err := h.db.GetCustomersBySalesman(sq.Salesman)
		if err != nil {
			res = types.RespondNotFound(err.Error(), "Failed")
			return c.Status(res.Status).JSON(res)
		}
		res = types.RespondOk(customers, "Success salesman")

	default:
		customers, err := h.db.GetCustomers()
		if err != nil {
			res = types.RespondNotFound(err.Error(), "Failed")
			return c.Status(res.Status).JSON(res)
		}
		res = types.RespondOk(customers, "Success")
	}

	return c.Status(res.Status).JSON(res)
}

func (h *customerHandler) GetCustomerByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	customer, err := h.db.GetCustomerByCode(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(customer, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *customerHandler) CreateCustomer(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateCustomerRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateCustomer(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *customerHandler) UpdateCustomer(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateCustomerRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateCustomer(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}
