-- name: GetConfigs :many
select *
from closs_config
;

-- name: GetConfigsByUsername :many
select *
from closs_config
where username = ?
;

-- name: CreateConfig :one
INSERT INTO closs_config (
    cnfg_idconfig,
    cnfg_clase,
    cnfg_tipo,
    cnfg_valnum,
    cnfg_valsino,
    cnfg_valtxt,
    cnfg_lentxt,
    cnfg_valfch,
    cnfg_activa,
    cnfg_etiq,
    cnfg_ttip,
    username,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateConfig :one
UPDATE closs_config SET 
    cnfg_clase = ?,
    cnfg_tipo = ?,
    cnfg_valnum = ?,
    cnfg_valsino = ?,
    cnfg_valtxt = ?,
    cnfg_lentxt = ?,
    cnfg_valfch = ?,
    cnfg_activa = ?,
    cnfg_etiq = ?,
    cnfg_ttip = ?,
    updated_at = ?
WHERE cnfg_idconfig = ? and username = ?
RETURNING *;

-- name: DeleteConfig :exec
delete from closs_config
where cnfg_idconfig = ? and username = ?
;

