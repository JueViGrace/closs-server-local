package types

import (
	"time"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
)

type OrderWithLinesResponse struct {
	OrderResponse
	Lines []OrderLineResponse `json:"lines"`
}

type OrderResponse struct {
	Agencia    string    `json:"agencia"`
	TipoDoc    string    `json:"tipodoc"`
	Documento  string    `json:"documento"`
	CodCliente string    `json:"codcliente"`
	NombreCli  string    `json:"nombrecli"`
	Emision    time.Time `json:"emision"`
	Upickup    string    `json:"upickup"`
	IdCarrito  string    `json:"idcarrito"`
	Almacen    string    `json:"almacen"`
}

type OrderLineResponse struct {
	Agencia   string `json:"agencia"`
	TipoDoc   string `json:"tipodoc"`
	Documento string `json:"documento"`
	Codigo    string `json:"codigo"`
	Almacen   string `json:"almacen"`
	CantRef   string `json:"cantref"`
	Cantidad  string `json:"cantidad"`
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
) *OrderResponse {
	return &OrderResponse{
		Agencia:    agencia,
		TipoDoc:    tipodoc,
		Documento:  documento,
		CodCliente: codcliente,
		NombreCli:  nombrecli,
		Emision:    emision,
		Upickup:    upickup,
		IdCarrito:  idcarrito,
		Almacen:    almacen,
	}
}

func mapToOrderLine(
	agencia string,
	tipodoc string,
	documento string,
	codigo string,
	almacen string,
	cantref string,
	cantidad string,
) *OrderLineResponse {
	return &OrderLineResponse{
		Agencia:   agencia,
		TipoDoc:   tipodoc,
		Documento: documento,
		Codigo:    codigo,
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

func GroupOrderByUserRow(rows []database.GetOrdersByUserRow) []OrderWithLinesResponse {
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

		group[*ord] = append(group[*ord], *mapGetOrdersByUserRowToOrderLine(&row))
	}

	for key, value := range group {
		resp = append(resp, *MapToOrderWithLines(&key, value))
	}

	return resp
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
	)
}

func mapGetOrdersByUserRowToOrderLine(row *database.GetOrdersByUserRow) *OrderLineResponse {
	return mapToOrderLine(
		row.Agencia,
		row.Tipodoc,
		row.Documento,
		row.Codigo.String,
		row.Almacen,
		row.Cantref.String,
		row.Cantidad.String,
	)
}

// func GroupOrderWithLinesRow(rows []database.GetOrdersWithLinesRow) []OrderWithLinesResponse {
// 	orders := make([]OrderWithLinesResponse, 0)
// 	group := make(map[OrderResponse][]OrderLineResponse)
// 	ord := new(OrderResponse)
//
// 	for _, row := range rows {
// 		if ord == nil {
// 			ord = mapGetOrdersRowToOrder(&row)
// 		}
//
// 		if ord.KtiNdoc != row.KtiNdoc {
// 			ord = mapGetOrdersRowToOrder(&row)
// 		}
//
// 		group[*ord] = append(group[*ord], *mapGetOrdersRowToOrderLine(&row))
// 	}
//
// 	for key, value := range group {
// 		orders = append(orders, *MapToOrderWithLines(&key, value))
// 	}
//
// 	return orders
// }
//
// func GroupOrderWithLinesByCodeRow(rows []database.GetOrderWithLinesByCodeRow) *OrderWithLinesResponse {
// 	order := new(OrderWithLinesResponse)
// 	group := make(map[OrderResponse][]OrderLineResponse)
// 	ord := new(OrderResponse)
//
// 	for _, row := range rows {
// 		if ord == nil {
// 			ord = mapGetOrdersRowByCodeToOrder(&row)
// 		}
//
// 		if order.KtiNdoc != row.KtiNdoc {
// 			ord = mapGetOrdersRowByCodeToOrder(&row)
// 		}
//
// 		group[*ord] = append(group[*ord], *mapGetOrdersRowByCodeToOrderLine(&row))
// 	}
//
// 	for key, value := range group {
// 		order = MapToOrderWithLines(&key, value)
// 	}
//
// 	return order
// }
//
// func mapGetOrdersRowToOrder(row *database.GetOrdersWithLinesRow) *OrderResponse {
// 	return mapToOrder()
// }
//
// func mapGetOrdersRowToOrderLine(row *database.GetOrdersWithLinesRow) *OrderLineResponse {
// 	return mapToOrderLine()
// }
//
// func mapGetOrdersRowByCodeToOrder(row *database.GetOrderWithLinesByCodeRow) *OrderResponse {
// 	return mapToOrder()
// }
//
// func mapGetOrdersRowByCodeToOrderLine(row *database.GetOrderWithLinesByCodeRow) *OrderLineResponse {
// 	return mapToOrderLine()
// }
