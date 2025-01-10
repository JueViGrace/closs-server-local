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
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	SoftDeleteProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}

type productHandler struct {
	db data.ProductStore
}

func NewProductHandler(db data.ProductStore) ProductHandler {
	return &productHandler{
		db: db,
	}
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	products, err := h.db.GetProducts()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(products, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetExistingProducts(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	products, err := h.db.GetExistingProducts()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(products, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetProductByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	product, err := h.db.GetProductByCode(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(product, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetExistingProductByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	product, err := h.db.GetExistingProductByCode(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(product, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateProductRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateProduct(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateProductRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateProduct(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) SoftDeleteProduct(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	err := h.db.SoftDeleteProduct(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	err := h.db.DeleteProduct(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}
