package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ProductRoutes(api fiber.Router) {
	group := api.Group("/products")
	handler := handlers.NewProductHandler(a.db, a.validator)

	group.Get("/", handler.GetProducts)
	group.Get("/:code", handler.GetProductByCode)
}
