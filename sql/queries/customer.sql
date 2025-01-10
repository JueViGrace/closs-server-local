-- name: GetCustomers :many
select *
from closs_customer
;

-- name: GetCustomerByCode :one
select *
from closs_customer
where codigo = ?
;

-- name: GetCustomersByManager :many
select closs_customer.*
from closs_customer
left join closs_salesman on closs_customer.vendedor = closs_salesman.codigo
left join closs_managers on closs_salesman.supervpor = closs_managers.kng_codcoord
where closs_managers.kng_codgcia =:manager
order by
    closs_salesman.supervpor asc, closs_customer.vendedor asc, closs_customer.nombre asc
;

-- name: GetCustomersBySalesman :many
select *
from closs_customer
where closs_customer.vendedor =:code
order by closs_customer.nombre asc
;

-- name: CreateCustomer :one
INSERT INTO closs_customer (
    codigo,
    nombre,
    email,
    direccion,
    telefonos,
    perscont,
    vendedor,
    contribespecial,
    status,
    sector,
    subsector,
    precio,
    kne_activa,
    kne_mtomin,
    noemifac,
    noeminota,
    fchultvta,
    mtoultvta,
    prcdpagdia,
    promdiasp,
    riesgocrd,
    cantdocs,
    totmtodocs,
    prommtodoc,
    diasultvta,
    promdiasvta,
    limcred,
    fchcrea,
    dolarflete,
    nodolarflete,
    created_at,
    updated_at
)
VALUES (
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

-- name: UpdateCustomer :one
UPDATE closs_customer SET 
    nombre = ?,
    direccion = ?,
    telefonos = ?,
    perscont = ?,
    vendedor = ?,
    contribespecial = ?,
    status = ?,
    sector = ?,
    subsector = ?,
    precio = ?,
    email = ?,
    kne_activa = ?,
    kne_mtomin = ?,
    noemifac = ?,
    noeminota = ?,
    fchultvta = ?,
    mtoultvta = ?,
    prcdpagdia = ?,
    promdiasp = ?,
    riesgocrd = ?,
    cantdocs = ?,
    totmtodocs = ?,
    prommtodoc = ?,
    diasultvta = ?,
    promdiasvta = ?,
    limcred = ?,
    dolarflete = ?,
    nodolarflete = ?,
    updated_at = ?
WHERE codigo = ?
RETURNING *;

