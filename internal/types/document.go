package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/util"
)

type DocumentWithLinesResponse struct {
	DocumentResponse
	Lines []DocumentLineResponse `json:"lines"`
}

type DocumentResponse struct {
	Agencia        string  `json:"agencia"`
	Tipodoc        string  `json:"tipodoc"`
	Documento      string  `json:"documento"`
	Tipodocv       string  `json:"tipodocv"`
	Codcliente     string  `json:"codcliente"`
	Nombrecli      string  `json:"nombrecli"`
	Contribesp     bool    `json:"contribesp"`
	RutaParme      bool    `json:"rutaParme"`
	Tipoprecio     int     `json:"tipoprecio"`
	Emision        string  `json:"emision"`
	Recepcion      string  `json:"recepcion"`
	Vence          string  `json:"vence"`
	Diascred       int     `json:"diascred"`
	Estatusdoc     int     `json:"estatusdoc"`
	Dtotneto       float64 `json:"dtotneto"`
	Dtotimpuest    float64 `json:"dtotimpuest"`
	Dtotalfinal    float64 `json:"dtotalfinal"`
	Dtotpagos      float64 `json:"dtotpagos"`
	Dtotdescuen    float64 `json:"dtotdescuen"`
	Dflete         float64 `json:"dFlete"`
	Dtotdev        float64 `json:"dtotdev"`
	Dvndmtototal   float64 `json:"dvndmtototal"`
	Dretencion     float64 `json:"dretencion"`
	Dretencioniva  float64 `json:"dretencioniva"`
	Vendedor       string  `json:"vendedor"`
	Codcoord       string  `json:"codcoord"`
	Aceptadev      bool    `json:"aceptadev"`
	KtiNegesp      bool    `json:"ktiNegesp"`
	Bsiva          float64 `json:"bsiva"`
	Bsflete        float64 `json:"bsflete"`
	Bsretencion    float64 `json:"bsretencion"`
	Bsretencioniva float64 `json:"bsretencioniva"`
	Tasadoc        float64 `json:"tasadoc"`
	Mtodcto        float64 `json:"mtodcto"`
	Fchvencedcto   string  `json:"fchvencedcto"`
	Tienedcto      bool    `json:"tienedcto"`
	Cbsret         float64 `json:"cbsret"`
	Cdret          float64 `json:"cdret"`
	Cbsretiva      float64 `json:"cbsretiva"`
	Cdretiva       float64 `json:"cdretiva"`
	Cbsrparme      float64 `json:"cbsrparme"`
	Cdrparme       float64 `json:"cdrparme"`
	Cbsretflete    float64 `json:"cbsretflete"`
	Cdretflete     float64 `json:"cdretflete"`
	Bsmtoiva       float64 `json:"bsmtoiva"`
	Bsmtofte       float64 `json:"bsmtofte"`
	RetmunMto      float64 `json:"retmunMto"`
	Dolarflete     bool    `json:"dolarflete"`
	Bsretflete     float64 `json:"bsretflete"`
	Dretflete      float64 `json:"dretflete"`
	DretmunMto     float64 `json:"dretmunMto"`
	Retivaoblig    bool    `json:"retivaoblig"`
	Edoentrega     bool    `json:"edoentrega"`
	Dmtoiva        float64 `json:"dmtoiva"`
	Prcdctoaplic   float64 `json:"prcdctoaplic"`
	Montodctodol   float64 `json:"montodctodol"`
	Montodctobs    float64 `json:"montodctobs"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
}

type DocumentLineResponse struct {
	Agencia       string  `json:"agencia"`
	Tipodoc       string  `json:"tipodoc"`
	Documento     string  `json:"documento"`
	Tipodocv      string  `json:"tipodocv"`
	Grupo         string  `json:"grupo"`
	Subgrupo      string  `json:"subgrupo"`
	Origen        int     `json:"origen"`
	Codigo        string  `json:"codigo"`
	Codhijo       string  `json:"codhijo"`
	Pid           string  `json:"pid"`
	Nombre        string  `json:"nombre"`
	Cantidad      int     `json:"cantidad"`
	Cntdevuelt    int     `json:"cntdevuelt"`
	Vndcntdevuelt int     `json:"vndcntdevuelt"`
	Dvndmtototal  float64 `json:"dvndmtototal"`
	Dpreciofin    float64 `json:"dpreciofin"`
	Dpreciounit   float64 `json:"dpreciounit"`
	Dmontoneto    float64 `json:"dmontoneto"`
	Dmontototal   float64 `json:"dmontototal"`
	Timpueprc     float64 `json:"timpueprc"`
	Unidevuelt    int     `json:"unidevuelt"`
	Fechadoc      string  `json:"fechadoc"`
	Vendedor      string  `json:"vendedor"`
	Codcoord      string  `json:"codcoord"`
}

type CreateDocumentRequest struct {
	Agencia        string                      `json:"agencia"`
	Tipodoc        string                      `json:"tipodoc"`
	Documento      string                      `json:"documento"`
	Tipodocv       string                      `json:"tipodocv"`
	Codcliente     string                      `json:"codcliente"`
	Nombrecli      string                      `json:"nombrecli"`
	Contribesp     bool                        `json:"contribesp"`
	RutaParme      bool                        `json:"rutaParme"`
	Tipoprecio     int                         `json:"tipoprecio"`
	Emision        time.Time                   `json:"emision"`
	Recepcion      time.Time                   `json:"recepcion"`
	Vence          time.Time                   `json:"vence"`
	Diascred       int                         `json:"diascred"`
	Estatusdoc     int                         `json:"estatusdoc"`
	Dtotneto       float64                     `json:"dtotneto"`
	Dtotimpuest    float64                     `json:"dtotimpuest"`
	Dtotalfinal    float64                     `json:"dtotalfinal"`
	Dtotpagos      float64                     `json:"dtotpagos"`
	Dtotdescuen    float64                     `json:"dtotdescuen"`
	Dflete         float64                     `json:"dFlete"`
	Dtotdev        float64                     `json:"dtotdev"`
	Dvndmtototal   float64                     `json:"dvndmtototal"`
	Dretencion     float64                     `json:"dretencion"`
	Dretencioniva  float64                     `json:"dretencioniva"`
	Vendedor       string                      `json:"vendedor"`
	Codcoord       string                      `json:"codcoord"`
	Aceptadev      bool                        `json:"aceptadev"`
	KtiNegesp      bool                        `json:"ktiNegesp"`
	Bsiva          float64                     `json:"bsiva"`
	Bsflete        float64                     `json:"bsflete"`
	Bsretencion    float64                     `json:"bsretencion"`
	Bsretencioniva float64                     `json:"bsretencioniva"`
	Tasadoc        float64                     `json:"tasadoc"`
	Mtodcto        float64                     `json:"mtodcto"`
	Fchvencedcto   time.Time                   `json:"fchvencedcto"`
	Tienedcto      bool                        `json:"tienedcto"`
	Cbsret         float64                     `json:"cbsret"`
	Cdret          float64                     `json:"cdret"`
	Cbsretiva      float64                     `json:"cbsretiva"`
	Cdretiva       float64                     `json:"cdretiva"`
	Cbsrparme      float64                     `json:"cbsrparme"`
	Cdrparme       float64                     `json:"cdrparme"`
	Cbsretflete    float64                     `json:"cbsretflete"`
	Cdretflete     float64                     `json:"cdretflete"`
	Bsmtoiva       float64                     `json:"bsmtoiva"`
	Bsmtofte       float64                     `json:"bsmtofte"`
	RetmunMto      float64                     `json:"retmunMto"`
	Dolarflete     bool                        `json:"dolarflete"`
	Bsretflete     float64                     `json:"bsretflete"`
	Dretflete      float64                     `json:"dretflete"`
	DretmunMto     float64                     `json:"dretmunMto"`
	Retivaoblig    bool                        `json:"retivaoblig"`
	Edoentrega     bool                        `json:"edoentrega"`
	Dmtoiva        float64                     `json:"dmtoiva"`
	Prcdctoaplic   float64                     `json:"prcdctoaplic"`
	Montodctodol   float64                     `json:"montodctodol"`
	Montodctobs    float64                     `json:"montodctobs"`
	Lines          []CreateDocumentLineRequest `json:"lines"`
}

type CreateDocumentLineRequest struct {
	Agencia       string    `json:"agencia"`
	Tipodoc       string    `json:"tipodoc"`
	Documento     string    `json:"documento"`
	Tipodocv      string    `json:"tipodocv"`
	Grupo         string    `json:"grupo"`
	Subgrupo      string    `json:"subgrupo"`
	Origen        int       `json:"origen"`
	Codigo        string    `json:"codigo"`
	Codhijo       string    `json:"codhijo"`
	Pid           string    `json:"pid"`
	Nombre        string    `json:"nombre"`
	Cantidad      int       `json:"cantidad"`
	Cntdevuelt    int       `json:"cntdevuelt"`
	Vndcntdevuelt int       `json:"vndcntdevuelt"`
	Dvndmtototal  float64   `json:"dvndmtototal"`
	Dpreciofin    float64   `json:"dpreciofin"`
	Dpreciounit   float64   `json:"dpreciounit"`
	Dmontoneto    float64   `json:"dmontoneto"`
	Dmontototal   float64   `json:"dmontototal"`
	Timpueprc     float64   `json:"timpueprc"`
	Unidevuelt    int32     `json:"unidevuelt"`
	Fechadoc      time.Time `json:"fechadoc"`
	Vendedor      string    `json:"vendedor"`
	Codcoord      string    `json:"codcoord"`
}

type UpdateDocumentRequest struct {
	Agencia        string                      `json:"agencia"`
	Tipodoc        string                      `json:"tipodoc"`
	Documento      string                      `json:"documento"`
	Tipodocv       string                      `json:"tipodocv"`
	Codcliente     string                      `json:"codcliente"`
	Nombrecli      string                      `json:"nombrecli"`
	Contribesp     bool                        `json:"contribesp"`
	RutaParme      bool                        `json:"rutaParme"`
	Tipoprecio     int                         `json:"tipoprecio"`
	Emision        time.Time                   `json:"emision"`
	Recepcion      time.Time                   `json:"recepcion"`
	Vence          time.Time                   `json:"vence"`
	Diascred       int                         `json:"diascred"`
	Estatusdoc     int                         `json:"estatusdoc"`
	Dtotneto       float64                     `json:"dtotneto"`
	Dtotimpuest    float64                     `json:"dtotimpuest"`
	Dtotalfinal    float64                     `json:"dtotalfinal"`
	Dtotpagos      float64                     `json:"dtotpagos"`
	Dtotdescuen    float64                     `json:"dtotdescuen"`
	Dflete         float64                     `json:"dFlete"`
	Dtotdev        float64                     `json:"dtotdev"`
	Dvndmtototal   float64                     `json:"dvndmtototal"`
	Dretencion     float64                     `json:"dretencion"`
	Dretencioniva  float64                     `json:"dretencioniva"`
	Vendedor       string                      `json:"vendedor"`
	Codcoord       string                      `json:"codcoord"`
	Aceptadev      bool                        `json:"aceptadev"`
	KtiNegesp      bool                        `json:"ktiNegesp"`
	Bsiva          float64                     `json:"bsiva"`
	Bsflete        float64                     `json:"bsflete"`
	Bsretencion    float64                     `json:"bsretencion"`
	Bsretencioniva float64                     `json:"bsretencioniva"`
	Tasadoc        float64                     `json:"tasadoc"`
	Mtodcto        float64                     `json:"mtodcto"`
	Fchvencedcto   time.Time                   `json:"fchvencedcto"`
	Tienedcto      bool                        `json:"tienedcto"`
	Cbsret         float64                     `json:"cbsret"`
	Cdret          float64                     `json:"cdret"`
	Cbsretiva      float64                     `json:"cbsretiva"`
	Cdretiva       float64                     `json:"cdretiva"`
	Cbsrparme      float64                     `json:"cbsrparme"`
	Cdrparme       float64                     `json:"cdrparme"`
	Cbsretflete    float64                     `json:"cbsretflete"`
	Cdretflete     float64                     `json:"cdretflete"`
	Bsmtoiva       float64                     `json:"bsmtoiva"`
	Bsmtofte       float64                     `json:"bsmtofte"`
	RetmunMto      float64                     `json:"retmunMto"`
	Dolarflete     bool                        `json:"dolarflete"`
	Bsretflete     float64                     `json:"bsretflete"`
	Dretflete      float64                     `json:"dretflete"`
	DretmunMto     float64                     `json:"dretmunMto"`
	Retivaoblig    bool                        `json:"retivaoblig"`
	Edoentrega     bool                        `json:"edoentrega"`
	Dmtoiva        float64                     `json:"dmtoiva"`
	Prcdctoaplic   float64                     `json:"prcdctoaplic"`
	Montodctodol   float64                     `json:"montodctodol"`
	Montodctobs    float64                     `json:"montodctobs"`
	Lines          []UpdateDocumentLineRequest `json:"lines"`
}

type UpdateDocumentLineRequest struct {
	Agencia       string    `json:"agencia"`
	Tipodoc       string    `json:"tipodoc"`
	Documento     string    `json:"documento"`
	Tipodocv      string    `json:"tipodocv"`
	Grupo         string    `json:"grupo"`
	Subgrupo      string    `json:"subgrupo"`
	Origen        int       `json:"origen"`
	Codigo        string    `json:"codigo"`
	Codhijo       string    `json:"codhijo"`
	Pid           string    `json:"pid"`
	Nombre        string    `json:"nombre"`
	Cantidad      int       `json:"cantidad"`
	Cntdevuelt    int       `json:"cntdevuelt"`
	Vndcntdevuelt int       `json:"vndcntdevuelt"`
	Dvndmtototal  float64   `json:"dvndmtototal"`
	Dpreciofin    float64   `json:"dpreciofin"`
	Dpreciounit   float64   `json:"dpreciounit"`
	Dmontoneto    float64   `json:"dmontoneto"`
	Dmontototal   float64   `json:"dmontototal"`
	Timpueprc     float64   `json:"timpueprc"`
	Unidevuelt    int32     `json:"unidevuelt"`
	Fechadoc      time.Time `json:"fechadoc"`
	Vendedor      string    `json:"vendedor"`
	Codcoord      string    `json:"codcoord"`
}

func DbDocToDocument(db *db.ClossDocument) *DocumentResponse {
	return mapToDocument(
		db.Agencia,
		db.Tipodoc,
		db.Documento,
		db.Tipodocv,
		db.Codcliente,
		db.Nombrecli,
		db.Contribesp == 1,
		db.RutaParme == 1,
		int(db.Tipoprecio),
		db.Emision,
		db.Recepcion,
		db.Vence,
		int(db.Diascred),
		int(db.Estatusdoc),
		db.Dtotneto,
		db.Dtotimpuest,
		db.Dtotalfinal,
		db.Dtotpagos,
		db.Dtotdescuen,
		db.Dflete,
		db.Dtotdev,
		db.Dvndmtototal,
		db.Dretencion,
		db.Dretencioniva,
		db.Vendedor,
		db.Codcoord,
		db.Aceptadev == 1,
		db.KtiNegesp == 1,
		db.Bsiva,
		db.Bsflete,
		db.Bsretencion,
		db.Bsretencioniva,
		db.Tasadoc,
		db.Mtodcto,
		db.Fchvencedcto,
		db.Tienedcto == 1,
		db.Cbsret,
		db.Cdret,
		db.Cbsretiva,
		db.Cdretiva,
		db.Cbsrparme,
		db.Cdrparme,
		db.Cbsretflete,
		db.Cdretflete,
		db.Bsmtoiva,
		db.Bsmtofte,
		db.RetmunMto,
		db.Dolarflete == 1,
		db.Bsretflete,
		db.Dretflete,
		db.DretmunMto,
		db.Retivaoblig == 1,
		db.Edoentrega == 1,
		db.Dmtoiva,
		db.Prcdctoaplic,
		db.Montodctodol,
		db.Montodctobs,
		db.CreatedAt,
		db.UpdatedAt,
	)
}

func DbDocLineToDocLine(db *db.ClossDocumentLine) *DocumentLineResponse {
	return mapToDocLine(
		db.Agencia,
		db.Tipodoc,
		db.Documento,
		db.Tipodocv,
		db.Grupo,
		db.Subgrupo,
		int(db.Origen),
		db.Codigo,
		db.Codhijo,
		db.Pid,
		db.Nombre,
		int(db.Cantidad),
		int(db.Cntdevuelt),
		int(db.Vndcntdevuelt),
		db.Dvndmtototal,
		db.Dpreciofin,
		db.Dpreciounit,
		db.Dmontoneto,
		db.Dmontototal,
		db.Timpueprc,
		int(db.Unidevuelt),
		db.Fechadoc,
		db.Vendedor,
		db.Codcoord,
	)
}

func CreateDocumentToDb(r *CreateDocumentRequest) *db.CreateDocumentParams {
	return &db.CreateDocumentParams{
		Agencia:        r.Agencia,
		Tipodoc:        r.Tipodoc,
		Documento:      r.Documento,
		Tipodocv:       r.Tipodocv,
		Codcliente:     r.Codcliente,
		Nombrecli:      r.Nombrecli,
		Contribesp:     int64(util.BoolToInt(r.Contribesp)),
		RutaParme:      int64(util.BoolToInt(r.RutaParme)),
		Tipoprecio:     int64(r.Tipoprecio),
		Emision:        r.Emision.String(),
		Recepcion:      r.Recepcion.String(),
		Vence:          r.Vence.String(),
		Diascred:       int64(r.Diascred),
		Estatusdoc:     int64(r.Estatusdoc),
		Dtotneto:       r.Dtotneto,
		Dtotimpuest:    r.Dtotimpuest,
		Dtotalfinal:    r.Dtotalfinal,
		Dtotpagos:      r.Dtotpagos,
		Dtotdescuen:    r.Dtotdescuen,
		Dflete:         r.Dflete,
		Dtotdev:        r.Dtotdev,
		Dvndmtototal:   r.Dvndmtototal,
		Dretencion:     r.Dretencion,
		Dretencioniva:  r.Dretencioniva,
		Vendedor:       r.Vendedor,
		Codcoord:       r.Codcoord,
		Aceptadev:      int64(util.BoolToInt(r.Aceptadev)),
		KtiNegesp:      int64(util.BoolToInt(r.KtiNegesp)),
		Bsiva:          r.Bsiva,
		Bsflete:        r.Bsflete,
		Bsretencion:    r.Bsretencion,
		Bsretencioniva: r.Bsretencioniva,
		Tasadoc:        r.Tasadoc,
		Mtodcto:        r.Mtodcto,
		Fchvencedcto:   r.Fchvencedcto.String(),
		Tienedcto:      int64(util.BoolToInt(r.Tienedcto)),
		Cbsret:         r.Cbsret,
		Cdret:          r.Cdret,
		Cbsretiva:      r.Cbsretiva,
		Cdretiva:       r.Cdretiva,
		Cbsrparme:      r.Cbsrparme,
		Cdrparme:       r.Cdrparme,
		Cbsretflete:    r.Cbsretflete,
		Cdretflete:     r.Cdretflete,
		Bsmtoiva:       r.Bsmtoiva,
		Bsmtofte:       r.Bsmtofte,
		RetmunMto:      r.RetmunMto,
		Dolarflete:     int64(util.BoolToInt(r.Dolarflete)),
		Bsretflete:     r.Bsretflete,
		Dretflete:      r.Dretflete,
		DretmunMto:     r.DretmunMto,
		Retivaoblig:    int64(util.BoolToInt(r.Retivaoblig)),
		Edoentrega:     int64(util.BoolToInt(r.Edoentrega)),
		Dmtoiva:        r.Dmtoiva,
		Prcdctoaplic:   r.Prcdctoaplic,
		Montodctodol:   r.Montodctodol,
		Montodctobs:    r.Montodctobs,
		CreatedAt:      time.Now().String(),
		UpdatedAt:      time.Now().String(),
	}
}

func UpdateDocumentToDb(r *UpdateDocumentRequest) *db.UpdateDocumentParams {
	return &db.UpdateDocumentParams{
		Agencia:        r.Agencia,
		Tipodoc:        r.Tipodoc,
		Documento:      r.Documento,
		Tipodocv:       r.Tipodocv,
		Codcliente:     r.Codcliente,
		Nombrecli:      r.Nombrecli,
		Contribesp:     int64(util.BoolToInt(r.Contribesp)),
		RutaParme:      int64(util.BoolToInt(r.RutaParme)),
		Tipoprecio:     int64(r.Tipoprecio),
		Emision:        r.Emision.String(),
		Recepcion:      r.Recepcion.String(),
		Vence:          r.Vence.String(),
		Diascred:       int64(r.Diascred),
		Estatusdoc:     int64(r.Estatusdoc),
		Dtotneto:       r.Dtotneto,
		Dtotimpuest:    r.Dtotimpuest,
		Dtotalfinal:    r.Dtotalfinal,
		Dtotpagos:      r.Dtotpagos,
		Dtotdescuen:    r.Dtotdescuen,
		Dflete:         r.Dflete,
		Dtotdev:        r.Dtotdev,
		Dvndmtototal:   r.Dvndmtototal,
		Dretencion:     r.Dretencion,
		Dretencioniva:  r.Dretencioniva,
		Vendedor:       r.Vendedor,
		Codcoord:       r.Codcoord,
		Aceptadev:      int64(util.BoolToInt(r.Aceptadev)),
		KtiNegesp:      int64(util.BoolToInt(r.KtiNegesp)),
		Bsiva:          r.Bsiva,
		Bsflete:        r.Bsflete,
		Bsretencion:    r.Bsretencion,
		Bsretencioniva: r.Bsretencioniva,
		Tasadoc:        r.Tasadoc,
		Mtodcto:        r.Mtodcto,
		Fchvencedcto:   r.Fchvencedcto.String(),
		Tienedcto:      int64(util.BoolToInt(r.Tienedcto)),
		Cbsret:         r.Cbsret,
		Cdret:          r.Cdret,
		Cbsretiva:      r.Cbsretiva,
		Cdretiva:       r.Cdretiva,
		Cbsrparme:      r.Cbsrparme,
		Cdrparme:       r.Cdrparme,
		Cbsretflete:    r.Cbsretflete,
		Cdretflete:     r.Cdretflete,
		Bsmtoiva:       r.Bsmtoiva,
		Bsmtofte:       r.Bsmtofte,
		RetmunMto:      r.RetmunMto,
		Dolarflete:     int64(util.BoolToInt(r.Dolarflete)),
		Bsretflete:     r.Bsretflete,
		Dretflete:      r.Dretflete,
		DretmunMto:     r.DretmunMto,
		Retivaoblig:    int64(util.BoolToInt(r.Retivaoblig)),
		Edoentrega:     int64(util.BoolToInt(r.Edoentrega)),
		Dmtoiva:        r.Dmtoiva,
		Prcdctoaplic:   r.Prcdctoaplic,
		Montodctodol:   r.Montodctodol,
		Montodctobs:    r.Montodctobs,
		UpdatedAt:      time.Now().String(),
	}
}

func CreateDocumentLineToDb(r *CreateDocumentLineRequest) *db.CreateDocumentLineParams {
	return &db.CreateDocumentLineParams{
		Agencia:       r.Agencia,
		Tipodoc:       r.Tipodoc,
		Documento:     r.Documento,
		Tipodocv:      r.Tipodocv,
		Grupo:         r.Grupo,
		Subgrupo:      r.Subgrupo,
		Origen:        int64(r.Origen),
		Codigo:        r.Codigo,
		Codhijo:       r.Codhijo,
		Pid:           r.Pid,
		Nombre:        r.Nombre,
		Cantidad:      int64(r.Cantidad),
		Cntdevuelt:    int64(r.Cntdevuelt),
		Vndcntdevuelt: int64(r.Vndcntdevuelt),
		Dvndmtototal:  r.Dvndmtototal,
		Dpreciofin:    r.Dpreciofin,
		Dpreciounit:   r.Dpreciounit,
		Dmontoneto:    r.Dmontoneto,
		Dmontototal:   r.Dmontototal,
		Timpueprc:     r.Timpueprc,
		Unidevuelt:    int64(r.Unidevuelt),
		Fechadoc:      r.Fechadoc.String(),
		Vendedor:      r.Vendedor,
		Codcoord:      r.Codcoord,
	}
}

func UpdateDocumentLineToDb(r *UpdateDocumentLineRequest) *db.UpdateDocumentLineParams {
	return &db.UpdateDocumentLineParams{
		Agencia:       r.Agencia,
		Tipodoc:       r.Tipodoc,
		Documento:     r.Documento,
		Tipodocv:      r.Tipodocv,
		Grupo:         r.Grupo,
		Subgrupo:      r.Subgrupo,
		Origen:        int64(r.Origen),
		Codigo:        r.Codigo,
		Codhijo:       r.Codhijo,
		Pid:           r.Pid,
		Nombre:        r.Nombre,
		Cantidad:      int64(r.Cantidad),
		Cntdevuelt:    int64(r.Cntdevuelt),
		Vndcntdevuelt: int64(r.Vndcntdevuelt),
		Dvndmtototal:  r.Dvndmtototal,
		Dpreciofin:    r.Dpreciofin,
		Dpreciounit:   r.Dpreciounit,
		Dmontoneto:    r.Dmontoneto,
		Dmontototal:   r.Dmontototal,
		Timpueprc:     r.Timpueprc,
		Unidevuelt:    int64(r.Unidevuelt),
		Fechadoc:      r.Fechadoc.String(),
		Vendedor:      r.Vendedor,
		Codcoord:      r.Codcoord,
	}
}

func mapToDocument(
	agencia string,
	tipodoc string,
	documento string,
	tipodocv string,
	codcliente string,
	nombrecli string,
	contribesp bool,
	rutaParme bool,
	tipoprecio int,
	emision string,
	recepcion string,
	vence string,
	diascred int,
	estatusdoc int,
	dtotneto float64,
	dtotimpuest float64,
	dtotalfinal float64,
	dtotpagos float64,
	dtotdescuen float64,
	dFlete float64,
	dtotdev float64,
	dvndmtototal float64,
	dretencion float64,
	dretencioniva float64,
	vendedor string,
	codcoord string,
	aceptadev bool,
	ktiNegesp bool,
	bsiva float64,
	bsflete float64,
	bsretencion float64,
	bsretencioniva float64,
	tasadoc float64,
	mtodcto float64,
	fchvencedcto string,
	tienedcto bool,
	cbsret float64,
	cdret float64,
	cbsretiva float64,
	cdretiva float64,
	cbsrparme float64,
	cdrparme float64,
	cbsretflete float64,
	cdretflete float64,
	bsmtoiva float64,
	bsmtofte float64,
	retmunMto float64,
	dolarflete bool,
	bsretflete float64,
	dretflete float64,
	dretmunMto float64,
	retivaoblig bool,
	edoentrega bool,
	dmtoiva float64,
	prcdctoaplic float64,
	montodctodol float64,
	montodctobs float64,
	createdAt string,
	updatedAt string,
) *DocumentResponse {
	return &DocumentResponse{
		Agencia:        agencia,
		Tipodoc:        tipodoc,
		Documento:      documento,
		Tipodocv:       tipodocv,
		Codcliente:     codcliente,
		Nombrecli:      nombrecli,
		Contribesp:     contribesp,
		RutaParme:      rutaParme,
		Tipoprecio:     tipoprecio,
		Emision:        emision,
		Recepcion:      recepcion,
		Vence:          vence,
		Diascred:       diascred,
		Estatusdoc:     estatusdoc,
		Dtotneto:       dtotneto,
		Dtotimpuest:    dtotimpuest,
		Dtotalfinal:    dtotalfinal,
		Dtotpagos:      dtotpagos,
		Dtotdescuen:    dtotdescuen,
		Dflete:         dFlete,
		Dtotdev:        dtotdev,
		Dvndmtototal:   dvndmtototal,
		Dretencion:     dretencion,
		Dretencioniva:  dretencioniva,
		Vendedor:       vendedor,
		Codcoord:       codcoord,
		Aceptadev:      aceptadev,
		KtiNegesp:      ktiNegesp,
		Bsiva:          bsiva,
		Bsflete:        bsflete,
		Bsretencion:    bsretencion,
		Bsretencioniva: bsretencioniva,
		Tasadoc:        tasadoc,
		Mtodcto:        mtodcto,
		Fchvencedcto:   fchvencedcto,
		Tienedcto:      tienedcto,
		Cbsret:         cbsret,
		Cdret:          cdret,
		Cbsretiva:      cbsretiva,
		Cdretiva:       cdretiva,
		Cbsrparme:      cbsrparme,
		Cdrparme:       cdrparme,
		Cbsretflete:    cbsretflete,
		Cdretflete:     cdretflete,
		Bsmtoiva:       bsmtoiva,
		Bsmtofte:       bsmtofte,
		RetmunMto:      retmunMto,
		Dolarflete:     dolarflete,
		Bsretflete:     bsretflete,
		Dretflete:      dretflete,
		DretmunMto:     dretmunMto,
		Retivaoblig:    retivaoblig,
		Edoentrega:     edoentrega,
		Dmtoiva:        dmtoiva,
		Prcdctoaplic:   prcdctoaplic,
		Montodctodol:   montodctodol,
		Montodctobs:    montodctobs,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

func mapToDocLine(
	agencia string,
	tipodoc string,
	documento string,
	tipodocv string,
	grupo string,
	subgrupo string,
	origen int,
	codigo string,
	codhijo string,
	pid string,
	nombre string,
	cantidad int,
	cntdevuelt int,
	vndcntdevuelt int,
	dvndmtototal float64,
	dpreciofin float64,
	dpreciounit float64,
	dmontoneto float64,
	dmontototal float64,
	timpueprc float64,
	unidevuelt int,
	fechadoc string,
	vendedor string,
	codcoord string,
) *DocumentLineResponse {
	return &DocumentLineResponse{
		Agencia:       agencia,
		Tipodoc:       tipodoc,
		Documento:     documento,
		Tipodocv:      tipodocv,
		Grupo:         grupo,
		Subgrupo:      subgrupo,
		Origen:        origen,
		Codigo:        codigo,
		Codhijo:       codhijo,
		Pid:           pid,
		Nombre:        nombre,
		Cantidad:      cantidad,
		Cntdevuelt:    cntdevuelt,
		Vndcntdevuelt: vndcntdevuelt,
		Dvndmtototal:  dvndmtototal,
		Dpreciofin:    dpreciofin,
		Dpreciounit:   dpreciounit,
		Dmontoneto:    dmontoneto,
		Dmontototal:   dmontototal,
		Timpueprc:     timpueprc,
		Unidevuelt:    unidevuelt,
		Fechadoc:      fechadoc,
		Vendedor:      vendedor,
		Codcoord:      codcoord,
	}
}

func MapToDocWithLines(head *DocumentResponse, lines []DocumentLineResponse) *DocumentWithLinesResponse {
	return &DocumentWithLinesResponse{
		DocumentResponse: *head,
		Lines:            lines,
	}
}

func GroupDocWithLinesRows(rows []db.GetDocumentsWithLinesRow) []DocumentWithLinesResponse {
	documents := make([]DocumentWithLinesResponse, 0)
	group := make(map[DocumentResponse][]DocumentLineResponse)
	doc := new(DocumentResponse)

	for _, row := range rows {
		if doc == nil {
			doc = mapGetDocumentsRowToDoc(&row)
		}

		if doc.Documento != row.Documento {
			doc = mapGetDocumentsRowToDoc(&row)
		}

		group[*doc] = append(group[*doc], *mapGetDocumentsRowToDocLine(&row))
	}

	for key, value := range group {
		documents = append(documents, *MapToDocWithLines(&key, value))
	}

	return documents
}

func GroupDocWithLinesByCodeRows(rows []db.GetDocumentWithLinesByCodeRow) *DocumentWithLinesResponse {
	document := new(DocumentWithLinesResponse)
	group := make(map[DocumentResponse][]DocumentLineResponse)
	doc := new(DocumentResponse)

	for _, row := range rows {
		if doc == nil {
			doc = mapGetDocumentsRowByCodeToDoc(&row)
		}

		if doc.Documento != row.Documento {
			doc = mapGetDocumentsRowByCodeToDoc(&row)
		}

		group[*doc] = append(group[*doc], *mapGetDocumentsRowByCodeToDocLine(&row))
	}

	for key, value := range group {
		document = MapToDocWithLines(&key, value)
	}

	return document
}

func mapGetDocumentsRowToDoc(row *db.GetDocumentsWithLinesRow) *DocumentResponse {
	return mapToDocument(
		row.Agencia,
		row.Tipodoc,
		row.Documento,
		row.Tipodocv,
		row.Codcliente,
		row.Nombrecli,
		row.Contribesp == 1,
		row.RutaParme == 1,
		int(row.Tipoprecio),
		row.Emision,
		row.Recepcion,
		row.Vence,
		int(row.Diascred),
		int(row.Estatusdoc),
		row.Dtotneto,
		row.Dtotimpuest,
		row.Dtotalfinal,
		row.Dtotpagos,
		row.Dtotdescuen,
		row.Dflete,
		row.Dtotdev,
		row.Dvndmtototal,
		row.Dretencion,
		row.Dretencioniva,
		row.Vendedor,
		row.Codcoord,
		row.Aceptadev == 1,
		row.KtiNegesp == 1,
		row.Bsiva,
		row.Bsflete,
		row.Bsretencion,
		row.Bsretencioniva,
		row.Tasadoc,
		row.Mtodcto,
		row.Fchvencedcto,
		row.Tienedcto == 1,
		row.Cbsret,
		row.Cdret,
		row.Cbsretiva,
		row.Cdretiva,
		row.Cbsrparme,
		row.Cdrparme,
		row.Cbsretflete,
		row.Cdretflete,
		row.Bsmtoiva,
		row.Bsmtofte,
		row.RetmunMto,
		row.Dolarflete == 1,
		row.Bsretflete,
		row.Dretflete,
		row.DretmunMto,
		row.Retivaoblig == 1,
		row.Edoentrega == 1,
		row.Dmtoiva,
		row.Prcdctoaplic,
		row.Montodctodol,
		row.Montodctobs,
		row.CreatedAt,
		row.UpdatedAt,
	)
}

func mapGetDocumentsRowToDocLine(row *db.GetDocumentsWithLinesRow) *DocumentLineResponse {
	return mapToDocLine(
		row.Agencia_2.String,
		row.Tipodoc_2.String,
		row.Documento_2.String,
		row.Tipodocv_2.String,
		row.Grupo.String,
		row.Subgrupo.String,
		int(row.Origen.Int64),
		row.Codigo.String,
		row.Codhijo.String,
		row.Pid.String,
		row.Nombre.String,
		int(row.Cantidad.Int64),
		int(row.Cntdevuelt.Int64),
		int(row.Vndcntdevuelt.Int64),
		row.Dvndmtototal_2.Float64,
		row.Dpreciofin.Float64,
		row.Dpreciounit.Float64,
		row.Dmontoneto.Float64,
		row.Dmontototal.Float64,
		row.Timpueprc.Float64,
		int(row.Unidevuelt.Int64),
		row.Fechadoc.String,
		row.Vendedor_2.String,
		row.Codcoord_2.String,
	)
}

func mapGetDocumentsRowByCodeToDoc(row *db.GetDocumentWithLinesByCodeRow) *DocumentResponse {
	return mapToDocument(
		row.Agencia,
		row.Tipodoc,
		row.Documento,
		row.Tipodocv,
		row.Codcliente,
		row.Nombrecli,
		row.Contribesp == 1,
		row.RutaParme == 1,
		int(row.Tipoprecio),
		row.Emision,
		row.Recepcion,
		row.Vence,
		int(row.Diascred),
		int(row.Estatusdoc),
		row.Dtotneto,
		row.Dtotimpuest,
		row.Dtotalfinal,
		row.Dtotpagos,
		row.Dtotdescuen,
		row.Dflete,
		row.Dtotdev,
		row.Dvndmtototal,
		row.Dretencion,
		row.Dretencioniva,
		row.Vendedor,
		row.Codcoord,
		row.Aceptadev == 1,
		row.KtiNegesp == 1,
		row.Bsiva,
		row.Bsflete,
		row.Bsretencion,
		row.Bsretencioniva,
		row.Tasadoc,
		row.Mtodcto,
		row.Fchvencedcto,
		row.Tienedcto == 1,
		row.Cbsret,
		row.Cdret,
		row.Cbsretiva,
		row.Cdretiva,
		row.Cbsrparme,
		row.Cdrparme,
		row.Cbsretflete,
		row.Cdretflete,
		row.Bsmtoiva,
		row.Bsmtofte,
		row.RetmunMto,
		row.Dolarflete == 1,
		row.Bsretflete,
		row.Dretflete,
		row.DretmunMto,
		row.Retivaoblig == 1,
		row.Edoentrega == 1,
		row.Dmtoiva,
		row.Prcdctoaplic,
		row.Montodctodol,
		row.Montodctobs,
		row.CreatedAt,
		row.UpdatedAt,
	)
}

func mapGetDocumentsRowByCodeToDocLine(row *db.GetDocumentWithLinesByCodeRow) *DocumentLineResponse {
	return mapToDocLine(
		row.Agencia_2.String,
		row.Tipodoc_2.String,
		row.Documento_2.String,
		row.Tipodocv_2.String,
		row.Grupo.String,
		row.Subgrupo.String,
		int(row.Origen.Int64),
		row.Codigo.String,
		row.Codhijo.String,
		row.Pid.String,
		row.Nombre.String,
		int(row.Cantidad.Int64),
		int(row.Cntdevuelt.Int64),
		int(row.Vndcntdevuelt.Int64),
		row.Dvndmtototal_2.Float64,
		row.Dpreciofin.Float64,
		row.Dpreciounit.Float64,
		row.Dmontoneto.Float64,
		row.Dmontototal.Float64,
		row.Timpueprc.Float64,
		int(row.Unidevuelt.Int64),
		row.Fechadoc.String,
		row.Vendedor_2.String,
		row.Codcoord_2.String,
	)
}
