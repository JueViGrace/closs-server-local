-- name: GetSessionById :one
select *
from ke_session
where user_id = ?
;

-- name: GetSessionByToken :one
select *
from ke_session
where token = ?
;

-- name: CreateSession :exec
INSERT INTO ke_session(
    token,
    user_id
)
VALUES (?, ?);

-- name: UpdateSession :exec
UPDATE ke_session SET
    token = ?
WHERE user_id = ?;

-- name: DeleteSessionById :exec
delete from ke_session
where user_id = ?
;

-- name: DeleteSessionByToken :exec
delete from ke_session
where token = ?
;

