package types

type OrderWithLinesResponse struct {
	OrderResponse
	Lines []OrderLineResponse `json:"lines"`
}

type OrderResponse struct {
	KtiNdoc        string  `json:"ktiNdoc"`
	KtiTdoc        string  `json:"ktiTdoc"`
	KtiCodcli      string  `json:"ktiCodcli"`
	KtiNombrecli   string  `json:"ktiNombrecli"`
	KtiCodven      string  `json:"ktiCodven"`
	KtiDocsol      string  `json:"ktiDocsol"`
	KtiCondicion   string  `json:"ktiCondicion"`
	KtiTipprec     int     `json:"ktiTipprec"`
	KtiTotneto     float64 `json:"ktiTotneto"`
	KtiStatus      int     `json:"ktiStatus"`
	KtiNroped      string  `json:"ktiNroped"`
	KtiFchdoc      string  `json:"ktiFchdoc"`
	KtiNegesp      bool    `json:"ktiNegesp"`
	KePedstatus    int     `json:"kePedstatus"`
	Dolarflete     bool    `json:"dolarflete"`
	Complemento    bool    `json:"complemento"`
	NroComplemento string  `json:"nroComplemento"`
	CreatedAt      string  `json:"createdAt"`
	UpdatedAt      string  `json:"updatedAt"`
}

type OrderLineResponse struct {
	KtiTdoc    string  `json:"ktiTdoc"`
	KtiNdoc    string  `json:"ktiNdoc"`
	KtiTipprec int     `json:"ktiTipprec"`
	KmvCodart  string  `json:"kmvCodart"`
	KmvNombre  string  `json:"kmvNombre"`
	KmvCant    int     `json:"kmvCant"`
	KmvArtprec float64 `json:"kmvArtprec"`
	KmvStot    float64 `json:"kmvStot"`
	KmvDctolin float64 `json:"kmvDctolin"`
}

// func DbOrderToOrder(db *db.ClossOrder) *OrderResponse {
// 	return mapToOrder(
// 		db.KtiNdoc,
// 		db.KtiTdoc,
// 		db.KtiCodcli,
// 		db.KtiNombrecli,
// 		db.KtiCodven,
// 		db.KtiDocsol,
// 		db.KtiCondicion,
// 		int(db.KtiTipprec),
// 		db.KtiTotneto,
// 		int(db.KtiStatus),
// 		db.KtiNroped,
// 		db.KtiFchdoc,
// 		db.KtiNegesp == 1,
// 		int(db.KePedstatus),
// 		db.Dolarflete == 1,
// 		db.Complemento == 1,
// 		db.NroComplemento,
// 		db.CreatedAt,
// 		db.UpdatedAt,
// 	)
// }
//
// func DbOrderLineToOrderLine(db *db.ClossOrderLine) *OrderLineResponse {
// 	return mapToOrderLine(
// 		db.KtiTdoc,
// 		db.KtiNdoc,
// 		int(db.KtiTipprec),
// 		db.KmvCodart,
// 		db.KmvNombre,
// 		int(db.KmvCant),
// 		db.KmvArtprec,
// 		db.KmvStot,
// 		db.KmvDctolin,
// 	)
// }
//
// func mapToOrder(
// 	ktiNdoc string,
// 	ktiTdoc string,
// 	ktiCodcli string,
// 	ktiNombrecli string,
// 	ktiCodven string,
// 	ktiDocsol string,
// 	ktiCondicion string,
// 	ktiTipprec int,
// 	ktiTotneto float64,
// 	ktiStatus int,
// 	ktiNroped string,
// 	ktiFchdoc string,
// 	ktiNegesp bool,
// 	kePedstatus int,
// 	dolarflete bool,
// 	complemento bool,
// 	nroComplemento string,
// 	createdAt string,
// 	updatedAt string,
// ) *OrderResponse {
// 	return &OrderResponse{
// 		KtiNdoc:        ktiNdoc,
// 		KtiTdoc:        ktiTdoc,
// 		KtiCodcli:      ktiCodcli,
// 		KtiNombrecli:   ktiNombrecli,
// 		KtiCodven:      ktiCodven,
// 		KtiDocsol:      ktiDocsol,
// 		KtiCondicion:   ktiCondicion,
// 		KtiTipprec:     ktiTipprec,
// 		KtiTotneto:     ktiTotneto,
// 		KtiStatus:      ktiStatus,
// 		KtiNroped:      ktiNroped,
// 		KtiFchdoc:      ktiFchdoc,
// 		KtiNegesp:      ktiNegesp,
// 		KePedstatus:    kePedstatus,
// 		Dolarflete:     dolarflete,
// 		Complemento:    complemento,
// 		NroComplemento: nroComplemento,
// 		CreatedAt:      createdAt,
// 		UpdatedAt:      updatedAt,
// 	}
// }
//
// func mapToOrderLine(
// 	ktiTdoc string,
// 	ktiNdoc string,
// 	ktiTipprec int,
// 	kmvCodart string,
// 	kmvNombre string,
// 	kmvCant int,
// 	kmvArtprec float64,
// 	kmvStot float64,
// 	kmvDctolin float64,
// ) *OrderLineResponse {
// 	return &OrderLineResponse{
// 		KtiTdoc:    ktiTdoc,
// 		KtiNdoc:    ktiNdoc,
// 		KtiTipprec: ktiTipprec,
// 		KmvCodart:  kmvCodart,
// 		KmvNombre:  kmvNombre,
// 		KmvCant:    kmvCant,
// 		KmvArtprec: kmvArtprec,
// 		KmvStot:    kmvStot,
// 		KmvDctolin: kmvDctolin,
// 	}
// }
//
// func MapToOrderWithLines(head *OrderResponse, lines []OrderLineResponse) *OrderWithLinesResponse {
// 	return &OrderWithLinesResponse{
// 		OrderResponse: *head,
// 		Lines:         lines,
// 	}
// }
//
// func GroupOrderWithLinesRow(rows []db.GetOrdersWithLinesRow) []OrderWithLinesResponse {
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
// func GroupOrderWithLinesByCodeRow(rows []db.GetOrderWithLinesByCodeRow) *OrderWithLinesResponse {
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
// func mapGetOrdersRowToOrder(row *db.GetOrdersWithLinesRow) *OrderResponse {
// 	return mapToOrder(
// 		row.KtiNdoc,
// 		row.KtiTdoc,
// 		row.KtiCodcli,
// 		row.KtiNombrecli,
// 		row.KtiCodven,
// 		row.KtiDocsol,
// 		row.KtiCondicion,
// 		int(row.KtiTipprec),
// 		row.KtiTotneto,
// 		int(row.KtiStatus),
// 		row.KtiNroped,
// 		row.KtiFchdoc,
// 		row.KtiNegesp == 1,
// 		int(row.KePedstatus),
// 		row.Dolarflete == 1,
// 		row.Complemento == 1,
// 		row.NroComplemento,
// 		row.CreatedAt,
// 		row.UpdatedAt,
// 	)
// }
//
// func mapGetOrdersRowToOrderLine(row *db.GetOrdersWithLinesRow) *OrderLineResponse {
// 	return mapToOrderLine(
// 		row.KtiTdoc_2.String,
// 		row.KtiNdoc_2.String,
// 		int(row.KtiTipprec_2.Int64),
// 		row.KmvCodart.String,
// 		row.KmvNombre.String,
// 		int(row.KmvCant.Int64),
// 		row.KmvArtprec.Float64,
// 		row.KmvStot.Float64,
// 		row.KmvDctolin.Float64,
// 	)
// }
//
// func mapGetOrdersRowByCodeToOrder(row *db.GetOrderWithLinesByCodeRow) *OrderResponse {
// 	return mapToOrder(
// 		row.KtiNdoc,
// 		row.KtiTdoc,
// 		row.KtiCodcli,
// 		row.KtiNombrecli,
// 		row.KtiCodven,
// 		row.KtiDocsol,
// 		row.KtiCondicion,
// 		int(row.KtiTipprec),
// 		row.KtiTotneto,
// 		int(row.KtiStatus),
// 		row.KtiNroped,
// 		row.KtiFchdoc,
// 		row.KtiNegesp == 1,
// 		int(row.KePedstatus),
// 		row.Dolarflete == 1,
// 		row.Complemento == 1,
// 		row.NroComplemento,
// 		row.CreatedAt,
// 		row.UpdatedAt,
// 	)
// }
//
// func mapGetOrdersRowByCodeToOrderLine(row *db.GetOrderWithLinesByCodeRow) *OrderLineResponse {
// 	return mapToOrderLine(
// 		row.KtiTdoc_2.String,
// 		row.KtiNdoc_2.String,
// 		int(row.KtiTipprec_2.Int64),
// 		row.KmvCodart.String,
// 		row.KmvNombre.String,
// 		int(row.KmvCant.Int64),
// 		row.KmvArtprec.Float64,
// 		row.KmvStot.Float64,
// 		row.KmvDctolin.Float64,
// 	)
// }
