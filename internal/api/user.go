package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) UserRoutes(api fiber.Router) {
	group := api.Group("/users")
	handler := handlers.NewUserHandler(a.db, a.validator)

	group.Get("/me", a.authenticatedHandler(handler.GetUserById))
}
