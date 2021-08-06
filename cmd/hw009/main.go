package main

import (
	"log"
	"os"

	"github.com/rodkevich/go-course/homework/hw009/book"
	"github.com/rodkevich/go-course/homework/hw009/book/pg"
	"github.com/rodkevich/go-course/homework/hw009/book/types"
)

func main() {
	os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@0.0.0.0:5432/postgres")
	var (
		cbd book.ContactsBookDataSource
		err error
	)
	cbd, err = pg.NewContactsBook()
	if err != nil {
		panic(err)
	}
	// init
	if err = cbd.Up(); err != nil {
		panic(err)
	}
	// create PeterPan with group
	PeterPan, err := book.NewContact(
		"PeterPan",
		"123-456-5678",
		types.Gopher,
	)
	if err != nil {
		panic(err)
	}
	// create Pinocchio with NO group
	Pinocchio := book.EmptyContact()
	Pinocchio.Name = "Pinocchio"
	Pinocchio.Phone = "231-456-1234"

	// create records in DB
	var contacts []*book.Contact
	contacts = append(contacts, PeterPan, Pinocchio)
	for _, person := range contacts {
		id, _ := cbd.Create(person)
		log.Printf("User created: %v", id)
	}
	// find contacts with no group - it will be `Pinocchio`
	findNoGroup, _ := cbd.FindByGroup("")
	PinocchioFromDB := findNoGroup[0]

	// update its group
	newCont := cbd.AssignContactToGroup(PinocchioFromDB, types.Gopher)
	log.Printf("Updated contact: %v", newCont)

	// find both from contacts by group
	allByGroup, _ := cbd.FindByGroup(types.Gopher)
	for _, contact := range allByGroup {
		log.Printf("Found `Gopher` contact %v\n", contact.ID.String())

	}
	// drop database
	err = cbd.Drop()
	if err != nil {
		return
	}
}
