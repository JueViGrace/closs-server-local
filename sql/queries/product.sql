-- name: GetProducts :many
select *
from articulo
;

-- name: GetExistingProducts :many
select *
from articulo
where (existencia - comprometido) > 0 and discont = 0
;

-- name: GetProductByCode :one
select *
from articulo
where codigo = ?
;

-- name: GetExistingProductByCode :one
select *
from articulo
where
    codigo = ?
    and (existencia - comprometido) > 0
    and discont = 0
    and deleted_at is null
;

