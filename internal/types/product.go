package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/util"
)

type ProductResponse struct {
	Grupo        string  `json:"grupo"`
	Subgrupo     string  `json:"subgrupo"`
	Nombre       string  `json:"nombre"`
	Codigo       string  `json:"codigo"`
	Referencia   string  `json:"referencia"`
	Marca        string  `json:"marca"`
	Unidad       string  `json:"unidad"`
	Existencia   int32   `json:"existencia"`
	Precio1      float64 `json:"precio1"`
	Precio2      float64 `json:"precio2"`
	Precio3      float64 `json:"precio3"`
	Precio4      float64 `json:"precio4"`
	Precio5      float64 `json:"precio5"`
	Precio6      float64 `json:"precio6"`
	Precio7      float64 `json:"precio7"`
	Discont      bool    `json:"discont"`
	VtaMax       int32   `json:"vta_max"`
	VtaMin       int32   `json:"vta_min"`
	Dctotope     float64 `json:"dctotope"`
	Preventa     bool    `json:"enpreventa"`
	Comprometido int32   `json:"comprometido"`
	VtaMinenx    int32   `json:"vta_minenx"`
	VtaSolofac   bool    `json:"vta_solofac"`
	VtaSolone    bool    `json:"vta_solone"`
	Codbarras    int     `json:"codbarras"`
	Detalles     string  `json:"detalles"`
	Cantbulto    int32   `json:"cantbulto"`
	CostoProm    float64 `json:"costoProm"`
	Util1        float64 `json:"util1"`
	Util2        float64 `json:"util2"`
	Util3        float64 `json:"util3"`
	Fchultcomp   string  `json:"fchultcomp"`
	Qtyultcomp   int32   `json:"qtyultcomp"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    string  `json:"-"`
}

type CreateProductRequest struct {
	Grupo        string    `json:"grupo"`
	Subgrupo     string    `json:"subgrupo"`
	Nombre       string    `json:"nombre"`
	Codigo       string    `json:"codigo"`
	Referencia   string    `json:"referencia"`
	Marca        string    `json:"marca"`
	Unidad       string    `json:"unidad"`
	Existencia   int32     `json:"existencia"`
	Precio1      float64   `json:"precio1"`
	Precio2      float64   `json:"precio2"`
	Precio3      float64   `json:"precio3"`
	Precio4      float64   `json:"precio4"`
	Precio5      float64   `json:"precio5"`
	Precio6      float64   `json:"precio6"`
	Precio7      float64   `json:"precio7"`
	Discont      bool      `json:"discont"`
	VtaMax       int32     `json:"vta_max"`
	VtaMin       int32     `json:"vta_min"`
	Dctotope     float64   `json:"dctotope"`
	Preventa     bool      `json:"enpreventa"`
	Comprometido int32     `json:"comprometido"`
	VtaMinenx    int32     `json:"vta_minenx"`
	VtaSolofac   bool      `json:"vta_solofac"`
	VtaSolone    bool      `json:"vta_solone"`
	Codbarras    int       `json:"codbarras"`
	Detalles     string    `json:"detalles"`
	Cantbulto    int32     `json:"cantbulto"`
	CostoProm    float64   `json:"costoProm"`
	Util1        float64   `json:"util1"`
	Util2        float64   `json:"util2"`
	Util3        float64   `json:"util3"`
	Fchultcomp   time.Time `json:"fchultcomp"`
	Qtyultcomp   int32     `json:"qtyultcomp"`
}

type UpdateProductRequest struct {
	Grupo        string    `json:"grupo"`
	Subgrupo     string    `json:"subgrupo"`
	Nombre       string    `json:"nombre"`
	Codigo       string    `json:"codigo"`
	Referencia   string    `json:"referencia"`
	Marca        string    `json:"marca"`
	Unidad       string    `json:"unidad"`
	Existencia   int32     `json:"existencia"`
	Precio1      float64   `json:"precio1"`
	Precio2      float64   `json:"precio2"`
	Precio3      float64   `json:"precio3"`
	Precio4      float64   `json:"precio4"`
	Precio5      float64   `json:"precio5"`
	Precio6      float64   `json:"precio6"`
	Precio7      float64   `json:"precio7"`
	Discont      bool      `json:"discont"`
	VtaMax       int32     `json:"vta_max"`
	VtaMin       int32     `json:"vta_min"`
	Dctotope     float64   `json:"dctotope"`
	Preventa     bool      `json:"enpreventa"`
	Comprometido int32     `json:"comprometido"`
	VtaMinenx    int32     `json:"vta_minenx"`
	VtaSolofac   bool      `json:"vta_solofac"`
	VtaSolone    bool      `json:"vta_solone"`
	Codbarras    int       `json:"codbarras"`
	Detalles     string    `json:"detalles"`
	Cantbulto    int32     `json:"cantbulto"`
	CostoProm    float64   `json:"costoProm"`
	Util1        float64   `json:"util1"`
	Util2        float64   `json:"util2"`
	Util3        float64   `json:"util3"`
	Fchultcomp   time.Time `json:"fchultcomp"`
	Qtyultcomp   int32     `json:"qtyultcomp"`
}

func DbProductToProduct(db *db.ClossProduct) *ProductResponse {
	return &ProductResponse{
		Grupo:        db.Grupo,
		Subgrupo:     db.Subgrupo,
		Nombre:       db.Nombre,
		Codigo:       db.Codigo,
		Referencia:   db.Referencia,
		Marca:        db.Marca,
		Unidad:       db.Unidad,
		Existencia:   int32(db.Existencia),
		Precio1:      db.Precio1,
		Precio2:      db.Precio2,
		Precio3:      db.Precio3,
		Precio4:      db.Precio4,
		Precio5:      db.Precio5,
		Precio6:      db.Precio6,
		Precio7:      db.Precio7,
		Discont:      db.Discont == 1,
		VtaMax:       int32(db.VtaMax),
		VtaMin:       int32(db.VtaMin),
		Dctotope:     db.Dctotope,
		Preventa:     db.Preventa == 1,
		Comprometido: int32(db.Comprometido),
		VtaMinenx:    int32(db.VtaMinenx),
		VtaSolofac:   db.VtaSolofac == 1,
		VtaSolone:    db.VtaSolone == 1,
		Codbarras:    int(db.Codbarras),
		Detalles:     db.Detalles,
		Cantbulto:    int32(db.Cantbulto),
		CostoProm:    db.CostoProm,
		Util1:        db.Util1,
		Util2:        db.Util2,
		Util3:        db.Util3,
		Fchultcomp:   db.Fchultcomp,
		Qtyultcomp:   int32(db.Qtyultcomp),
		CreatedAt:    db.CreatedAt,
		UpdatedAt:    db.UpdatedAt,
		DeletedAt:    db.DeletedAt.String,
	}
}

func CreateProductToDb(r *CreateProductRequest) *db.CreateProductParams {
	return &db.CreateProductParams{
		Codigo:       r.Codigo,
		Grupo:        r.Grupo,
		Subgrupo:     r.Subgrupo,
		Nombre:       r.Nombre,
		Referencia:   r.Referencia,
		Marca:        r.Marca,
		Unidad:       r.Unidad,
		Existencia:   int64(r.Existencia),
		Precio1:      r.Precio1,
		Precio2:      r.Precio2,
		Precio3:      r.Precio3,
		Precio4:      r.Precio4,
		Precio5:      r.Precio5,
		Precio6:      r.Precio6,
		Precio7:      r.Precio7,
		Discont:      int64(util.BoolToInt(r.Discont)),
		VtaMax:       int64(r.VtaMax),
		VtaMin:       int64(r.VtaMin),
		Dctotope:     r.Dctotope,
		Preventa:     int64(util.BoolToInt(r.Preventa)),
		Comprometido: int64(r.Comprometido),
		VtaMinenx:    int64(r.VtaMinenx),
		VtaSolofac:   int64(util.BoolToInt(r.VtaSolofac)),
		VtaSolone:    int64(util.BoolToInt(r.VtaSolone)),
		Codbarras:    int64(r.Codbarras),
		Detalles:     r.Detalles,
		Cantbulto:    int64(r.Cantbulto),
		CostoProm:    r.CostoProm,
		Util1:        r.Util1,
		Util2:        r.Util2,
		Util3:        r.Util3,
		Fchultcomp:   r.Fchultcomp.String(),
		Qtyultcomp:   int64(r.Qtyultcomp),
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}
}

func UpdateProductToDb(r *UpdateProductRequest) *db.UpdateProductParams {
	return &db.UpdateProductParams{
		Codigo:       r.Codigo,
		Grupo:        r.Grupo,
		Subgrupo:     r.Subgrupo,
		Nombre:       r.Nombre,
		Referencia:   r.Referencia,
		Marca:        r.Marca,
		Unidad:       r.Unidad,
		Existencia:   int64(r.Existencia),
		Precio1:      r.Precio1,
		Precio2:      r.Precio2,
		Precio3:      r.Precio3,
		Precio4:      r.Precio4,
		Precio5:      r.Precio5,
		Precio6:      r.Precio6,
		Precio7:      r.Precio7,
		Discont:      int64(util.BoolToInt(r.Discont)),
		VtaMax:       int64(r.VtaMax),
		VtaMin:       int64(r.VtaMin),
		Dctotope:     r.Dctotope,
		Preventa:     int64(util.BoolToInt(r.Preventa)),
		Comprometido: int64(r.Comprometido),
		VtaMinenx:    int64(r.VtaMinenx),
		VtaSolofac:   int64(util.BoolToInt(r.VtaSolofac)),
		VtaSolone:    int64(util.BoolToInt(r.VtaSolone)),
		Codbarras:    int64(r.Codbarras),
		Detalles:     r.Detalles,
		Cantbulto:    int64(r.Cantbulto),
		CostoProm:    r.CostoProm,
		Util1:        r.Util1,
		Util2:        r.Util2,
		Util3:        r.Util3,
		Fchultcomp:   r.Fchultcomp.String(),
		Qtyultcomp:   int64(r.Qtyultcomp),
		UpdatedAt:    time.Now().String(),
	}
}
