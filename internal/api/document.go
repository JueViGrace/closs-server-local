package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) DocumentRoutes(api fiber.Router) {
	group := api.Group("/documents")
	handler := handlers.NewDocumentHandler(a.db.DocumentStore())

	group.Get("/", handler.GetDocuments)
}
