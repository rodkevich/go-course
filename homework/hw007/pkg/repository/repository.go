package repository

import "sync"

// User represents persons in app
type User struct {
	UserID     uint64 `json:"user_id"`
	UniqueName string `json:"unique_name"`
}

// Repository represents the repositories for usage
type Repository interface {
	sync.Locker
	GetAllUsers() map[uint64]User
	GetNewUserID() (rtn uint64)
}
