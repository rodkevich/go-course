package client

import (
	"context"
	"github.com/rodkevich/go-course/homework/hw007/api/v1/users"
	"google.golang.org/grpc"
)
const serverAddress = "127.0.0.1:9090"

var ctx = context.Background()

// Client ...
type Client struct {
	Conn *grpc.ClientConn
}

// NewClient for grpc server
func NewClient() *Client {
	opts := []grpc.DialOption{grpc.WithInsecure()} // disable tls
	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		panic(err)
	}
	return &Client{Conn: conn}
}

// Registration for new person in fake db
func (c *Client) Registration(req *users.RegistrationRequest) (string, error) {
	client := users.NewRegistrationClient(c.Conn)
	reg, err := client.Registration(ctx, req)
	if err != nil {
		panic(err)
	}
	return reg.GetMessage(), nil
}

// List all persons from fake db
func (c Client) List(req *users.ListRequest) (*users.ListResponse, error) {
	client := users.NewListClient(c.Conn)
	usersList, err := client.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return usersList, nil
}
