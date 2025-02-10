package data

import (
	"context"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	"github.com/JueViGrace/closs-server-local/internal/types"
)

type UserStore interface {
	GetUsers() ([]types.UserResponse, error)
	GetDbUserByUsername(username string) (user *database.KeWusuario, err error)
	GetUser(a *types.AuthData) (user *types.UserResponse, err error)
	GetUserInfo(a *types.AuthData) (picker *types.PickerInfoResponse, err error)
}

func (s *storage) UserStore() UserStore {
	return NewUserStore(s.ctx, s.myStore.Queries)
}

type userStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewUserStore(ctx context.Context, db *database.Queries) UserStore {
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
		users = append(users, *types.DbUserToUser([16]byte{}, &dbUser))
	}

	return users, nil
}

func (s *userStore) GetUser(a *types.AuthData) (*types.UserResponse, error) {
	user := new(types.UserResponse)

	dbUser, err := s.GetDbUserByUsername(a.Username)
	if err != nil {
		return nil, err
	}

	user = types.DbUserToUser(a.UserId, dbUser)

	return user, nil
}

func (s *userStore) GetUserInfo(a *types.AuthData) (*types.PickerInfoResponse, error) {
	info := new(types.PickerInfoResponse)

	dbUser, err := s.GetDbUserByUsername(a.Username)
	if err != nil {
		return nil, err
	}

	info = &types.PickerInfoResponse{
		Name:    dbUser.Nombre,
		Almacen: dbUser.Almacen,
	}

	return info, nil
}

func (s *userStore) GetDbUserByUsername(username string) (*database.KeWusuario, error) {
	dbUser, err := s.db.GetUserByUsername(s.ctx, username)
	if err != nil {
		return nil, err
	}

	return &dbUser, nil
}
