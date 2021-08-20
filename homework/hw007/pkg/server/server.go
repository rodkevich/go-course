package server

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/rodkevich/go-course/homework/hw007/api/v1/users"
	repo "github.com/rodkevich/go-course/homework/hw007/pkg/repository"
	"github.com/rodkevich/go-course/homework/hw007/pkg/repository/fakedb"
)

var errNotFound = errors.New("not found")
var errDuplicate = errors.New("name exists in DB")
var errInitFakeDb = errors.New("DB was not initialized")

// GRPCServer base type
type GRPCServer struct {
	users.UnimplementedRegistrationServer
	users.UnimplementedListServer
	db repo.Repository
}

// InitFakeDb method for initialising of new fake-db
func (s *GRPCServer) InitFakeDb() error {
	db, err := fakedb.NewDb()
	if err != nil {
		return errInitFakeDb
	}
	s.db = db
	return nil
}

// Registration for new person
func (s *GRPCServer) Registration(ctx context.Context, req *users.RegistrationRequest) (resp *users.RegistrationResponse, err error) {
	s.db.Lock()
	defer s.db.Unlock()

	for _, u := range s.db.AllUsers() {
		if strings.EqualFold(u.UniqueName, req.Name) {
			return nil, errDuplicate
		}
	}
	var newUserID = s.db.GetNewUserID()
	newUser := repo.User{
		UserID:     newUserID,
		UniqueName: req.Name,
	}
	s.db.AllUsers()[newUserID] = newUser
	return &users.RegistrationResponse{
		Id:      newUserID,
		Name:    req.Name,
		Message: "ok, " + req.Name + " - successfully registered",
	}, nil
}

// List existing persons
func (s *GRPCServer) List(context.Context, *users.ListRequest) (resp *users.ListResponse, err error) {
	s.db.Lock()
	defer s.db.Unlock()
	if len(s.db.AllUsers()) == 0 {
		return nil, errNotFound
	}
	var rtn []*users.ListResponse_User
	for _, user := range s.db.AllUsers() {
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
func (s *GRPCServer) Run(address string) (net.Listener, error) {
	return net.Listen("tcp", address)
}
