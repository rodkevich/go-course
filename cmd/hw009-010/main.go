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
	PeterPan, Pinocchio *types.Contact
)

func main() {
	os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/postgres")
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")

	// new books
	mongoBook, err = mongodb.NewContactsBook()
	if err != nil {
		panic(err)
	}
	pgBook, err = pg.NewContactsBook()
	if err != nil {
		panic(err)
	}
	// init pgBook + mongoBook
	if err = mongoBook.Up(); err != nil {
		panic(err)
	}
	if err = pgBook.Up(); err != nil {
		panic(err)
	}
	// create PeterPan with group
	PeterPan, err = types.NewContact("Питер Джеймсович Пэн", "+91 (123) 456-7890")
	if err != nil {
		panic(err)
	}
	// Pinocchio created with NO group and nil uuID
	Pinocchio = &types.Contact{
		UUID:  nil,
		Name:  "Пинок Карлович Кио",
		Phone: "123.456.7890",
		Group: types.NoGroup,
	}

	batch = append(batch, PeterPan, Pinocchio)

	// create records in PG & Mongo
	for _, contact := range batch {
		// add batch to postgres + get string uuID
		uuID, err = pgBook.Create(contact)
		if err != nil {
			panic(err)
		}

		// convert from returned str to valid uuid type
		decodedUUID, err := uuid.Parse(uuID)
		if err != nil {
			panic(err)
		}

		// set uuid as a contact field
		contact.UUID = &decodedUUID
		// add contact to mongo using postgres generated uuID
		uuID, err = mongoBook.Create(contact)
		if err != nil {
			panic(err)
		}
		// show in logs
		log.Printf("pg: contact created: %v", contact.UUID)
		log.Printf("mongo: contact created: %v", uuID)
	}

	// Postgres search + update
	// find batch with no group - it will be `Пинок Карлович Кио`
	batch, err = pgBook.FindByGroup(types.NoGroup) // "" can be used
	if err != nil {
		panic(err)
	}
	// from pgFoundData update 1rst contact's `group` field
	Pinocchio = pgBook.AssignContactToGroup(batch[0], types.Gopher)
	log.Printf("pg: updated contact: %v", Pinocchio)

	// find both from batch by group
	batch, _ = pgBook.FindByGroup(types.Gopher)
	for _, contact := range batch {
		log.Printf("pg: found `Gopher` contact %v\n", contact.UUID)
	}

	// Mongo search + update
	// find batch with no group - it will be `Пинок Карлович Кио`
	batch, err = mongoBook.FindByGroup(types.NoGroup)
	if err != nil {
		panic(err)
	}
	// from mongoFoundData update 1rst contact's `group` field
	Pinocchio = mongoBook.AssignContactToGroup(batch[0], types.Gopher)
	log.Printf("mongo: updated contact: %v", Pinocchio)

	// find both from batch by group
	batch, _ = mongoBook.FindByGroup(types.Gopher)
	for _, record := range batch {
		log.Printf("mongo: found `Gopher` contact %v\n", record.UUID)
	}

	shutdownApp(pgBook, mongoBook)
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
