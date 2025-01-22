package types

import "github.com/google/uuid"

type SessionStore = map[uuid.UUID]Session

type Session struct {
	Username     string
	RefreshToken string
}
