package api

import (
	"github.com/JueViGrace/closs-server-local/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ConfigRoutes(api fiber.Router) {
	group := api.Group("/config")
	handler := handlers.NewConfigHandler(a.db, a.validator)

	group.Get("/:id", a.authenticatedRoute(handler.GetConfigsByUser))
}
