package types

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

// func DbConfigToConfig(db *db.ClossConfig) *ConfigResponse {
// 	return &ConfigResponse{
// 		CnfgIdconfig: db.CnfgIdconfig,
// 		CnfgClase:    db.CnfgClase,
// 		CnfgTipo:     db.CnfgTipo,
// 		CnfgValnum:   db.CnfgValnum,
// 		CnfgValsino:  db.CnfgValsino == 1,
// 		CnfgValtxt:   db.CnfgValtxt,
// 		CnfgLentxt:   int16(db.CnfgLentxt),
// 		CnfgValfch:   db.CnfgValfch,
// 		CnfgActiva:   db.CnfgActiva == 1,
// 		CnfgEtiq:     db.CnfgEtiq,
// 		CnfgTtip:     db.CnfgTtip,
// 		Username:     db.Username,
// 		CreatedAt:    db.CreatedAt,
// 		UpdatedAt:    db.UpdatedAt,
// 		DeletedAt:    db.DeletedAt.String,
// 	}
// }
