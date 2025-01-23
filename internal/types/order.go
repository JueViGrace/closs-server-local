package types

import (
	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
)

type OrderWithLinesResponse struct {
	OrderResponse
	Lines []OrderLineResponse `json:"lines"`
}

type OrderResponse struct{}

type OrderLineResponse struct{}

func DbOrderToOrder(db *database.Operti) *OrderResponse {
	return mapToOrder()
}

func DbOrderLineToOrderLine(db *database.Opermv) *OrderLineResponse {
	return mapToOrderLine()
}

func mapToOrder() *OrderResponse {
	return &OrderResponse{}
}

func mapToOrderLine() *OrderLineResponse {
	return &OrderLineResponse{}
}

// func MapToOrderWithLines(head *OrderResponse, lines []OrderLineResponse) *OrderWithLinesResponse {
// 	return &OrderWithLinesResponse{
// 		OrderResponse: *head,
// 		Lines:         lines,
// 	}
// }
//
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
