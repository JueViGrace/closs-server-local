package types

import (
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

type AuthData struct {
	Jwt JwtData
}

type JwtData struct {
	Token  *jwt.Token
	Claims util.JWTClaims
}
