package client

import (
	"context"
	"time"

	"github.com/rodkevich/go-course/homework/hw007/api/v1/users"
	"google.golang.org/grpc"
)

var ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)

// Client ...
type Client struct {
	Conn *grpc.ClientConn
}

// NewClient for grpc server
func NewClient(address string) *Client {
	opts := []grpc.DialOption{grpc.WithInsecure()} // disable tls
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		panic(err)
	}
	return &Client{Conn: conn}
}

// Registration for new person in fake db
func (c *Client) Registration(req *users.RegistrationRequest) (string, error) {
	defer cancel()

	client := users.NewRegistrationClient(c.Conn)
	reg, err := client.Registration(ctx, req)
	if err != nil {
		panic(err)
	}
	return reg.GetMessage(), nil
}

// List all persons from fake db
func (c Client) List(req *users.ListRequest) (*users.ListResponse, error) {
	defer cancel()

	client := users.NewListClient(c.Conn)
	usersList, err := client.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return usersList, nil
}
