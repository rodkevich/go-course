package users

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var errNotFound = errors.New("not found")
var errDuplicate = errors.New("name exists in DB")

// User ..
type User struct {
	UserID     uint64 `json:"user_id,omitempty"`
	UniqueName string `json:"unique_name,omitempty"`
}

// Db ...
type Db struct {
	Users  map[uint64]User
	locker sync.RWMutex
}

// NewDb ...
func NewDb() (Db, error) {
	d := Db{
		Users: map[uint64]User{},
	}
	return d, nil
}

func (db *Db) getNewUserID() (rtn uint64) {
	for i := range db.Users {
		if rtn < i {
			rtn = i
		}
	}
	rtn++
	return rtn
}

// GRPCServer ...
type GRPCServer struct {
	UnimplementedRegisterServer
	UnimplementedListServer
	db *Db
}

// InitDb ...
func (s *GRPCServer) InitDb() {
	db, _ := NewDb()
	s.db = &db
}

// Register ...
func (s *GRPCServer) Register(ctx context.Context, req *RegisterRequest) (resp *RegisterResponse, err error) {
	s.db.locker.RLock()
	defer s.db.locker.RUnlock()
	for _, u := range s.db.Users {
		if strings.EqualFold(u.UniqueName, req.Name) {
			return nil, errDuplicate
		}
	}
	var newUserID = s.db.getNewUserID()
	newUser := User{
		UserID:     newUserID,
		UniqueName: req.Name,
	}
	s.db.Users[newUserID] = newUser
	return &RegisterResponse{
		Id:      newUserID,
		Name:    req.Name,
		Message: "ok, user registered",
	}, nil
}

// List ...
func (s *GRPCServer) List(ctx context.Context, req *ListRequest) (resp *ListResponse, err error) {
	s.db.locker.RLock()
	defer s.db.locker.RUnlock()
	if len(s.db.Users) == 0 {
		return nil, errNotFound
	}
	var rtn []*ListResponse_Result
	for _, user := range s.db.Users {
		rtn = append(rtn, &ListResponse_Result{
			Id:   user.UserID,
			Name: user.UniqueName,
		})
	}
	return &ListResponse{
		Results: rtn,
	}, nil
}
