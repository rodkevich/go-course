package client

import (
	"context"
	"github.com/rodkevich/go-course/homework/hw007/users"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"log"
)

func register(c *cli.Context) error {
	opts := []grpc.DialOption{grpc.WithInsecure()} // disable tls
	conn, err := grpc.Dial(c.String(users.ServerAddress), opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := users.NewListClient(conn)
	stream, err := client.List(context.Background(), &users.ListRequest{}, nil)
	if err != nil {
		panic(err)
	}
	results := stream.GetResults()
	log.Println(results)

	return nil
}
