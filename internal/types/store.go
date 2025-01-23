package types

import "github.com/google/uuid"

type SessionStore = map[uuid.UUID]Session

type Session struct {
	UserId       uuid.UUID
	Username     string
	RefreshToken string
}
