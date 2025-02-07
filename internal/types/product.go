package types

import (
	"strconv"
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
	Existencia int    `json:"existencia"`
	Image      string `json:"image"`
	CreatedAt  string `json:"created_at"`
}

func DbProductToProduct(db *database.Articulo) *ProductResponse {
	existencia, err := strconv.Atoi(db.Existencia)
	if err != nil {
		existencia = 0
	}

	return &ProductResponse{
		Nombre:     db.Nombre,
		Codigo:     strings.TrimSpace(db.Codigo),
		Referencia: strings.TrimSpace(db.Referencia),
		Marca:      strings.TrimSpace(db.Marca),
		Unidad:     strings.TrimSpace(db.Unidad),
		Existencia: existencia,
		Image:      db.Rutafoto,
		CreatedAt:  util.FormatDateForResponse(db.Fechacrea),
	}
}
