package main

import (
	"fmt"
	"github.com/rodkevich/go-course/homework/hw007/users"
	"log"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	grpcUsersServer := &users.GRPCServer{}
	// init fake database
	if err := grpcUsersServer.InitDb(); err != nil {
		log.Fatal(err)
	}
	// register services
	users.RegisterListServer(s, grpcUsersServer)
	users.RegisterRegistrationServer(s, grpcUsersServer)

	usersInstance, err := grpcUsersServer.Run()
	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	// serve grpc
	if err := s.Serve(usersInstance); err != nil {
		log.Fatal(err)
	}
}

