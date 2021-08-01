package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rodkevich/go-course/homework/hw007/users"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &users.GRPCServer{}
	if err := srv.InitDb(); err != nil {
		log.Fatal(err)
	}
	users.RegisterRegisterServer(s, srv)
	users.RegisterListServer(s, srv)
	listener, err := net.Listen("tcp", users.ServerAddress)

	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}

	// go func() {
	// 	if err := s.Serve(listener); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	//
	// cli := users.NewListClient()
	// }
}
