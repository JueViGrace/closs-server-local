package api

import (
	"github.com/JueViGrace/clo-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func (a *api) UserRoutes(api fiber.Router) {
	group := api.Group("/users")
	handler := handlers.NewUserHandler(a.db.UserStore())

	group.Get("/", handler.GetUsers)
	group.Get("/:id", handler.GetUserById)
	group.Delete("/:id", handler.DeleteUser)
}
