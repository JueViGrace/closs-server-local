package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) CustomerRoutes(api fiber.Router) {
	group := api.Group("/customers")
	adminGroup := group.Group("/admin")

	handler := handlers.NewCustomerHandler(a.db.CustomerStore())

	group.Get("/", handler.GetCustomers)
	group.Get("/:code", handler.GetCustomerByCode)
	adminGroup.Post("/", handler.CreateCustomer)
	adminGroup.Put("/", handler.UpdateCustomer)
}

// todo: Middlewares
