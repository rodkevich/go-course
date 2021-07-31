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
	srv.InitDb()
	users.RegisterRegisterServer(s, srv)
	users.RegisterListServer(s, srv)
	listener, err := net.Listen("tcp", ":9090")

	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
