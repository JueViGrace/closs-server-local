package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) AuthRoutes(api fiber.Router) {
	group := api.Group("/auth")
	handler := handlers.NewAuthHandler(a.db, a.validator)

	group.Post("/signIn", handler.SignIn)
	group.Post("/refresh", a.authenticatedHandler(handler.Refresh))
	group.Post("/recover/password", handler.RecoverPassword)
}
