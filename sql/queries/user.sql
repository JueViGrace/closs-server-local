-- name: GetUsers :many
select *
from ke_wusuarios
;

-- name: GetUserByUsername :one
select *
from ke_wusuarios
where username = ?
;

-- name: UpdatePassword :exec
update ke_wusuarios set
    mail_password = ?
where username = ?;

