package types

import (
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthDataHandler = func(*fiber.Ctx, *AuthData) error

type AuthData struct {
	UserId   uuid.UUID
	Username string
	Code     string
}

type JwtData struct {
	Token  *jwt.Token
	Claims util.JWTClaims
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	Token string `json:"refresh_token" validate:"required"`
}

type RecoverPasswordResquest struct {
	Username string `json:"username" validate:"required"`
}
