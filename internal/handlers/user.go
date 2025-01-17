package handlers

import (
	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUserById(c *fiber.Ctx) error
}

type userHandler struct {
	db *data.Storage
}

func NewUserHandler(db *data.Storage) UserHandler {
	return &userHandler{
		db: db,
	}
}

func (h *userHandler) GetUserById(c *fiber.Ctx) error {
	// res := new(types.APIResponse)
	// id, err := util.GetIdFromParams(c.Params("id"))
	// if err != nil {
	// 	res = types.RespondBadRequest(err.Error(), "Invalid request")
	// 	return c.Status(res.Status).JSON(res)
	// }
	//
	// user, err := h.db.Queries.GetUserById(id)
	// if err != nil {
	// 	res = types.RespondNotFound(err.Error(), "Failed")
	// 	return c.Status(res.Status).JSON(res)
	// }
	//
	// res = types.RespondOk(user, "Success")
	// return c.Status(res.Status).JSON(res)
	return nil
}
