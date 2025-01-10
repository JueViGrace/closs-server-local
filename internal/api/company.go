package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) CompanyRoutes(api fiber.Router) {
	group := api.Group("/company")
	adminGroup := group.Group("/admin")

	handler := handlers.NewCompanyHandler(a.db.CompanyStore())

	adminGroup.Get("/", handler.GetCompanies)
	adminGroup.Get("/:code", handler.GetCompanyByCode)
	group.Get("/:code", handler.GetExistingCompanyByCode)
	adminGroup.Post("/", handler.CreateCompany)
	adminGroup.Put("/", handler.UpdateCompany)
	adminGroup.Delete("/:code", handler.SoftDeleteCompany)
	adminGroup.Delete("/forever/:code", handler.DeleteCompany)
}
