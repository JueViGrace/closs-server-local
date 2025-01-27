package types

import "github.com/google/uuid"

type Session struct {
	UserId       uuid.UUID
	Username     string
	RefreshToken string
    AccessToken string
}
