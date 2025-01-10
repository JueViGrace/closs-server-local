package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
)

type SessionStore interface {
	GetSessionById(id string) (*db.ClossSession, error)
	GetSessionByToken(token string) (*db.ClossSession, error)
	DeleteSessionById(id string) error
	DeleteSessionByToken(token string) error
}

func (s *storage) SessionStore() SessionStore {
	return NewSessionStore(s.ctx, s.queries)
}

type sessionStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewSessionStore(ctx context.Context, db *db.Queries) SessionStore {
	return &sessionStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *sessionStore) GetSessionById(id string) (*db.ClossSession, error) {
	session, err := s.db.GetSessionById(s.ctx, id)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *sessionStore) GetSessionByToken(token string) (*db.ClossSession, error) {
	session, err := s.db.GetSessionByToken(s.ctx, token)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *sessionStore) DeleteSessionById(id string) error {
	err := s.db.DeleteSessionById(s.ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) DeleteSessionByToken(token string) error {
	err := s.db.DeleteSessionByToken(s.ctx, token)
	if err != nil {
		return err
	}
	return nil
}
