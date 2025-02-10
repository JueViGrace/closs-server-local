package handlers

import (
	"fmt"
	"os"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetProductByCode(c *fiber.Ctx) error
	GetProductImageByCode(c *fiber.Ctx) error
}

type productHandler struct {
	db        data.ProductStore
	validator *types.XValidator
}

func NewProductHandler(db data.ProductStore, v *types.XValidator) ProductHandler {
	return &productHandler{
		db:        db,
		validator: v,
	}
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	products, err := h.db.GetProducts()
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(products, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetProductByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	product, err := h.db.GetProduct(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(product, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetProductImageByCode(c *fiber.Ctx) error {
	file := fmt.Sprintf("%s%s.jpg", os.Getenv("IMGS_DIR"), c.Params("code"))

	return c.Status(fiber.StatusOK).SendFile(file, true)
}
