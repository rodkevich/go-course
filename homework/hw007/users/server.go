package users

import (
	"context"
)

// GRPCServer ...
type GRPCServer struct {
	UnimplementedRegisterServer
	UnimplementedListServer
	db map[string]string
}

func (s *GRPCServer) Register(ctx context.Context, req *RegisterRequest) (resp *RegisterResponse, err error) {
	return &RegisterResponse{
		Id:   "1",
		Name: "1",
		Text: "1",
	}, nil
}

func (s *GRPCServer) List(ctx context.Context, req *ListRequest) (resp *ListResponse, err error) {

	return &ListResponse{
		Id:   "1",
		Name: "",
		Text: "",
	}, nil
}
