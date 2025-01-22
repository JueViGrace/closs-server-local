package types

type SessionStore = map[string]string

type Session struct {
	Username     string
	RefreshToken string
}
