-- name: GetManagersByCode :many
select *
from closs_managers
where kng_codgcia = ?
;

-- name: GetManagerByCode :one
select *
from closs_managers
where kng_codgcia = ? and kng_codcoord = ?
;

-- name: CreateManager :one
INSERT OR REPLACE INTO closs_managers(
    kng_codgcia,
    kng_codcoord
)
VALUES (?, ?)
RETURNING *;

