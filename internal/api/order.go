package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) OrderRoutes(api fiber.Router) {
	group := api.Group("/orders")
	handler := handlers.NewOrderHandler(a.db)

	group.Get("/", a.authenticatedRoute(handler.GetOrders))
	group.Get("/:code", a.authenticatedRoute(handler.GetOrderByCode))
}
