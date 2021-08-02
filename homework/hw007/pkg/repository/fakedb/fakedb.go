package fakedb

import (
	"github.com/rodkevich/go-course/homework/hw007/pkg/repository"
	"sync"
)


// Db stores persons for app
type Db struct {
	Locker sync.RWMutex
	Users  map[uint64]repository.User
}

// NewDb constructs new instance of fake-db
func NewDb() (*Db, error) {
	d := Db{
		Users: map[uint64]repository.User{},
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
