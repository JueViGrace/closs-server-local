package data

import (
	"context"
	"errors"

	mydb "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	cache "github.com/JueViGrace/closs-server-local/internal/database/sqlite"
	database "github.com/JueViGrace/closs-server-local/internal/database/sqlite"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/google/uuid"
)

type AuthStore interface {
	SignIn(r *types.SignInRequest) (res *types.AuthResponse, err error)
	Refresh(r *types.RefreshRequest, a *types.AuthData) (res *types.AuthResponse, err error)
	ForgotPassword() (err error)
}

func (s *storage) AuthStore() AuthStore {
	return NewAuthStore(s.ctx, s.myStore.Queries, s.cacheStore.queries)
}

type authStore struct {
	ctx   context.Context
	db    *mydb.Queries
	cache *cache.Queries
}

func NewAuthStore(ctx context.Context, db *mydb.Queries, cache *cache.Queries) AuthStore {
	return &authStore{
		ctx:   ctx,
		db:    db,
		cache: cache,
	}
}

func (s *authStore) SignIn(r *types.SignInRequest) (*types.AuthResponse, error) {
	dbSession, err := s.cache.GetSessionByUsername(s.ctx, r.Username)
	if err == nil {
		err = s.cache.DeleteSessionById(s.ctx, dbSession.UserID)
		if err != nil {
			return nil, err
		}
	}

	dbUser, err := s.db.GetUserByUsername(s.ctx, r.Username)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	user := types.DbUserToUser(id, &dbUser)

	if !util.ValidatePassword(r.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := createTokens(user)
	if err != nil {
		return nil, err
	}

	err = s.cache.CreateSession(s.ctx, database.CreateSessionParams{
		UserID:       user.ID.String(),
		Username:     user.Username,
		RefreshToken: token.RefreshToken,
		AccessToken:  token.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (s *authStore) Refresh(r *types.RefreshRequest, a *types.AuthData) (*types.AuthResponse, error) {
	dbUser, err := s.db.GetUserByUsername(s.ctx, a.Username)
	if err != nil {
		return nil, err
	}

	user := types.DbUserToUser(a.UserId, &dbUser)

	newTokens, err := createTokens(user)
	if err != nil {
		return nil, err
	}

	err = s.cache.UpdateSession(s.ctx, database.UpdateSessionParams{
		RefreshToken: newTokens.RefreshToken,
		AccessToken:  newTokens.AccessToken,
		UserID:       a.UserId.String(),
	})
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  newTokens.AccessToken,
		RefreshToken: newTokens.RefreshToken,
	}, nil
}

func (s *authStore) ForgotPassword() error {
	return nil
}

func createTokens(user *types.UserResponse) (*types.AuthResponse, error) {
	accessToken, err := util.CreateAccessToken(user.ID.String(), user.Username, user.Code)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.CreateRefreshToken(user.ID.String(), user.Username, user.Code)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
