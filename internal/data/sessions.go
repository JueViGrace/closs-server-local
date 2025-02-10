package data

import (
	"context"

	database "github.com/JueViGrace/closs-server-local/internal/database/sqlite"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/google/uuid"
)

type SessionStore interface {
	GetSessionById(id uuid.UUID) (session *types.Session, err error)
	GetSessionByUsername(username string) (session *types.Session, err error)
	CreateSession(r *types.Session) (err error)
	UpdateSession(r *types.Session) (err error)
	DeleteSession(id uuid.UUID) (err error)
	DeleteSessionByToken(token string) (err error)
}

func (s *storage) SessionStorage() SessionStore {
	return NewSessionStorage(s.ctx, s.cacheStore.queries)
}

type sessionStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewSessionStorage(ctx context.Context, queries *database.Queries) SessionStore {
	return &sessionStore{
		ctx: ctx,
		db:  queries,
	}
}

func (s *sessionStore) GetSessionById(id uuid.UUID) (*types.Session, error) {
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

func (s *sessionStore) GetSessionByUsername(username string) (*types.Session, error) {
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

func (s *sessionStore) CreateSession(r *types.Session) error {
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

func (s *sessionStore) UpdateSession(r *types.Session) error {
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

func (s *sessionStore) DeleteSession(id uuid.UUID) error {
	err := s.db.DeleteSessionById(s.ctx, id.String())
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) DeleteSessionByToken(token string) error {
	err := s.db.DeleteSessionByToken(s.ctx, database.DeleteSessionByTokenParams{
		RefreshToken: token,
		AccessToken:  token,
	})
	if err != nil {
		return err
	}
	return nil
}
