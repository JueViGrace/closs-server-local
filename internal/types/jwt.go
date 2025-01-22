package types

import (
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthData struct {
	UserId   uuid.UUID
	Username string
	Code     string
}

type JwtData struct {
	Token  *jwt.Token
	Claims util.JWTClaims
}
