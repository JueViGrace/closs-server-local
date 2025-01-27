// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sessions.sql

package database

import (
	"context"
)

const createSession = `-- name: CreateSession :exec
;

INSERT  INTO closs_session(
    refresh_token,
    access_token,
    username,
    user_id
)
VALUES (?, ?, ?, ?)
`

type CreateSessionParams struct {
	RefreshToken string
	AccessToken  string
	Username     string
	UserID       string
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.db.ExecContext(ctx, createSession,
		arg.RefreshToken,
		arg.AccessToken,
		arg.Username,
		arg.UserID,
	)
	return err
}

const deleteSessionById = `-- name: DeleteSessionById :exec
;

delete from closs_session
where user_id = ?
`

func (q *Queries) DeleteSessionById(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, deleteSessionById, userID)
	return err
}

const getSessionById = `-- name: GetSessionById :one
select refresh_token, access_token, username, user_id
from closs_session
where user_id = ?
`

func (q *Queries) GetSessionById(ctx context.Context, userID string) (ClossSession, error) {
	row := q.db.QueryRowContext(ctx, getSessionById, userID)
	var i ClossSession
	err := row.Scan(
		&i.RefreshToken,
		&i.AccessToken,
		&i.Username,
		&i.UserID,
	)
	return i, err
}

const getSessionByUsername = `-- name: GetSessionByUsername :one
;

select refresh_token, access_token, username, user_id
from closs_session
where username = ?
`

func (q *Queries) GetSessionByUsername(ctx context.Context, username string) (ClossSession, error) {
	row := q.db.QueryRowContext(ctx, getSessionByUsername, username)
	var i ClossSession
	err := row.Scan(
		&i.RefreshToken,
		&i.AccessToken,
		&i.Username,
		&i.UserID,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :exec
;

UPDATE closs_session SET
    refresh_token = ?,
    access_token = ?
WHERE user_id = ?
`

type UpdateSessionParams struct {
	RefreshToken string
	AccessToken  string
	UserID       string
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) error {
	_, err := q.db.ExecContext(ctx, updateSession, arg.RefreshToken, arg.AccessToken, arg.UserID)
	return err
}
