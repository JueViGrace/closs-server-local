package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ProductRoutes(api fiber.Router) {
	group := api.Group("/products")
	handler := handlers.NewProductHandler(a.db.ProductStore())

	group.Get("/", handler.GetProducts)
	group.Get("/:code", handler.GetProductByCode)
	group.Post("/", handler.CreateProduct)
	group.Patch("/", handler.UpdateProduct)
	group.Delete("/:code", handler.DeleteProduct)
}
