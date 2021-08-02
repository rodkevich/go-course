package main

import (
	"fmt"
	"log"

	"github.com/rodkevich/go-course/homework/hw007/api/v1/users"
	"github.com/rodkevich/go-course/homework/hw007/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	startApp("127.0.0.1:9090")
}

func startApp(address string) {
	s := grpc.NewServer()

	grpcUsersServer := &server.GRPCServer{}
	// init fake database
	if err := grpcUsersServer.InitDb(); err != nil {
		log.Fatal(err)
	}
	// register services
	users.RegisterListServer(s, grpcUsersServer)
	users.RegisterRegistrationServer(s, grpcUsersServer)

	usersInstance, err := grpcUsersServer.Run(address)
	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	// serve grpc
	if err := s.Serve(usersInstance); err != nil {
		log.Fatal(err)
	}
}
