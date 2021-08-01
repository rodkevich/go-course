package client

import (
	"context"
	"github.com/rodkevich/go-course/homework/hw007/internal/constants"
	"github.com/rodkevich/go-course/homework/hw007/server"

	"google.golang.org/grpc"
)

var ctx = context.Background()

// Client ...
type Client struct {
	Conn *grpc.ClientConn
}

// NewClient for grpc server
func NewClient() *Client {
	opts := []grpc.DialOption{grpc.WithInsecure()} // disable tls
	conn, err := grpc.Dial(constants.ServerAddress, opts...)
	if err != nil {
		panic(err)
	}
	return &Client{Conn: conn}
}

// Registration for new person in fake db
func (c *Client) Registration(req *server.RegistrationRequest) (string, error) {
	client := server.NewRegistrationClient(c.Conn)
	reg, err := client.Registration(ctx, req)
	if err != nil {
		panic(err)
	}
	return reg.GetMessage(), nil
}

// List all persons from fake db
func (c Client) List(req *server.ListRequest) (*server.ListResponse, error) {
	client := server.NewListClient(c.Conn)
	usersList, err := client.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return usersList, nil
}
