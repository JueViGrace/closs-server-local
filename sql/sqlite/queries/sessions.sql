-- name: GetSessionById :one
select *
from closs_session
where user_id = ?
;

-- name: GetSessionByUsername :one
select *
from closs_session
where username = ?
;

-- name: CreateSession :exec
INSERT  INTO closs_session(
    refresh_token,
    access_token,
    username,
    user_id
)
VALUES (?, ?, ?, ?)
;

-- name: UpdateSession :exec
UPDATE closs_session SET
    refresh_token = ?,
    access_token = ?
WHERE user_id = ?
;

-- name: DeleteSessionById :exec
delete from closs_session
where user_id = ?
;

-- name: DeleteSessionByToken :exec
delete from closs_session
where refresh_token = ? or access_token = ?
;

