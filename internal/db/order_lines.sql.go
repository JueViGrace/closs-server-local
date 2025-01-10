// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order_lines.sql

package db

import (
	"context"
)

const createOrderLine = `-- name: CreateOrderLine :one
insert into closs_order_lines (
    kti_tdoc,
    kti_ndoc,
    kti_tipprec,
    kmv_codart,
    kmv_nombre,
    kmv_cant,
    kmv_artprec,
    kmv_stot,
    kmv_dctolin
)
values (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING kti_tdoc, kti_ndoc, kti_tipprec, kmv_codart, kmv_nombre, kmv_cant, kmv_artprec, kmv_stot, kmv_dctolin
`

type CreateOrderLineParams struct {
	KtiTdoc    string
	KtiNdoc    string
	KtiTipprec int64
	KmvCodart  string
	KmvNombre  string
	KmvCant    int64
	KmvArtprec float64
	KmvStot    float64
	KmvDctolin float64
}

func (q *Queries) CreateOrderLine(ctx context.Context, arg CreateOrderLineParams) (ClossOrderLine, error) {
	row := q.db.QueryRowContext(ctx, createOrderLine,
		arg.KtiTdoc,
		arg.KtiNdoc,
		arg.KtiTipprec,
		arg.KmvCodart,
		arg.KmvNombre,
		arg.KmvCant,
		arg.KmvArtprec,
		arg.KmvStot,
		arg.KmvDctolin,
	)
	var i ClossOrderLine
	err := row.Scan(
		&i.KtiTdoc,
		&i.KtiNdoc,
		&i.KtiTipprec,
		&i.KmvCodart,
		&i.KmvNombre,
		&i.KmvCant,
		&i.KmvArtprec,
		&i.KmvStot,
		&i.KmvDctolin,
	)
	return i, err
}

const updateOrderLine = `-- name: UpdateOrderLine :one
update closs_order_lines set 
    kti_tdoc = ?,
    kti_tipprec = ?,
    kmv_nombre = ?,
    kmv_cant = ?,
    kmv_artprec = ?,
    kmv_stot = ?,
    kmv_dctolin = ?
WHERE kti_ndoc = ? and kmv_codart = ?
RETURNING kti_tdoc, kti_ndoc, kti_tipprec, kmv_codart, kmv_nombre, kmv_cant, kmv_artprec, kmv_stot, kmv_dctolin
`

type UpdateOrderLineParams struct {
	KtiTdoc    string
	KtiTipprec int64
	KmvNombre  string
	KmvCant    int64
	KmvArtprec float64
	KmvStot    float64
	KmvDctolin float64
	KtiNdoc    string
	KmvCodart  string
}

func (q *Queries) UpdateOrderLine(ctx context.Context, arg UpdateOrderLineParams) (ClossOrderLine, error) {
	row := q.db.QueryRowContext(ctx, updateOrderLine,
		arg.KtiTdoc,
		arg.KtiTipprec,
		arg.KmvNombre,
		arg.KmvCant,
		arg.KmvArtprec,
		arg.KmvStot,
		arg.KmvDctolin,
		arg.KtiNdoc,
		arg.KmvCodart,
	)
	var i ClossOrderLine
	err := row.Scan(
		&i.KtiTdoc,
		&i.KtiNdoc,
		&i.KtiTipprec,
		&i.KmvCodart,
		&i.KmvNombre,
		&i.KmvCant,
		&i.KmvArtprec,
		&i.KmvStot,
		&i.KmvDctolin,
	)
	return i, err
}
