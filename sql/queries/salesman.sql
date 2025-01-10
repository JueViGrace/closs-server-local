-- name: GetSalesmenByManager :many
select
    closs_salesman.codigo,
    closs_user.username,
    closs_salesman.nombre,
    closs_salesman.email,
    closs_salesman.telefono,
    closs_salesman.telefonos,
    closs_salesman.supervpor,
    closs_sector.zona as sector,
    closs_subsector.subsector as subsector,
    closs_user.ult_sinc,
    closs_user.version,
    closs_salesman.created_at,
    closs_salesman.updated_at
from closs_salesman
left join closs_sector on closs_sector.codigo = closs_salesman.sector
left join closs_subsector on closs_subsector.subcodigo = closs_salesman.subcodigo
left join closs_user on closs_salesman.codigo = closs_user.codigo
left join closs_managers on closs_salesman.supervpor = closs_managers.kng_codcoord
where closs_managers.kng_codgcia = ?
group by closs_salesman.codigo
;

-- name: GetExistingSalesmenByManager :many
select
    closs_salesman.codigo,
    closs_user.username,
    closs_salesman.nombre,
    closs_salesman.email,
    closs_salesman.telefono,
    closs_salesman.telefonos,
    closs_salesman.supervpor,
    closs_sector.zona as sector,
    closs_subsector.subsector as subsector,
    closs_user.ult_sinc,
    closs_user.version,
    closs_salesman.created_at,
    closs_salesman.updated_at
from closs_salesman
left join closs_sector on closs_sector.codigo = closs_salesman.sector
left join closs_subsector on closs_subsector.subcodigo = closs_salesman.subcodigo
left join closs_user on closs_salesman.codigo = closs_user.codigo
left join closs_managers on closs_salesman.supervpor = closs_managers.kng_codcoord
where closs_salesman.status = 1 and closs_managers.kng_codgcia = ?
group by closs_salesman.codigo
;

-- name: GetSalesmanByCode :one
select
    closs_salesman.codigo,
    closs_user.username,
    closs_salesman.nombre,
    closs_salesman.email,
    closs_salesman.telefono,
    closs_salesman.telefonos,
    closs_salesman.supervpor,
    closs_sector.zona as sector,
    closs_subsector.subsector as subsector,
    closs_user.ult_sinc,
    closs_user.version,
    closs_salesman.created_at,
    closs_salesman.updated_at
from closs_salesman
left join closs_sector on closs_sector.codigo = closs_salesman.sector
left join closs_subsector on closs_subsector.subcodigo = closs_salesman.subcodigo
left join closs_user on closs_salesman.codigo = closs_user.codigo
where closs_salesman.codigo = ?
;

-- name: GetExistingSalesmanByCode :one
select
    closs_salesman.codigo,
    closs_user.username,
    closs_salesman.nombre,
    closs_salesman.email,
    closs_salesman.telefono,
    closs_salesman.telefonos,
    closs_salesman.supervpor,
    closs_sector.zona as sector,
    closs_subsector.subsector as subsector,
    closs_user.ult_sinc,
    closs_user.version,
    closs_salesman.created_at,
    closs_salesman.updated_at
from closs_salesman
left join closs_sector on closs_sector.codigo = closs_salesman.sector
left join closs_subsector on closs_subsector.subcodigo = closs_salesman.subcodigo
left join closs_user on closs_salesman.codigo = closs_user.codigo
where closs_salesman.codigo = ? and closs_salesman.status = 1
;

-- name: CreateSalesman :one
insert into closs_salesman (
    codigo,
    nombre,
    telefono,
    telefonos,
    status,
    supervpor,
    sector,
    subcodigo,
    email,
    nivgcial,
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
    ?
)
RETURNING *;

-- name: UpdateSalesman :one
update closs_salesman set 
    nombre = ?,
    telefono = ?,
    telefonos = ?,
    supervpor = ?,
    sector = ?,
    subcodigo = ?,
    email = ?,
    updated_at = ?
where codigo = ?
RETURNING *;

