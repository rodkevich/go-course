package repository

// User represents persons in app
type User struct {
	UserID     uint64 `json:"user_id"`
	UniqueName string `json:"unique_name"`
}

// // Repository represents the repositories for usage
// type Repository interface {
// 	GetNewUserID() uint64
// 	GetAllUsers() map[uint64]User
// }
