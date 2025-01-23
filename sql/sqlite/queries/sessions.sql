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

-- name: GetSessionByToken :one
select *
from closs_session
where token = ?
;

-- name: CreateSession :exec
INSERT  INTO closs_session(
    token,
    username,
    user_id
)
VALUES (?, ?, ?)
;

-- name: UpdateSession :exec
UPDATE closs_session SET
    token = ?
WHERE user_id = ?
;

-- name: DeleteSessionById :exec
delete from closs_session
where user_id = ?
;

-- name: DeleteSessionByToken :exec
delete from closs_session
where token = ?
;

