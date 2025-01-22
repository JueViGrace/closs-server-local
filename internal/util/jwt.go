package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	jwtSecret string           = os.Getenv("JWT_SECRET")
	Issuer    string           = "ClossServer"
	Audience  jwt.ClaimStrings = jwt.ClaimStrings{
		"api",
	}
	accessExpiration  time.Time = time.Now().Add(1 * time.Hour)
	refreshExpiration time.Time = time.Now().Add(24 * time.Hour)
)

type userClaims struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Code     string    `json:"code"`
}

type JWTClaims struct {
	userClaims
	jwt.RegisteredClaims
}

func CreateAccessToken(id, username, code string) (string, error) {
	return createJWT(id, username, code, accessExpiration)
}

func CreateRefreshToken(id, username, code string) (string, error) {
	return createJWT(id, username, code, refreshExpiration)
}

func createJWT(id, username, code string, expiration time.Time) (string, error) {
	tokenId, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	claims := JWTClaims{
		userClaims{
			UserId:   userId,
			Username: username,
			Code:     code,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    Issuer,
			Subject:   username,
			ID:        tokenId.String(),
			Audience:  Audience,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	claims := new(JWTClaims)

	return jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
}
