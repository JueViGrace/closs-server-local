package types

import (
	"time"

	"github.com/JueViGrace/closs-server-local/internal/database"
)

type UserResponse struct {
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Code      string    `json:"codigo"`
	Nombre    string    `json:"nombre"`
	Almacen   string    `json:"almacen"`
	Desactivo bool      `json:"-"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func DbUserToUser(db *database.KeWusuario) *UserResponse {
	return &UserResponse{
		Username:  db.Username,
		Password:  db.PasswordApp,
		Code:      "",
		Nombre:    db.Nombre,
		Almacen:   db.Almacen,
		Desactivo: db.Desactivo,
		UpdatedAt: db.Fechamodifi,
	}
}
