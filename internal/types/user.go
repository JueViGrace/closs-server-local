package types

import (
	"time"

	"github.com/JueViGrace/closs-server-local/internal/database"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Code      string    `json:"codigo"`
	LastSync  time.Time `json:"lastSync"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func DbUserToUser(id uuid.UUID, db *database.KeWusuario) *UserResponse {
	return &UserResponse{
		ID:        id,
		Username:  db.Username,
		Password:  db.PasswordApp,
		Code:      "",
		LastSync:  time.Now(),
		Version:   "",
		CreatedAt: db.Fechanac,
		UpdatedAt: db.Fechamodifi,
	}
}
