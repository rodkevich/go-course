package repository

import "sync"

// User represents persons in app
type User struct {
	UserID     uint64 `json:"user_id"`
	UniqueName string `json:"unique_name"`
}

// Db stores persons for app
type Db struct {
	Locker sync.RWMutex
	Users  map[uint64]User
}

// NewDb constructs new instance of fake-db
func NewDb() (*Db, error) {
	d := Db{
		Users: map[uint64]User{},
	}
	return &d, nil
}

// GetNewUserID ...
func (db *Db) GetNewUserID() (rtn uint64) {
	for i := range db.Users {
		if rtn < i {
			rtn = i
		}
	}
	rtn++
	return rtn
}
