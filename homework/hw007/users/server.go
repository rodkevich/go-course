package users

import (
	"context"
	"errors"
	"github.com/rodkevich/go-course/homework/hw007/internal/constants"
	"net"
	"strings"
	"sync"
)

var errNotFound = errors.New("not found")
var errDuplicate = errors.New("name exists in DB")

// User represents persons in app
type User struct {
	UserID     uint64 `json:"user_id"`
	UniqueName string `json:"unique_name"`
}

// Db stores persons for app
type Db struct {
	locker sync.RWMutex
	Users  map[uint64]User
}

// newDb constructs new instance of fake-db
func newDb() (*Db, error) {
	d := Db{
		Users: map[uint64]User{},
	}
	return &d, nil
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

// GRPCServer base type
type GRPCServer struct {
	Address string
	UnimplementedRegistrationServer
	UnimplementedListServer
	db *Db
}

// InitDb method for initialising of new fake-db
func (s *GRPCServer) InitDb() error {
	db, _ := newDb()
	s.db = db
	s.Address = constants.ServerAddress
	return nil
}

// Registration for new person
func (s *GRPCServer) Registration(_ context.Context, req *RegistrationRequest) (resp *RegistrationResponse, err error) {
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
	return &RegistrationResponse{
		Id:      newUserID,
		Name:    req.Name,
		Message: "ok, " + req.Name + " - successfully registered",
	}, nil
}

// List existing persons
func (s *GRPCServer) List(context.Context, *ListRequest) (resp *ListResponse, err error) {
	s.db.locker.RLock()
	defer s.db.locker.RUnlock()
	if len(s.db.Users) == 0 {
		return nil, errNotFound
	}
	var rtn []*ListResponse_User
	for _, user := range s.db.Users {
		rtn = append(rtn, &ListResponse_User{
			Id:   user.UserID,
			Name: user.UniqueName,
		})
	}
	return &ListResponse{
		Users: rtn,
	}, nil
}

// Run start a server
func (s *GRPCServer) Run() (net.Listener, error) {
	return net.Listen("tcp", constants.ServerAddress)
}
