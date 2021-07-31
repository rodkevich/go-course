package users

import (
	"context"
)

// User ..
type User struct {
	UserID     uint64 `json:"user_id,omitempty"`
	UniqueName string `json:"unique_name,omitempty"`
}

// Db ...
type Db struct {
	Users map[uint64]User
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

	var newUserID uint64 = s.db.getNewUserID()
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
	return
}
