package main

import (
	"context"
	"github.com/rodkevich/go-course/homework/hw007/users"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// flag.Parse()
	// if flag.NArg() < 1 {
	// 	log.Fatal("not enough args")
	// }
	conn, err := grpc.Dial(users.ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := users.NewRegisterClient(conn)
	broReg, err := c.Register(context.Background(), &users.RegisterRequest{Name: "Bro"})
	broReg2, err := c.Register(context.Background(), &users.RegisterRequest{Name: "Bro2"})
	broReg3, err := c.Register(context.Background(), &users.RegisterRequest{Name: "Bro3"})
	broReg4, err := c.Register(context.Background(), &users.RegisterRequest{Name: "Bro4"})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(broReg.GetMessage())
	log.Println(broReg2.GetMessage())
	log.Println(broReg3.GetMessage())
	log.Println(broReg4.GetMessage())

	l := users.NewListClient(conn)
	broList, err := l.List(context.Background(), &users.ListRequest{})
	log.Println("ok")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(broList.Results)

}
