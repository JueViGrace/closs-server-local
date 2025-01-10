package data

import (
	"context"
	"errors"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/JueViGrace/clo-backend/internal/util"
)

type AuthStore interface {
	SignIn(r *types.SignInRequest) (*types.AuthResponse, error)
	Refresh(r *types.RefreshRequest) (*types.AuthResponse, error)
	RecoverPassword(r *types.RecoverPasswordResquest) (*types.AuthResponse, error)
}

func (s *storage) AuthStore() AuthStore {
	return NewAuthStore(s.ctx, s.queries)
}

type authStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewAuthStore(ctx context.Context, db *db.Queries) AuthStore {
	return &authStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *authStore) SignIn(r *types.SignInRequest) (*types.AuthResponse, error) {
	user, err := s.db.GetUserByUsername(s.ctx, r.Username)
	if err != nil {
		return nil, err
	}

	if user.DeletedAt.Valid {
		return nil, errors.New("this user was deleted")
	}

	if !util.ValidatePassword(r.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	res, err := createTokens(&user)
	if err != nil {
		return nil, err
	}

	_, err = s.db.CreateSession(s.ctx, db.CreateSessionParams{
		UserID: user.ID,
		Token:  res.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authStore) Refresh(r *types.RefreshRequest) (*types.AuthResponse, error) {
	token, err := util.ValidateJWT(r.Token)
	if err != nil {
		s.db.DeleteSessionByToken(s.ctx, r.Token)
		return nil, err
	}

	claims, ok := token.Claims.(util.JWTClaims)
	if !ok {
		return nil, errors.New("invalid request")
	}

	user, err := s.db.GetUserById(s.ctx, claims.ID)
	if err != nil {
		return nil, err
	}

	res, err := createTokens(&user)
	if err != nil {
		return nil, err
	}

	_, err = s.db.CreateSession(s.ctx, db.CreateSessionParams{
		UserID: user.ID,
		Token:  res.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authStore) RecoverPassword(r *types.RecoverPasswordResquest) (*types.AuthResponse, error) {
	return nil, nil
}

func createTokens(user *db.ClossUser) (*types.AuthResponse, error) {
	accessToken, err := util.CreateAccessToken(user.ID, user.Username, user.Codigo)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.CreateRefreshToken(user.ID, user.Username, user.Codigo)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
