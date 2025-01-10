-- name: GetStatistics :many
select *
from closs_salesman_statistic
;

-- name: GetStatisticsByManager :many
select closs_salesman_statistic.*
from closs_salesman_statistic
left join
    closs_managers on closs_salesman_statistic.codcoord = closs_managers.kng_codcoord
where closs_managers.kng_codgcia = ?
;

-- name: GetStatisticBySalesman :one
select *
from closs_salesman_statistic
where vendedor = ?
;

-- name: CreateStatistic :one
insert into closs_salesman_statistic(
    codcoord,
    nomcoord,
    vendedor,
    nombrevend,
    cntpedidos,
    mtopedidos,
    cntfacturas,
    mtofacturas,
    metavend,
    prcmeta,
    cntclientes,
    clivisit,
    prcvisitas,
    lom_montovtas,
    lom_prcvtas,
    lom_prcvisit,
    rlom_montovtas,
    rlom_prcvtas,
    rlom_prcvisit,
    fecha_estad,
    ppgdol_totneto,
    devdol_totneto,
    defdol_totneto,
    totdolcob,
    cntrecl,
    mtorecl,
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
    ?
)
RETURNING *;

-- name: UpdateStatistic :one
update closs_salesman_statistic set 
    codcoord = ?,
    nomcoord = ?,
    nombrevend = ?,
    cntpedidos = ?,
    mtopedidos = ?,
    cntfacturas = ?,
    mtofacturas = ?,
    metavend = ?,
    prcmeta = ?,
    cntclientes = ?,
    clivisit = ?,
    prcvisitas = ?,
    lom_montovtas = ?,
    lom_prcvtas = ?,
    lom_prcvisit = ?,
    rlom_montovtas = ?,
    rlom_prcvtas = ?,
    rlom_prcvisit = ?,
    fecha_estad = ?,
    ppgdol_totneto = ?,
    devdol_totneto = ?,
    defdol_totneto = ?,
    totdolcob = ?,
    cntrecl = ?,
    mtorecl = ?,
    updated_at = ?
where vendedor = ?
RETURNING *;

