package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/util"
)

type CustomerResponse struct {
	Codigo          string  `json:"codigo"`
	Nombre          string  `json:"nombre"`
	Direccion       string  `json:"direccion"`
	Telefonos       string  `json:"telefonos"`
	Perscont        string  `json:"perscont"`
	Vendedor        string  `json:"vendedor"`
	Contribespecial bool    `json:"contribespecial"`
	Status          int16   `json:"status"`
	Sector          string  `json:"sector"`
	Subsector       string  `json:"subcodigo"`
	Precio          int16   `json:"precio"`
	Email           string  `json:"email"`
	KneActiva       bool    `json:"kne_activa"`
	KneMtomin       float64 `json:"kne_mtomin"`
	Noemifac        bool    `json:"noemifac"`
	Noeminota       bool    `json:"noeminota"`
	Fchultvta       string  `json:"fchultvta"`
	Mtoultvta       float64 `json:"mtoultvta"`
	Prcdpagdia      float64 `json:"prcdpagdia"`
	Promdiasp       float64 `json:"promdiasp"`
	Riesgocrd       float64 `json:"riesgocrd"`
	Cantdocs        int32   `json:"cantdocs"`
	Totmtodocs      float64 `json:"totmtodocs"`
	Prommtodoc      float64 `json:"prommtodoc"`
	Diasultvta      float64 `json:"diasultvta"`
	Promdiasvta     float64 `json:"promdiasvta"`
	Limcred         float64 `json:"limcred"`
	Dolarflete      bool    `json:"dolarflete"`
	Nodolarflete    bool    `json:"nodolarflete"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type CreateCustomerRequest struct {
	Codigo          string    `json:"codigo"`
	Nombre          string    `json:"nombre"`
	Direccion       string    `json:"direccion"`
	Telefonos       string    `json:"telefonos"`
	Perscont        string    `json:"perscont"`
	Vendedor        string    `json:"vendedor"`
	Contribespecial bool      `json:"contribespecial"`
	Status          int16     `json:"status"`
	Sector          string    `json:"sector"`
	Subcodigo       string    `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       float64   `json:"kne_mtomin"`
	Noemifac        bool      `json:"noemifac"`
	Noeminota       bool      `json:"noeminota"`
	Fchultvta       time.Time `json:"fchultvta"`
	Mtoultvta       float64   `json:"mtoultvta"`
	Prcdpagdia      float64   `json:"prcdpagdia"`
	Promdiasp       float64   `json:"promdiasp"`
	Riesgocrd       float64   `json:"riesgocrd"`
	Cantdocs        int32     `json:"cantdocs"`
	Totmtodocs      float64   `json:"totmtodocs"`
	Prommtodoc      float64   `json:"prommtodoc"`
	Diasultvta      float64   `json:"diasultvta"`
	Promdiasvta     float64   `json:"promdiasvta"`
	Limcred         float64   `json:"limcred"`
	Dolarflete      bool      `json:"dolarflete"`
	Nodolarflete    bool      `json:"nodolarflete"`
}

type UpdateCustomerRequest struct {
	Codigo          string    `json:"codigo"`
	Nombre          string    `json:"nombre"`
	Direccion       string    `json:"direccion"`
	Telefonos       string    `json:"telefonos"`
	Perscont        string    `json:"perscont"`
	Vendedor        string    `json:"vendedor"`
	Contribespecial bool      `json:"contribespecial"`
	Status          int16     `json:"status"`
	Sector          string    `json:"sector"`
	Subcodigo       string    `json:"subcodigo"`
	Precio          int16     `json:"precio"`
	Email           string    `json:"email"`
	KneActiva       bool      `json:"kne_activa"`
	KneMtomin       float64   `json:"kne_mtomin"`
	Noemifac        bool      `json:"noemifac"`
	Noeminota       bool      `json:"noeminota"`
	Fchultvta       time.Time `json:"fchultvta"`
	Mtoultvta       float64   `json:"mtoultvta"`
	Prcdpagdia      float64   `json:"prcdpagdia"`
	Promdiasp       float64   `json:"promdiasp"`
	Riesgocrd       float64   `json:"riesgocrd"`
	Cantdocs        int32     `json:"cantdocs"`
	Totmtodocs      float64   `json:"totmtodocs"`
	Prommtodoc      float64   `json:"prommtodoc"`
	Diasultvta      float64   `json:"diasultvta"`
	Promdiasvta     float64   `json:"promdiasvta"`
	Limcred         float64   `json:"limcred"`
	Dolarflete      bool      `json:"dolarflete"`
	Nodolarflete    bool      `json:"nodolarflete"`
}

func DbCustomerToCustomer(db *db.ClossCustomer) *CustomerResponse {
	return &CustomerResponse{
		Codigo:          db.Codigo,
		Nombre:          db.Nombre,
		Direccion:       db.Direccion,
		Telefonos:       db.Telefonos,
		Perscont:        db.Perscont,
		Vendedor:        db.Vendedor,
		Contribespecial: db.Contribespecial == 1,
		Status:          int16(db.Status),
		Sector:          db.Sector,
		Subsector:       db.Subsector,
		Precio:          int16(db.Precio),
		Email:           db.Email,
		KneActiva:       db.KneActiva == 1,
		KneMtomin:       db.KneMtomin,
		Noemifac:        db.Noemifac == 1,
		Noeminota:       db.Noeminota == 1,
		Fchultvta:       db.Fchultvta,
		Mtoultvta:       db.Mtoultvta,
		Prcdpagdia:      db.Prcdpagdia,
		Promdiasp:       db.Promdiasp,
		Riesgocrd:       db.Riesgocrd,
		Cantdocs:        int32(db.Cantdocs),
		Totmtodocs:      db.Totmtodocs,
		Prommtodoc:      db.Prommtodoc,
		Diasultvta:      db.Diasultvta,
		Promdiasvta:     db.Promdiasvta,
		Limcred:         db.Limcred,
		Dolarflete:      db.Dolarflete == 1,
		Nodolarflete:    db.Nodolarflete == 1,
		CreatedAt:       db.CreatedAt,
		UpdatedAt:       db.UpdatedAt,
	}
}

func CreateCustomerToDb(r *CreateCustomerRequest) *db.CreateCustomerParams {
	return &db.CreateCustomerParams{
		Codigo:          r.Codigo,
		Nombre:          r.Nombre,
		Email:           r.Email,
		Direccion:       r.Direccion,
		Telefonos:       r.Telefonos,
		Perscont:        r.Perscont,
		Vendedor:        r.Vendedor,
		Contribespecial: int64(util.BoolToInt(r.Contribespecial)),
		Status:          int64(r.Status),
		Sector:          r.Sector,
		Subsector:       r.Subcodigo,
		Precio:          int64(r.Precio),
		KneActiva:       int64(util.BoolToInt(r.KneActiva)),
		KneMtomin:       r.KneMtomin,
		Noemifac:        int64(util.BoolToInt(r.KneActiva)),
		Noeminota:       int64(util.BoolToInt(r.Noeminota)),
		Fchultvta:       r.Fchultvta.String(),
		Mtoultvta:       r.Mtoultvta,
		Prcdpagdia:      r.Prcdpagdia,
		Promdiasp:       r.Promdiasp,
		Riesgocrd:       r.Riesgocrd,
		Cantdocs:        int64(r.Cantdocs),
		Totmtodocs:      r.Totmtodocs,
		Prommtodoc:      r.Prommtodoc,
		Diasultvta:      r.Diasultvta,
		Promdiasvta:     r.Promdiasvta,
		Limcred:         r.Limcred,
		Fchcrea:         r.Fchultvta.String(),
		Dolarflete:      int64(util.BoolToInt(r.Dolarflete)),
		Nodolarflete:    int64(util.BoolToInt(r.Nodolarflete)),
		CreatedAt:       time.Now().String(),
		UpdatedAt:       time.Now().String(),
	}
}

func UpdateCustomerToDb(r *UpdateCustomerRequest) *db.UpdateCustomerParams {
	return &db.UpdateCustomerParams{
		Nombre:          r.Nombre,
		Direccion:       r.Direccion,
		Telefonos:       r.Telefonos,
		Perscont:        r.Perscont,
		Vendedor:        r.Vendedor,
		Contribespecial: int64(util.BoolToInt(r.Contribespecial)),
		Status:          int64(r.Status),
		Sector:          r.Sector,
		Subsector:       r.Subcodigo,
		Precio:          int64(r.Precio),
		Email:           r.Email,
		KneActiva:       int64(util.BoolToInt(r.KneActiva)),
		KneMtomin:       r.KneMtomin,
		Noemifac:        int64(util.BoolToInt(r.KneActiva)),
		Noeminota:       int64(util.BoolToInt(r.Noeminota)),
		Fchultvta:       r.Fchultvta.String(),
		Mtoultvta:       r.Mtoultvta,
		Prcdpagdia:      r.Prcdpagdia,
		Promdiasp:       r.Promdiasp,
		Riesgocrd:       r.Riesgocrd,
		Cantdocs:        int64(r.Cantdocs),
		Totmtodocs:      r.Totmtodocs,
		Prommtodoc:      r.Prommtodoc,
		Diasultvta:      r.Diasultvta,
		Promdiasvta:     r.Promdiasvta,
		Limcred:         r.Limcred,
		Dolarflete:      int64(util.BoolToInt(r.Dolarflete)),
		Nodolarflete:    int64(util.BoolToInt(r.Nodolarflete)),
		UpdatedAt:       time.Now().String(),
		Codigo:          r.Codigo,
	}
}
