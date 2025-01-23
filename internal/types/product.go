package types

import (
	"strings"
	"time"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
)

type ProductResponse struct {
	Agencia      string    `json:"agencia"`
	Grupo        string    `json:"grupo"`
	Subgrupo     string    `json:"subgrupo"`
	Nombre       string    `json:"nombre"`
	Codigo       string    `json:"codigo"`
	Referencia   string    `json:"referencia"`
	Marca        string    `json:"marca"`
	Unidad       string    `json:"unidad"`
	Existencia   string    `json:"existencia"`
	Comprometido string    `json:"comprometido"`
	Precio1      string    `json:"precio1"`
	Precio2      string    `json:"precio2"`
	Precio3      string    `json:"precio3"`
	Precio4      string    `json:"precio4"`
	Precio5      string    `json:"precio5"`
	Precio6      string    `json:"precio6"`
	Precio7      string    `json:"precio7"`
	Discont      bool      `json:"discont"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func DbProductToProduct(db *database.Articulo) *ProductResponse {
	return &ProductResponse{
		Agencia:      db.Agencia,
		Grupo:        strings.TrimSpace(db.Grupo),
		Subgrupo:     strings.TrimSpace(db.Subgrupo),
		Nombre:       db.Nombre,
		Codigo:       strings.TrimSpace(db.Codigo),
		Referencia:   strings.TrimSpace(db.Referencia),
		Marca:        strings.TrimSpace(db.Marca),
		Unidad:       strings.TrimSpace(db.Unidad),
		Existencia:   db.Existencia,
		Comprometido: db.Comprometido,
		Precio1:      db.Precio1,
		Precio2:      db.Precio2,
		Precio3:      db.Precio3,
		Precio4:      db.Precio4,
		Precio5:      db.Precio5,
		Precio6:      db.Precio6,
		Precio7:      db.Precio7,
		Discont:      db.Discont,
		CreatedAt:    db.Fechacrea,
		UpdatedAt:    db.Fechamodifi,
	}
}
