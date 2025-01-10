package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Code      string `json:"codigo"`
	Role      Role   `json:"-"`
	LastSync  string `json:"lastSync"`
	Version   string `json:"version"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Code     string `json:"codigo"`
	Role     Role   `json:"-"`
	LastSync string `json:"lastSync"`
	Version  string `json:"version"`
}

type UpdateLastSyncRequest struct {
	LastSync time.Time `json:"lastSync"`
	Version  string    `json:"version"`
	ID       uuid.UUID `json:"id"`
}

type UpdatePasswordRequest struct {
	Password string    `json:"password"`
	ID       uuid.UUID `json:"id"`
}

type Role string

func (r Role) String() string {
	return string(r)
}

const (
	RoleCustomer Role = "customer"
	RoleSalesman Role = "salesman"
	RoleManager  Role = "manager"
	RoleAdmin    Role = "admin"
)

func DbUserToUser(db *db.ClossUser) *UserResponse {
	return &UserResponse{
		ID:        db.ID,
		Username:  db.Username,
		Password:  db.Password,
		Role:      Role(db.Role),
		LastSync:  db.UltSinc,
		Version:   db.Version,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		DeletedAt: db.DeletedAt.String,
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
		Role:      RoleSalesman.String(),
		UltSinc:   r.LastSync,
		Version:   r.Version,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}, nil
}
