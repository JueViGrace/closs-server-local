package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) ConfigRoutes(api fiber.Router) {
	group := api.Group("/config")
	adminGroup := group.Group("/config")

	handler := handlers.NewConfigHandler(a.db.ConfigStore())

	adminGroup.Get("/", handler.GetConfigs)
	group.Get("/:id", handler.GetConfigsByUser)
	adminGroup.Post("/", handler.CreateConfig)
	adminGroup.Put("/", handler.UpdateConfig)
	adminGroup.Delete("/", handler.DeleteConfig)
}
