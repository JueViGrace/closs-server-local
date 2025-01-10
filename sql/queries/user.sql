-- name: GetUsers :many
select *
from closs_user
;

-- name: GetUserById :one
select *
from closs_user
where id = ?
;

-- name: GetUserByUsername :one
select *
from closs_user
where username = ? and deleted_at is null
;

-- name: CreateUser :one
insert into closs_user (
    id,
    username,
    password,
    codigo,
    role,
    ult_sinc,
    version,
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
    ?
)
RETURNING *;

-- name: UpdatePassword :exec
update closs_user set
    password = ?,
    updated_at = ?
where id = ?
RETURNING *;

-- name: UpdateLastSync :exec
update closs_user set
    ult_sinc = ?,
    version = ?,
    updated_at = ?
where id = ?
RETURNING *;

-- name: SoftDeleteUser :exec
update closs_user set 
    updated_at = ?,
    deleted_at = ?
where id = ?;

-- name: DeleteUserById :exec
delete from closs_user
where id = ?
;

