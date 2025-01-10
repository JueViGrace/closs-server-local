package api

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Api interface {
	Init() error
}

type api struct {
	*fiber.App
	db data.Storage
}

func New() Api {
	return &api{
		App: fiber.New(fiber.Config{
			ServerHeader: "CloServer",
			AppName:      "CloServer",
		}),
		db: data.NewStorage(),
	}
}

func (a *api) Init() (err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	a.Use(cors.New())
	a.Use(logger.New())

	a.RegisterRoutes()

	err = a.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		a.db.Close()
		return err
	}

	return
}
