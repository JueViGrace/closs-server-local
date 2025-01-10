-- name: CreateOrderLine :one
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
RETURNING *;

-- name: UpdateOrderLine :one
update closs_order_lines set 
    kti_tdoc = ?,
    kti_tipprec = ?,
    kmv_nombre = ?,
    kmv_cant = ?,
    kmv_artprec = ?,
    kmv_stot = ?,
    kmv_dctolin = ?
WHERE kti_ndoc = ? and kmv_codart = ?
RETURNING *;

