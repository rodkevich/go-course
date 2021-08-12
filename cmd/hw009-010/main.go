package main

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/rodkevich/go-course/homework/hw010/mongodb"

	"github.com/rodkevich/go-course/homework/hw009/book"
	"github.com/rodkevich/go-course/homework/hw009/book/pg"
	"github.com/rodkevich/go-course/homework/hw009/book/types"
)

var (
	uuID                string
	err                 error
	batch               []*types.Contact
	pgBook, mongoBook   book.ContactBookDataSource
	peterPan, pinocchio *types.Contact
)

func main() {
	os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/postgres")
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")

	// new books
	mongoBook, err = mongodb.NewContactsBook()
	if err != nil {
		log.Printf("err: %v", err)
	}
	pgBook, err = pg.NewContactsBook()
	if err != nil {
		log.Printf("err: %v", err)
	}

	// init pgBook + mongoBook
	if err = mongoBook.Up(); err != nil {
		log.Printf("err: %v", err)
	}
	if err = pgBook.Up(); err != nil {
		log.Printf("err: %v", err)
	}

	// schedule app turn off
	defer shutdownApp(pgBook, mongoBook)

	// create peterPan with group
	peterPan, err = types.NewContact("Питер Джеймсович Пэн", "+91 (123) 456-7890")
	if err != nil {
		log.Printf("err: %v", err)
	}
	// pinocchio created with NO group and nil uuID
	pinocchio = &types.Contact{
		UUID:  nil,
		Name:  "Пинок Карлович Кио",
		Phone: "123.456.7890",
		Group: types.NoGroup,
	}

	batch = append(batch, peterPan, pinocchio)

	// create records in PG & Mongo
	for _, contact := range batch {
		// add batch to postgres + get string uuID
		uuID, err = pgBook.Create(contact)
		if err != nil {
			log.Printf("err: %v", err)
		}

		// convert from returned str to valid uuid type
		decodedUUID, err := uuid.Parse(uuID)
		if err != nil {
			log.Printf("err: %v", err)
		}

		// set uuid as a contact field
		contact.UUID = &decodedUUID
		// add contact to mongo using postgres generated uuID
		uuID, err = mongoBook.Create(contact)
		if err != nil {
			log.Printf("err: %v", err)
		}
		// show in logs
		log.Printf("pg: contact created: %v", contact.UUID)
		log.Printf("mongo: contact created: %v", uuID)
	}

	// Postgres search + update
	// find batch with no group - it will be `Пинок Карлович Кио`
	batch, err = pgBook.FindByGroup(types.NoGroup) // "" can be used
	if err != nil {
		log.Printf("err: %v", err)
	}
	// from pgFoundData update 1rst contact's `group` field
	pinocchio = pgBook.AssignContactToGroup(batch[0], types.Gopher)
	log.Printf("pg: updated contact: %v", pinocchio)

	// find both from batch by group
	batch, _ = pgBook.FindByGroup(types.Gopher)
	for _, contact := range batch {
		log.Printf("pg: found `Gopher` contact %v\n", contact.UUID)
	}

	// Mongo search + update
	// find batch with no group - it will be `Пинок Карлович Кио`
	batch, err = mongoBook.FindByGroup(types.NoGroup)
	if err != nil {
		log.Printf("err: %v", err)
	}
	// from mongoFoundData update 1rst contact's `group` field
	pinocchio = mongoBook.AssignContactToGroup(batch[0], types.Gopher)
	log.Printf("mongo: updated contact: %v", pinocchio)

	// find both from batch by group
	batch, _ = mongoBook.FindByGroup(types.Gopher)
	for _, record := range batch {
		log.Printf("mongo: found `Gopher` contact %v\n", record.UUID)
	}

}

func shutdownApp(pgBook book.ContactBookDataSource, mongoBook book.ContactBookDataSource) {
	// // delete records
	// err := pgBook.Truncate()
	// if err != nil {
	// 	return
	// }
	// err = mongoBook.Truncate()
	// if err != nil {
	// 	return
	// }

	// drop databases
	err = pgBook.Drop()
	if err != nil {
		return
	}
	err = mongoBook.Drop()
	if err != nil {
		return
	}
	// close connections
	pgBook.Close()
	mongoBook.Close()
}
