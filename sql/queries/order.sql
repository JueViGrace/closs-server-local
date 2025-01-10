-- name: GetOrders :many
select *
from closs_order
;

-- name: GetOrdersWithLines :many
select *
from closs_order
left join closs_order_lines on closs_order.kti_ndoc = closs_order_lines.kti_ndoc
left join closs_product on closs_product.codigo = closs_order_lines.kmv_codart
;

-- name: GetOrderByCode :one
select *
from closs_order
where kti_ndoc = ?
;

-- name: GetOrderWithLinesByCode :many
select *
from closs_order
left join closs_order_lines on closs_order.kti_ndoc = closs_order_lines.kti_ndoc
left join closs_product on closs_product.codigo = closs_order_lines.kmv_codart
where closs_order.kti_ndoc = ?
;

-- name: GetAllOrdersByManager :many
select closs_order.*
from closs_order
left join closs_salesman on closs_order.kti_codven = closs_salesman.codigo
left join closs_managers on closs_salesman.supervpor = closs_managers.kng_codcoord
where closs_managers.kng_codgcia = ?
order by closs_order.kti_codven asc, closs_order.kti_fchdoc asc
;

-- name: GetOrdersByManager :many
select closs_order.*
from closs_order
left join closs_salesman on closs_order.kti_codven = closs_salesman.codigo
left join closs_managers on closs_salesman.supervpor = closs_managers.kng_codcoord
where
    closs_managers.kng_codgcia = ?
    and (closs_order.kti_status = 4 or closs_order.kti_status = 5)
    and closs_order.kti_nroped != ''
    and month(closs_order.kti_fchdoc) = month(now())
    and year(kti_fchdoc) = year(now())
order by closs_order.kti_codven asc, closs_order.kti_fchdoc asc
;

-- name: GetAllOrdersBySalesman :many
select closs_order.*
from closs_order
left join closs_salesman on closs_order.kti_codven = closs_salesman.codigo
where closs_salesman.codigo = ?
order by closs_order.kti_codven asc, closs_order.kti_fchdoc asc
;

-- name: GetOrdersBySalesman :many
select closs_order.*
from closs_order
left join closs_salesman on closs_order.kti_codven = closs_salesman.codigo
where
    closs_salesman.codigo = ?
    and (closs_order.kti_status = 4 or closs_order.kti_status = 5)
    and closs_order.kti_nroped != ''
    and month(closs_order.kti_fchdoc) = month(now())
    and year(kti_fchdoc) = year(now())
order by closs_order.kti_codven asc, closs_order.kti_fchdoc asc
;

-- name: GetAllOrdersByCustomer :many
select closs_order.*
from closs_order
left join closs_customer on closs_order.kti_codcli = closs_customer.codigo
where closs_customer.codigo = ?
order by closs_order.kti_codven asc, closs_order.kti_fchdoc asc
;

-- name: GetOrdersByCustomer :many
select closs_order.*
from closs_order
left join closs_customer on closs_order.kti_codcli = closs_customer.codigo
where
    closs_customer.codigo = ?
    and (closs_order.kti_status = 4 or closs_order.kti_status = 5)
    and closs_order.kti_nroped != ''
    and month(closs_order.kti_fchdoc) = month(now())
    and year(kti_fchdoc) = year(now())
order by closs_order.kti_codven asc, closs_order.kti_fchdoc asc
;

-- name: CreateOrder :one
insert into closs_order (
    kti_ndoc,
    kti_tdoc,
    kti_codcli,
    kti_nombrecli,
    kti_codven,
    kti_docsol,
    kti_condicion,
    kti_tipprec,
    kti_totneto,
    kti_status,
    kti_nroped,
    kti_fchdoc,
    kti_negesp,
    ke_pedstatus,
    dolarflete,
    complemento,
    nro_complemento,
    created_at,
    updated_at
)
values(
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

-- name: UpdateOrder :one
update closs_order set 
    kti_tdoc = ?,
    kti_codcli = ?,
    kti_nombrecli = ?,
    kti_codven = ?,
    kti_docsol = ?,
    kti_condicion = ?,
    kti_tipprec = ?,
    kti_totneto = ?,
    kti_status = ?,
    kti_nroped = ?,
    kti_fchdoc = ?,
    kti_negesp = ?,
    ke_pedstatus = ?,
    dolarflete = ?,
    complemento = ?,
    nro_complemento = ?,
    updated_at = ?
WHERE kti_ndoc = ?
RETURNING *;

