package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUsers() (users []types.UserResponse, err error)
	GetUserById(id *uuid.UUID) (user *types.UserResponse, err error)
	GetUserByUsername(username string) (user *types.UserResponse, err error)
	CreateUser(r *types.CreateUserRequest) (user *types.UserResponse, err error)
	UpdateLastSync(r *types.UpdateLastSyncRequest) (err error)
	UpdatePassword(r *types.UpdatePasswordRequest) (err error)
	SoftDeleteUser(id *uuid.UUID) (err error)
	DeleteUser(id *uuid.UUID) (err error)
}

func (s *storage) UserStore() UserStore {
	return NewUserStore(s.ctx, s.queries)
}

type userStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewUserStore(ctx context.Context, db *db.Queries) UserStore {
	return &userStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *userStore) GetUsers() ([]types.UserResponse, error) {
	users := make([]types.UserResponse, 0)

	dbUsers, err := s.db.GetUsers(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbUser := range dbUsers {
		users = append(users, *types.DbUserToUser(&dbUser))
	}

	return users, nil
}

func (s *userStore) GetUserById(id *uuid.UUID) (*types.UserResponse, error) {
	dbUser, err := s.db.GetUserById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	return types.DbUserToUser(&dbUser), nil
}

func (s *userStore) GetUserByUsername(username string) (*types.UserResponse, error) {
	dbUser, err := s.db.GetUserById(s.ctx, username)
	if err != nil {
		return nil, err
	}

	return types.DbUserToUser(&dbUser), nil
}

func (s *userStore) CreateUser(r *types.CreateUserRequest) (*types.UserResponse, error) {
	cr, err := types.CreateUserToDb(r)
	if err != nil {
		return nil, err
	}

	dbUser, err := s.db.CreateUser(s.ctx, *cr)
	if err != nil {
		return nil, err
	}

	return types.DbUserToUser(&dbUser), nil
}

func (s *userStore) UpdatePassword(r *types.UpdatePasswordRequest) error {
	err := s.db.UpdatePassword(s.ctx, db.UpdatePasswordParams{
		Password:  r.Password,
		UpdatedAt: time.Now().String(),
		ID:        r.ID.String(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *userStore) UpdateLastSync(r *types.UpdateLastSyncRequest) error {
	err := s.db.UpdateLastSync(s.ctx, db.UpdateLastSyncParams{
		UltSinc:   r.LastSync.String(),
		Version:   r.Version,
		UpdatedAt: time.Now().String(),
		ID:        r.ID.String(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *userStore) SoftDeleteUser(id *uuid.UUID) error {
	err := s.db.SoftDeleteUser(s.ctx, db.SoftDeleteUserParams{
		UpdatedAt: time.Now().String(),
		DeletedAt: sql.NullString{
			String: time.Now().String(),
			Valid:  true,
		},
		ID: id.String(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *userStore) DeleteUser(id *uuid.UUID) error {
	err := s.db.DeleteUserById(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}
