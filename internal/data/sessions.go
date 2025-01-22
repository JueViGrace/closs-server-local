package data

import (
	"context"
	"errors"

	"github.com/JueViGrace/closs-server-local/internal/types"
)

type SessionStorage interface {
	GetSessionByKey(key string) (session *types.Session, err error)
	CreateSession(r *types.Session) (err error)
	UpdateSession(r *types.Session) (err error)
	DeleteSession(key string) (err error)
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

func (s *sessionStorage) GetSessionByKey(key string) (*types.Session, error) {
	v, ok := s.store[key]
	if !ok {
		return nil, errors.New("this entry doesn't exists")
	}

	session := &types.Session{
		Username:     key,
		RefreshToken: v,
	}

	return session, nil
}

func (s *sessionStorage) CreateSession(r *types.Session) error {
	_, ok := s.store[r.Username]
	if ok {
		return errors.New("This entry already exists")
	}

	s.store[r.Username] = r.RefreshToken

	return nil
}

func (s *sessionStorage) UpdateSession(r *types.Session) error {
	_, ok := s.store[r.Username]
	if !ok {
		return errors.New("this entry doesn't exists")
	}

	s.store[r.Username] = r.RefreshToken

	return nil
}

func (s *sessionStorage) DeleteSession(key string) error {
	_, ok := s.store[key]
	if !ok {
		return errors.New("this entry doesn't exists")
	}

	delete(s.store, key)

	return nil
}
