package types

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
