package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) OrderRoutes(api fiber.Router) {
	group := api.Group("/orders")
	handler := handlers.NewOrderHandler(a.db)

	group.Get("/", handler.GetOrders)
	group.Get("/:code", handler.GetOrderByCode)
}
