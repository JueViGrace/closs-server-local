package types

import (
	"strings"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	"github.com/JueViGrace/closs-server-local/internal/util"
)

type ProductResponse struct {
	Agencia      string  `json:"agencia"`
	Grupo        string  `json:"grupo"`
	Subgrupo     string  `json:"subgrupo"`
	Nombre       string  `json:"nombre"`
	Codigo       string  `json:"codigo"`
	Referencia   string  `json:"referencia"`
	Marca        string  `json:"marca"`
	Unidad       string  `json:"unidad"`
	Existencia   int     `json:"existencia"`
	Comprometido int     `json:"comprometido"`
	Volumen      float64 `json:"volumen"`
	Peso         float64 `json:"peso"`
	Diametro     float64 `json:"diametro"`
	Image        string  `json:"image"`
	CreatedAt    string  `json:"created_at"`
}

func DbProductToProduct(db *database.Articulo) *ProductResponse {
	return &ProductResponse{
		Agencia:    db.Agencia,
		Grupo:      strings.TrimSpace(db.Grupo),
		Subgrupo:   strings.TrimSpace(db.Subgrupo),
		Nombre:     db.Nombre,
		Codigo:     strings.TrimSpace(db.Codigo),
		Referencia: strings.TrimSpace(db.Referencia),
		Marca:      strings.TrimSpace(db.Marca),
		Unidad:     strings.TrimSpace(db.Unidad),
		Volumen:    util.StringToFloat(db.Volumen, 2),
		Peso:       util.StringToFloat(db.Peso, 2),
		Diametro:   util.StringToFloat(db.Diametro, 2),
		Image:      db.Rutafoto,
		CreatedAt:  util.FormatDateForResponse(db.Fechacrea),
	}
}
