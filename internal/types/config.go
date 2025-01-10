package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/util"
)

type ConfigResponse struct {
	CnfgIdconfig string  `json:"cnfg_idconfig"`
	CnfgClase    string  `json:"cnfg_clase"`
	CnfgTipo     string  `json:"cnfg_tipo"`
	CnfgValnum   float64 `json:"cnfg_valnum"`
	CnfgValsino  bool    `json:"cnfg_valsino"`
	CnfgValtxt   string  `json:"cnfg_valtxt"`
	CnfgLentxt   int16   `json:"cnfg_lentxt"`
	CnfgValfch   string  `json:"cnfg_valfch"`
	CnfgActiva   bool    `json:"cnfg_activa"`
	CnfgEtiq     string  `json:"cnfg_etiq"`
	CnfgTtip     string  `json:"cnfg_ttip"`
	Username     string  `json:"username"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    string  `json:"-"`
}

type CreateConfigRequest struct {
	CnfgIdconfig string    `json:"cnfg_idconfig"`
	CnfgClase    string    `json:"cnfg_clase"`
	CnfgTipo     string    `json:"cnfg_tipo"`
	CnfgValnum   float64   `json:"cnfg_valnum"`
	CnfgValsino  bool      `json:"cnfg_valsino"`
	CnfgValtxt   string    `json:"cnfg_valtxt"`
	CnfgLentxt   int16     `json:"cnfg_lentxt"`
	CnfgValfch   time.Time `json:"cnfg_valfch"`
	CnfgActiva   bool      `json:"cnfg_activa"`
	CnfgEtiq     string    `json:"cnfg_etiq"`
	CnfgTtip     string    `json:"cnfg_ttip"`
	Username     string    `json:"username"`
}

type UpdateConfigRequest struct {
	CnfgClase    string    `json:"cnfg_clase"`
	CnfgTipo     string    `json:"cnfg_tipo"`
	CnfgValnum   float64   `json:"cnfg_valnum"`
	CnfgValsino  bool      `json:"cnfg_valsino"`
	CnfgValtxt   string    `json:"cnfg_valtxt"`
	CnfgLentxt   int16     `json:"cnfg_lentxt"`
	CnfgValfch   time.Time `json:"cnfg_valfch"`
	CnfgActiva   bool      `json:"cnfg_activa"`
	CnfgEtiq     string    `json:"cnfg_etiq"`
	CnfgTtip     string    `json:"cnfg_ttip"`
	Username     string    `json:"username"`
	CnfgIdconfig string    `json:"cnfg_idconfig"`
}

type DeleteConfigRequest struct {
	Username     string `json:"username"`
	CnfgIdconfig string `json:"cnfg_idconfig"`
}

func DbConfigToConfig(db *db.ClossConfig) *ConfigResponse {
	return &ConfigResponse{
		CnfgIdconfig: db.CnfgIdconfig,
		CnfgClase:    db.CnfgClase,
		CnfgTipo:     db.CnfgTipo,
		CnfgValnum:   db.CnfgValnum,
		CnfgValsino:  db.CnfgValsino == 1,
		CnfgValtxt:   db.CnfgValtxt,
		CnfgLentxt:   int16(db.CnfgLentxt),
		CnfgValfch:   db.CnfgValfch,
		CnfgActiva:   db.CnfgActiva == 1,
		CnfgEtiq:     db.CnfgEtiq,
		CnfgTtip:     db.CnfgTtip,
		Username:     db.Username,
		CreatedAt:    db.CreatedAt,
		UpdatedAt:    db.UpdatedAt,
		DeletedAt:    db.DeletedAt.String,
	}
}

func CreateConfigToDb(r *CreateConfigRequest) *db.CreateConfigParams {
	return &db.CreateConfigParams{
		CnfgIdconfig: r.CnfgIdconfig,
		CnfgClase:    r.CnfgClase,
		CnfgTipo:     r.CnfgTipo,
		CnfgValnum:   r.CnfgValnum,
		CnfgValsino:  int64(util.BoolToInt(r.CnfgValsino)),
		CnfgValtxt:   r.CnfgValtxt,
		CnfgLentxt:   int64(r.CnfgLentxt),
		CnfgValfch:   r.CnfgValfch.String(),
		CnfgActiva:   int64(util.BoolToInt(r.CnfgValsino)),
		CnfgEtiq:     r.CnfgEtiq,
		CnfgTtip:     r.CnfgTtip,
		Username:     r.Username,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}
}

func UpdateConfigToDb(r *UpdateConfigRequest) *db.UpdateConfigParams {
	return &db.UpdateConfigParams{
		CnfgClase:    r.CnfgClase,
		CnfgTipo:     r.CnfgTipo,
		CnfgValnum:   r.CnfgValnum,
		CnfgValsino:  int64(util.BoolToInt(r.CnfgValsino)),
		CnfgValtxt:   r.CnfgValtxt,
		CnfgLentxt:   int64(r.CnfgLentxt),
		CnfgValfch:   r.CnfgValfch.String(),
		CnfgActiva:   int64(util.BoolToInt(r.CnfgActiva)),
		CnfgEtiq:     r.CnfgEtiq,
		CnfgTtip:     r.CnfgTtip,
		UpdatedAt:    time.Now().String(),
		CnfgIdconfig: r.CnfgIdconfig,
		Username:     r.Username,
	}
}
