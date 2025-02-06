package types

import (
	"fmt"
	"os"
	"strings"
	"time"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	"github.com/JueViGrace/closs-server-local/internal/util"
)

type OrderWithLinesResponse struct {
	OrderResponse
	Lines []OrderLineResponse `json:"lines"`
}

type OrderResponse struct {
	Agencia     string `json:"agencia"`
	TipoDoc     string `json:"tipodoc"`
	Documento   string `json:"documento"`
	CodCliente  string `json:"codcliente"`
	NombreCli   string `json:"nombrecli"`
	Emision     string `json:"emision"`
	Upickup     string `json:"upickup"`
	IdCarrito   string `json:"idcarrito"`
	Almacen     string `json:"almacen"`
	RutaCodigo  string `json:"ruta_codigo"`
	RutaDescrip string `json:"ruta_descrip"`
	KePedStatus string `json:"ke_pedstatus"`
}

type OrderLineResponse struct {
	Agencia   string          `json:"agencia"`
	TipoDoc   string          `json:"tipodoc"`
	Documento string          `json:"documento"`
	Almacen   string          `json:"almacen"`
	Product   ProductResponse `json:"product"`
	CantRef   int             `json:"cantref"`
	Cantidad  int             `json:"cantidad"`
}

func mapToOrder(
	agencia string,
	tipodoc string,
	documento string,
	codcliente string,
	nombrecli string,
	emision time.Time,
	upickup string,
	idcarrito string,
	almacen string,
	rutaCodigo string,
	rutaDescrip string,
	kePedstatus string,
) *OrderResponse {
	return &OrderResponse{
		Agencia:     agencia,
		TipoDoc:     tipodoc,
		Documento:   documento,
		CodCliente:  codcliente,
		NombreCli:   nombrecli,
		Emision:     util.FormatDateForResponse(emision),
		Upickup:     upickup,
		IdCarrito:   strings.TrimSpace(idcarrito),
		Almacen:     almacen,
		RutaCodigo:  rutaCodigo,
		RutaDescrip: strings.TrimSpace(rutaDescrip),
		KePedStatus: kePedstatus,
	}
}

func mapToOrderLine(
	agencia string,
	tipodoc string,
	documento string,
	almacen string,
	product *ProductResponse,
	cantref int,
	cantidad int,
) *OrderLineResponse {
	return &OrderLineResponse{
		Agencia:   agencia,
		TipoDoc:   tipodoc,
		Documento: documento,
		Almacen:   almacen,
		Product:   *product,
		CantRef:   cantref,
		Cantidad:  cantidad,
	}
}

func MapToOrderWithLines(head *OrderResponse, lines []OrderLineResponse) *OrderWithLinesResponse {
	return &OrderWithLinesResponse{
		OrderResponse: *head,
		Lines:         lines,
	}
}

func GroupOrdersByUserRow(rows []database.GetOrdersByUserRow) ([]OrderWithLinesResponse, error) {
	resp := make([]OrderWithLinesResponse, 0)
	group := make(map[OrderResponse][]OrderLineResponse)
	ord := new(OrderResponse)

	for _, row := range rows {
		if ord == nil {
			ord = mapGetOrdersByUserRowToOrder(&row)
		}

		if ord.Documento != row.Documento {
			ord = mapGetOrdersByUserRowToOrder(&row)
		}

		line := mapGetOrdersByUserRowToOrderLine(&row)

		group[*ord] = append(group[*ord], *line)
	}

	for key, value := range group {
		resp = append(resp, *MapToOrderWithLines(&key, value))
	}

	return resp, nil
}

func mapGetOrdersByUserRowToOrder(row *database.GetOrdersByUserRow) *OrderResponse {
	return mapToOrder(
		row.Agencia,
		row.Tipodoc,
		row.Documento,
		row.Codcliente,
		row.Nombrecli,
		row.Emision,
		row.Upickup,
		row.Idcarrito,
		row.Almacen,
		row.RutaCodigo.String,
		row.RutaDescrip.String,
		row.KePedstatus,
	)
}

func mapGetOrdersByUserRowToOrderLine(row *database.GetOrdersByUserRow) *OrderLineResponse {
	cantRef := util.StringToFloat(row.Cantref.String, 0)
	cantidad := util.StringToFloat(row.Cantidad.String, 0)

	product := &ProductResponse{
		Nombre:     strings.TrimSpace(row.Nombre.String),
		Codigo:     strings.TrimSpace(row.Codigo.String),
		Referencia: strings.TrimSpace(row.Referencia.String),
		Marca:      strings.TrimSpace(row.Marca.String),
		Unidad:     strings.TrimSpace(row.Unidad.String),
		Image:      fmt.Sprintf("%s/api/products/%s/image", os.Getenv("BASE_URL"), row.Codigo.String),
		CreatedAt:  util.FormatDateForResponse(row.Fechacrea.Time),
	}

	return mapToOrderLine(
		row.Agencia_2.String,
		row.Tipodoc_2.String,
		row.Documento_2.String,
		row.Almacen_2.String,
		product,
		int(cantRef),
		int(cantidad),
	)
}

func GroupOrderByCodeRow(rows []database.GetOrderByCodeRow) (*OrderWithLinesResponse, error) {
	resp := new(OrderWithLinesResponse)
	group := make(map[OrderResponse][]OrderLineResponse)
	ord := new(OrderResponse)

	for _, row := range rows {
		if ord == nil {
			ord = mapGetOrderByCodeRowToOrder(&row)
		}

		if ord.Documento != row.Documento {
			ord = mapGetOrderByCodeRowToOrder(&row)
		}

		line := mapGetOrderByCodeRowToOrderLine(&row)

		group[*ord] = append(group[*ord], *line)
	}

	for key, value := range group {
		resp = MapToOrderWithLines(&key, value)
	}

	return resp, nil
}

func mapGetOrderByCodeRowToOrder(row *database.GetOrderByCodeRow) *OrderResponse {
	return mapToOrder(
		row.Agencia,
		row.Tipodoc,
		row.Documento,
		row.Codcliente,
		row.Nombrecli,
		row.Emision,
		row.Upickup,
		row.Idcarrito,
		row.Almacen,
		row.RutaCodigo.String,
		row.RutaDescrip.String,
		row.KePedstatus,
	)
}

func mapGetOrderByCodeRowToOrderLine(row *database.GetOrderByCodeRow) *OrderLineResponse {
	cantRef := util.StringToFloat(row.Cantref.String, 0)
	cantidad := util.StringToFloat(row.Cantidad.String, 0)

	product := &ProductResponse{
		Nombre:     strings.TrimSpace(row.Nombre.String),
		Codigo:     strings.TrimSpace(row.Codigo.String),
		Referencia: strings.TrimSpace(row.Referencia.String),
		Marca:      strings.TrimSpace(row.Marca.String),
		Unidad:     strings.TrimSpace(row.Unidad.String),
		Image:      fmt.Sprintf("%s/api/products/%s/image", os.Getenv("BASE_URL"), row.Codigo.String),
		CreatedAt:  util.FormatDateForResponse(row.Fechacrea.Time),
	}

	return mapToOrderLine(
		row.Agencia_2.String,
		row.Tipodoc_2.String,
		row.Documento_2.String,
		row.Almacen_2.String,
		product,
		int(cantRef),
		int(cantidad),
	)
}
