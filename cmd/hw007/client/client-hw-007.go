package main

import (
	"log"

	"github.com/rodkevich/go-course/homework/hw007/client"
	"github.com/rodkevich/go-course/homework/hw007/users"
)

func main() {
	// create new client & reg some users
	cl := client.NewClient()

	persons := [3]string{"name 1", "name 2", "name 3"}
	for _, user := range persons {
		resp, err := cl.Registration(&users.RegistrationRequest{Name: user})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
	}
	// call List of users
	if personsList, err := cl.List(&users.ListRequest{}); err != nil {
		log.Fatal(err)
	} else {
		log.Println(personsList)
	}
}
