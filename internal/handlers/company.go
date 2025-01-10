package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type CompanyHandler interface {
	GetCompanies(c *fiber.Ctx) error
	GetCompanyByCode(c *fiber.Ctx) error
	GetExistingCompanyByCode(c *fiber.Ctx) error
	CreateCompany(c *fiber.Ctx) error
	UpdateCompany(c *fiber.Ctx) error
	SoftDeleteCompany(c *fiber.Ctx) error
	DeleteCompany(c *fiber.Ctx) error
}

type companyHandler struct {
	db data.CompanyStore
}

func NewCompanyHandler(db data.CompanyStore) CompanyHandler {
	return &companyHandler{
		db: db,
	}
}

func (h *companyHandler) GetCompanies(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	companies, err := h.db.GetCompanies()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(companies, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) GetCompanyByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	company, err := h.db.GetCompanyByCode(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(company, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) GetExistingCompanyByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	company, err := h.db.GetExistingCompanyByCode(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(company, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) CreateCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateCompanyRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateCompany(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) UpdateCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateCompanyRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(err.Error(), "Invalid request")
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateCompany(r)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) SoftDeleteCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	err := h.db.SoftDeleteCompany(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *companyHandler) DeleteCompany(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	err := h.db.DeleteCompany(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondNoContent("Deleted", "Success")
	return c.Status(res.Status).JSON(res)
}
