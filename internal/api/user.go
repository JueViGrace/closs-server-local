package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) UserRoutes(api fiber.Router) {
	group := api.Group("/users")
	handler := handlers.NewUserHandler(a.db.UserStore(), a.validator)

	group.Get("/", handler.GetUsers)
	group.Get("/me", a.authenticatedHandler(handler.GetUserById))
	group.Get("/me/info", a.authenticatedHandler(handler.GetUserInfoById))
}
