package data

import (
	"context"
	"errors"

	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/google/uuid"
)

type SessionStorage interface {
	GetSessionByKey(key uuid.UUID) (session *types.Session, err error)
	CreateSession(id uuid.UUID, r *types.Session) (err error)
	UpdateSession(id uuid.UUID, r *types.Session) (err error)
	DeleteSession(key uuid.UUID) (err error)
}

func (s *cacheStorage) SessionStorage() SessionStorage {
	return NewSessionStorage(s.ctx)
}

type sessionStorage struct {
	ctx   context.Context
	store types.SessionStore
}

func NewSessionStorage(ctx context.Context) SessionStorage {
	store := make(types.SessionStore)
	return &sessionStorage{
		ctx:   ctx,
		store: store,
	}
}

func (s *sessionStorage) GetSessionByKey(key uuid.UUID) (*types.Session, error) {
	v, ok := s.store[key]
	if !ok {
		return nil, errors.New("this entry doesn't exists")
	}

	session := &types.Session{
		Username:     v.Username,
		RefreshToken: v.RefreshToken,
	}

	return session, nil
}

func (s *sessionStorage) CreateSession(id uuid.UUID, r *types.Session) error {
	_, ok := s.store[id]
	if ok {
		return errors.New("This entry already exists")
	}

	s.store[id] = types.Session{
		Username:     r.Username,
		RefreshToken: r.RefreshToken,
	}

	return nil
}

func (s *sessionStorage) UpdateSession(id uuid.UUID, r *types.Session) error {
	_, ok := s.store[id]
	if !ok {
		return errors.New("this entry doesn't exists")
	}

	s.store[id] = types.Session{
		Username:     r.Username,
		RefreshToken: r.RefreshToken,
	}

	return nil
}

func (s *sessionStorage) DeleteSession(key uuid.UUID) error {
	_, ok := s.store[key]
	if !ok {
		return errors.New("this entry doesn't exists")
	}

	delete(s.store, key)

	return nil
}
