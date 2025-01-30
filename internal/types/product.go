package types

import (
	"strings"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	"github.com/JueViGrace/closs-server-local/internal/util"
)

type ProductResponse struct {
	Nombre     string `json:"nombre"`
	Codigo     string `json:"codigo"`
	Referencia string `json:"referencia"`
	Marca      string `json:"marca"`
	Unidad     string `json:"unidad"`
	Image      string `json:"image"`
	CreatedAt  string `json:"created_at"`
}

func DbProductToProduct(db *database.Articulo) *ProductResponse {
	return &ProductResponse{
		Nombre:     db.Nombre,
		Codigo:     strings.TrimSpace(db.Codigo),
		Referencia: strings.TrimSpace(db.Referencia),
		Marca:      strings.TrimSpace(db.Marca),
		Unidad:     strings.TrimSpace(db.Unidad),
		Image:      db.Rutafoto,
		CreatedAt:  util.FormatDateForResponse(db.Fechacrea),
	}
}
