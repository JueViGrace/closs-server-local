package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) AuthRoutes(api fiber.Router) {
	group := api.Group("/auth")
	handler := handlers.NewAuthHandler(a.db)

	group.Post("/signIn", handler.SignIn)
	group.Post("/recover/password", handler.RecoverPassword)
	group.Post("/refresh", handler.Refresh)
}
