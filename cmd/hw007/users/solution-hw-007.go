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
	users.RegisterRegisterServer(s, srv)
	users.RegisterListServer(s, srv)
	l, err := net.Listen("tcp", ":9090")

	err = initDataBase(err)

	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func initDataBase(err error) error {
	return nil
}
