package types

import (
	"time"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Code      string    `json:"code"`
	LastSync  time.Time `json:"last_sync"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PickerInfoResponse struct {
	Name    string `json:"name"`
	Almacen string `json:"almacen"`
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
