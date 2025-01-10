package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
)

type CompanyResponse struct {
	KedCodigo string `json:"ked_codigo"`
	KedNombre string `json:"ked_nombre"`
	KedStatus int    `json:"ked_status"`
	KedEnlace string `json:"ked_enlace"`
	KedAgen   string `json:"ked_agen"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type CreateCompanyRequest struct {
	KedCodigo string `json:"ked_codigo"`
	KedNombre string `json:"ked_nombre"`
	KedStatus int    `json:"ked_status"`
	KedEnlace string `json:"ked_enlace"`
	KedAgen   string `json:"ked_agen"`
}

type UpdateCompanyRequest struct {
	KedNombre string `json:"ked_nombre"`
	KedStatus int    `json:"ked_status"`
	KedEnlace string `json:"ked_enlace"`
	KedAgen   string `json:"ked_agen"`
	KedCodigo string `json:"ked_codigo"`
}

func DbComanyToCompany(db *db.ClossCompany) *CompanyResponse {
	return &CompanyResponse{
		KedCodigo: db.KedCodigo,
		KedNombre: db.KedNombre,
		KedStatus: int(db.KedStatus),
		KedEnlace: db.KedEnlace,
		KedAgen:   db.KedAgen,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		DeletedAt: db.DeletedAt.String,
	}
}

func CreateCompanyToDb(r *CreateCompanyRequest) *db.CreateCompanyParams {
	return &db.CreateCompanyParams{
		KedCodigo: r.KedCodigo,
		KedNombre: r.KedNombre,
		KedStatus: int64(r.KedStatus),
		KedEnlace: r.KedEnlace,
		KedAgen:   r.KedAgen,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
}

func UpdateCompanyToDb(r *UpdateCompanyRequest) *db.UpdateCompanyParams {
	return &db.UpdateCompanyParams{
		KedNombre: r.KedNombre,
		KedStatus: int64(r.KedStatus),
		KedEnlace: r.KedEnlace,
		KedAgen:   r.KedAgen,
		UpdatedAt: time.Now().String(),
		KedCodigo: r.KedCodigo,
	}
}
