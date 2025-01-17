package types

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Code      string `json:"codigo"`
	Version   string `json:"version"`
	UpdatedAt string `json:"updatedAt"`
}

// func DbUserToUser(db *db.ClossUser) *UserResponse {
// 	return &UserResponse{
// 		ID:        db.ID,
// 		Username:  db.Username,
// 		Password:  db.Password,
// 		Version:   db.Version,
// 		UpdatedAt: db.UpdatedAt,
// 	}
// }
