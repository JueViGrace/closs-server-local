package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) OrderRoutes(api fiber.Router) {
	group := api.Group("/orders")
	handler := handlers.NewOrderHandler(a.db.OrderStorage(), a.validator)

	group.Get("/", a.authenticatedHandler(handler.GetOrders))
	group.Get("/:code", a.authenticatedHandler(handler.GetOrderByCode))
}
