package api

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (a *api) RegisterRoutes() {
	a.ApiRoutes()
}

func (a *api) ApiRoutes() {
	api := a.App.Group("/api")

	api.Get("/health", a.HealthRoute)
	api.Get("/metrics", monitor.New(monitor.Config{
		Refresh: time.Duration(time.Second),
	}))

	a.AuthRoutes(api)
	a.ConfigRoutes(api)
	a.OrderRoutes(api)
	a.ProductRoutes(api)
	a.UserRoutes(api)
}

func (a *api) HealthRoute(c *fiber.Ctx) error {
	res := types.RespondOk(a.db.Health(), "Success")
	return c.Status(res.Status).JSON(res)
}
