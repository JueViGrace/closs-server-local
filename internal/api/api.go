package api

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Api interface {
	Init() error
}

type api struct {
	*fiber.App
	db        *data.Storage
	validator *types.XValidator
}

func New() Api {
	validator := &types.XValidator{
		Validator: validator.New(),
	}
	return &api{
		App: fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				res := types.RespondBadRequest(nil, err.Error())
				return c.Status(res.Status).JSON(res)
			},
			ServerHeader: "CloServer",
			AppName:      "CloServer",
		}),
		db:        data.NewStorage(),
		validator: validator,
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
