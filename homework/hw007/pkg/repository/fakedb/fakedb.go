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

// Lock to prevent changing Storage
func (db *Db) Lock() {
	db.Locker.RLock()
}

// Unlock to allow changing Storage
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

// AllUsers return all existing users from storage
func (db *Db) AllUsers() map[uint64]repo.User {
	return db.Storage  // to skip Save-like methods for now
}

// GetNewUserID ...
func (db *Db) GetNewUserID() (rtn uint64) {
	for i := range db.AllUsers() {
		if rtn < i {
			rtn = i
		}
	}
	rtn++
	return rtn
}
