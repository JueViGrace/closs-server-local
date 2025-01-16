package types

import (
	"time"

	"github.com/JueViGrace/closs-server-local/internal/db"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Code      string `json:"codigo"`
	Version   string `json:"version"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"codigo"`
	LastSync string `json:"lastSync"`
	Version  string `json:"version"`
}

type UpdatePasswordRequest struct {
	Password string    `json:"password"`
	ID       uuid.UUID `json:"id"`
}

func DbUserToUser(db *db.ClossUser) *UserResponse {
	return &UserResponse{
		ID:        db.ID,
		Username:  db.Username,
		Password:  db.Password,
		Version:   db.Version,
		UpdatedAt: db.UpdatedAt,
	}
}

func CreateUserToDb(r *CreateUserRequest) (*db.CreateUserParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &db.CreateUserParams{
		ID:        id.String(),
		Username:  r.Username,
		Password:  r.Password,
		Codigo:    r.Code,
		UltSinc:   r.LastSync,
		Version:   r.Version,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}, nil
}
