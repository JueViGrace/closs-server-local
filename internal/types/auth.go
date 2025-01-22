package types

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	Token string `json:"refreshToken" validate:"required"`
}

type RecoverPasswordResquest struct {
	Username string `json:"username" validate:"required"`
}
