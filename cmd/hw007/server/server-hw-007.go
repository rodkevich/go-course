package main

import (
	"fmt"
	"log"

	"github.com/rodkevich/go-course/homework/hw007/server"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	grpcUsersServer := &server.GRPCServer{}
	// init fake database
	if err := grpcUsersServer.InitDb(); err != nil {
		log.Fatal(err)
	}
	// register services
	server.RegisterListServer(s, grpcUsersServer)
	server.RegisterRegistrationServer(s, grpcUsersServer)

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
