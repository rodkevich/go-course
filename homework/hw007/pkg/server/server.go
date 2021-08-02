package server

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/rodkevich/go-course/homework/hw007/api/v1/users"
	"github.com/rodkevich/go-course/homework/hw007/internal/constants"
	"github.com/rodkevich/go-course/homework/hw007/pkg/repository"
	"github.com/rodkevich/go-course/homework/hw007/pkg/repository/fakedb"
)

var errNotFound = errors.New("not found")
var errDuplicate = errors.New("name exists in DB")

// GRPCServer base type
type GRPCServer struct {
	Address string
	users.UnimplementedRegistrationServer
	users.UnimplementedListServer
	db *fakedb.Db
}

// InitDb method for initialising of new fake-db
func (s *GRPCServer) InitDb() error {
	db, _ := fakedb.NewDb()
	s.db = db
	s.Address = constants.ServerAddress
	return nil
}

// Registration for new person
func (s *GRPCServer) Registration(_ context.Context, req *users.RegistrationRequest) (resp *users.RegistrationResponse, err error) {
	s.db.Locker.RLock()
	defer s.db.Locker.RUnlock()
	for _, u := range s.db.Users {
		if strings.EqualFold(u.UniqueName, req.Name) {
			return nil, errDuplicate
		}
	}
	var newUserID = s.db.GetNewUserID()
	newUser := repository.User{
		UserID:     newUserID,
		UniqueName: req.Name,
	}
	s.db.Users[newUserID] = newUser
	return &users.RegistrationResponse{
		Id:      newUserID,
		Name:    req.Name,
		Message: "ok, " + req.Name + " - successfully registered",
	}, nil
}

// List existing persons
func (s *GRPCServer) List(context.Context, *users.ListRequest) (resp *users.ListResponse, err error) {
	s.db.Locker.RLock()
	defer s.db.Locker.RUnlock()
	if len(s.db.Users) == 0 {
		return nil, errNotFound
	}
	var rtn []*users.ListResponse_User
	for _, user := range s.db.Users {
		rtn = append(rtn, &users.ListResponse_User{
			Id:   user.UserID,
			Name: user.UniqueName,
		})
	}
	return &users.ListResponse{
		Users: rtn,
	}, nil
}

// Run start a server
func (s *GRPCServer) Run() (net.Listener, error) {
	return net.Listen("tcp", constants.ServerAddress)
}
