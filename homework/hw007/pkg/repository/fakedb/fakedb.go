package fakedb

import (
	"sync"

	repo "github.com/rodkevich/go-course/homework/hw007/pkg/repository"
)

// Db stores persons for app
type Db struct {
	Locker  sync.RWMutex
	Storage map[uint64]repo.User
}

// Lock ...
func (db *Db) Lock() {
	db.Locker.RLock()
}

// Unlock ...
func (db *Db) Unlock() {
	db.Locker.RUnlock()
}

// NewDb constructs new instance of fake-db
func NewDb() (*Db, error) {
	d := Db{
		Storage: map[uint64]repo.User{},
	}
	return &d, nil
}

// GetAllUsers return all existing users from storage
func (db *Db) GetAllUsers() map[uint64]repo.User {
	return db.Storage
}

// GetNewUserID ...
func (db *Db) GetNewUserID() (rtn uint64) {
	for i := range db.GetAllUsers() {
		if rtn < i {
			rtn = i
		}
	}
	rtn++
	return rtn
}
