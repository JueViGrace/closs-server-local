package data

import (
	"context"

	database "github.com/JueViGrace/closs-server-local/internal/database/sqlite"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/google/uuid"
)

type SessionStorage interface {
	GetSessionById(id uuid.UUID) (session *types.Session, err error)
	GetSessionByUsername(username string) (session *types.Session, err error)
	CreateSession(r *types.Session) (err error)
	UpdateSession(r *types.Session) (err error)
	DeleteSession(id uuid.UUID) (err error)
}

func (s *cacheStore) SessionStorage() SessionStorage {
	return NewSessionStorage(s.ctx, s.queries)
}

type sessionStorage struct {
	ctx context.Context
	db  *database.Queries
}

func NewSessionStorage(ctx context.Context, queries *database.Queries) SessionStorage {
	return &sessionStorage{
		ctx: ctx,
		db:  queries,
	}
}

func (s *sessionStorage) GetSessionById(id uuid.UUID) (*types.Session, error) {
	session, err := s.db.GetSessionById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(session.UserID)
	if err != nil {
		return nil, err
	}

	return &types.Session{
		UserId:       userId,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		AccessToken:  session.AccessToken,
	}, nil
}

func (s *sessionStorage) GetSessionByUsername(username string) (*types.Session, error) {
	session, err := s.db.GetSessionByUsername(s.ctx, username)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(session.UserID)
	if err != nil {
		return nil, err
	}

	return &types.Session{
		UserId:       userId,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		AccessToken:  session.AccessToken,
	}, nil
}

func (s *sessionStorage) CreateSession(r *types.Session) error {
	err := s.db.CreateSession(s.ctx, database.CreateSessionParams{
		RefreshToken: r.RefreshToken,
		AccessToken:  r.AccessToken,
		Username:     r.Username,
		UserID:       r.UserId.String(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionStorage) UpdateSession(r *types.Session) error {
	err := s.db.UpdateSession(s.ctx, database.UpdateSessionParams{
		RefreshToken: r.RefreshToken,
		AccessToken:  r.AccessToken,
		UserID:       r.UserId.String(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionStorage) DeleteSession(id uuid.UUID) error {
	err := s.db.DeleteSessionById(s.ctx, id.String())
	if err != nil {
		return err
	}
	return nil
}
