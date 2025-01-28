package types

import (
	"strconv"
	"time"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
)

type OrderWithLinesResponse struct {
	OrderResponse
	Lines []OrderLineResponse `json:"lines"`
}

type OrderResponse struct {
	Agencia     string    `json:"agencia"`
	TipoDoc     string    `json:"tipodoc"`
	Documento   string    `json:"documento"`
	CodCliente  string    `json:"codcliente"`
	NombreCli   string    `json:"nombrecli"`
	Emision     time.Time `json:"emision"`
	Upickup     string    `json:"upickup"`
	IdCarrito   string    `json:"idcarrito"`
	Almacen     string    `json:"almacen"`
	KePedStatus string    `json:"ke_pedstatus"`
}

type OrderLineResponse struct {
	Agencia   string `json:"agencia"`
	TipoDoc   string `json:"tipodoc"`
	Documento string `json:"documento"`
	Codigo    string `json:"codigo"`
	Nombre    string `json:"nombre"`
	Almacen   string `json:"almacen"`
	CantRef   int    `json:"cantref"`
	Cantidad  int    `json:"cantidad"`
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
	kePedstatus string,
) *OrderResponse {
	return &OrderResponse{
		Agencia:     agencia,
		TipoDoc:     tipodoc,
		Documento:   documento,
		CodCliente:  codcliente,
		NombreCli:   nombrecli,
		Emision:     emision,
		Upickup:     upickup,
		IdCarrito:   idcarrito,
		Almacen:     almacen,
		KePedStatus: kePedstatus,
	}
}

func mapToOrderLine(
	agencia string,
	tipodoc string,
	documento string,
	codigo string,
	nombre string,
	almacen string,
	cantref int,
	cantidad int,
) *OrderLineResponse {
	return &OrderLineResponse{
		Agencia:   agencia,
		TipoDoc:   tipodoc,
		Documento: documento,
		Codigo:    codigo,
		Nombre:    nombre,
		Almacen:   almacen,
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

func GroupOrderByUserRow(rows []database.GetOrdersByUserRow) ([]OrderWithLinesResponse, error) {
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

		line, err := mapGetOrdersByUserRowToOrderLine(&row)
		if err != nil {
			return nil, err
		}

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
		row.KePedstatus,
	)
}

func mapGetOrdersByUserRowToOrderLine(row *database.GetOrdersByUserRow) (*OrderLineResponse, error) {
	cantRef, err := strconv.ParseFloat(row.Cantref.String, 0)
	if err != nil {
		return nil, err
	}
	cantidad, err := strconv.ParseFloat(row.Cantidad.String, 0)
	if err != nil {
		return nil, err
	}

	return mapToOrderLine(
		row.Agencia,
		row.Tipodoc,
		row.Documento,
		row.Codigo.String,
		row.Nombre.String,
		row.Almacen,
		int(cantRef),
		int(cantidad),
	), nil
}
