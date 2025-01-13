package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetExistingProducts(c *fiber.Ctx) error
	GetProductByCode(c *fiber.Ctx) error
	GetExistingProductByCode(c *fiber.Ctx) error
}

type productHandler struct {
	db *data.Storage
}

func NewProductHandler(db *data.Storage) ProductHandler {
	return &productHandler{
		db: db,
	}
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	products, err := h.db.Queries.GetProducts()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(products, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetExistingProducts(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	products, err := h.db.Queries.GetExistingProducts()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(products, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetProductByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	product, err := h.db.Queries.GetProductByCode(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(product, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetExistingProductByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	product, err := h.db.Queries.GetExistingProductByCode(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(product, "Success")
	return c.Status(res.Status).JSON(res)
}
