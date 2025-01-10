-- name: GetProducts :many
select *
from closs_product
;

-- name: GetExistingProducts :many
select *
from closs_product
where (existencia - comprometido) > 0 and discont = 0 and deleted_at is null
;

-- name: GetProductByCode :one
select *
from closs_product
where codigo = ?
;

-- name: GetExistingProductByCode :one
select *
from closs_product
where
    codigo = ?
    and (existencia - comprometido) > 0
    and discont = 0
    and deleted_at is null
;

-- name: CreateProduct :one
insert into closs_product (
    codigo,
    grupo,
    subgrupo,
    nombre,
    referencia,
    marca,
    unidad,
    existencia,
    precio1,
    precio2,
    precio3,
    precio4,
    precio5,
    precio6,
    precio7,
    discont,
    vta_max,
    vta_min,
    dctotope,
    preventa,
    comprometido,
    vta_minenx,
    vta_solofac,
    vta_solone,
    codbarras,
    detalles,
    cantbulto,
    costo_prom,
    util1,
    util2,
    util3,
    fchultcomp,
    qtyultcomp,
    created_at,
    updated_at
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
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
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

-- name: UpdateProduct :one
update closs_product set 
    grupo = ?,
    subgrupo = ?,
    nombre = ?,
    referencia = ?,
    marca = ?,
    unidad = ?,
    existencia = ?,
    precio1 = ?,
    precio2 = ?,
    precio3 = ?,
    precio4 = ?,
    precio5 = ?,
    precio6 = ?,
    precio7 = ?,
    discont = ?,
    vta_max = ?,
    vta_min = ?,
    dctotope = ?,
    preventa = ?,
    comprometido = ?,
    vta_minenx = ?,
    vta_solofac = ?,
    vta_solone = ?,
    codbarras = ?,
    detalles = ?,
    cantbulto = ?,
    costo_prom = ?,
    util1 = ?,
    util2 = ?,
    util3 = ?,
    fchultcomp = ?,
    qtyultcomp = ?,
    updated_at = ?
where codigo = ?
RETURNING *;

-- name: SoftDeleteProduct :exec
update closs_product set 
    discont = 1,
    updated_at = ?,
    deleted_at = ?
where codigo = ?;

-- name: DeleteProduct :exec
delete from closs_product
where codigo = ?
;

