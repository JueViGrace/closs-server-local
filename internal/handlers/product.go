package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetProductByCode(c *fiber.Ctx) error
}

type productHandler struct {
	db        *data.Storage
	validator *types.XValidator
}

func NewProductHandler(db *data.Storage, v *types.XValidator) ProductHandler {
	return &productHandler{
		db:        db,
		validator: v,
	}
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	products := make([]types.ProductResponse, 0)

	dbProducts, err := h.db.Queries.GetProducts(h.db.Ctx)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	for _, p := range dbProducts {
		products = append(products, *types.DbProductToProduct(&p))
	}

	res = types.RespondOk(products, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *productHandler) GetProductByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	product, err := h.db.Queries.GetProductByCode(h.db.Ctx, c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(product, "Success")
	return c.Status(res.Status).JSON(res)
}
