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
	err                               error
	batch                             []*book.Contact
	pgBook, mongoBook                 book.ContactsBookDataSource
	contactPeterPan, contactPinocchio *book.Contact
)

func main() {
	os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/postgres")
	os.Setenv("MONGO_URL", "mongodb://localhost:27017/")

	// new books
	mongoBook, _ = mongodb.NewContactsBook()
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

	// create contactPeterPan with group
	contactPeterPan, err = book.NewContact(
		"Питер Джеймсович Пэн",
		"+91 (123) 456-7890",
		types.Gopher,
	)
	if err != nil {
		panic(err)
	}
	// contactPinocchio created with NO group
	// book.UnsafeEmptyContact() can be used for purpose
	// !WARNING UNSAFE! : both methods allow invalid fields to be set
	contactPinocchio = &book.Contact{
		UUID:  nil,
		Name:  "Пинок Карлович Кио",
		Phone: "123.456.7890",
		Group: "",
	}

	batch = append(batch, contactPeterPan, contactPinocchio)

	// create records in PG & Mongo
	for _, contact := range batch {
		// add batch to postgres + get string UUID
		pgUUID, err := pgBook.Create(contact)
		if err != nil {
			log.Println("can not Create contact in PG", contact)
			panic(err)
		}

		// convert from returned str to valid uuid type
		decodedUUID, err := uuid.Parse(pgUUID)
		if err != nil {
			log.Println("can not decode UUID from PG", pgUUID)
			panic(err)
		}

		// set uuid as a contact field
		contact.UUID = &decodedUUID
		// add contact to mongo using postgres generated UUID
		contactMongoID, err := mongoBook.Create(contact)
		if err != nil {
			log.Printf("can not create contact in mongo, UUID: %v", contact.UUID)
			panic(err)
		}
		// show in logs
		log.Printf("pg: contact created: %v", contact.UUID)
		log.Printf("mongo: contact created: %v", contactMongoID)
	}

	// Postgres search + update
	// find batch with no group - it will be `Пинок Карлович Кио`
	batch, err = pgBook.FindByGroup("")
	if err != nil {
		log.Println("pg: search failed:", err)
		panic(err)
	}
	// from pgFoundData update 1rst contact's `group` field
	contactPinocchio = *(pgBook.AssignContactToGroup(batch[0], types.Gopher))
	log.Printf("pg: updated contact: %v", contactPinocchio)

	// find both from batch by group
	batch, _ = pgBook.FindByGroup(types.Gopher)
	for _, contact := range batch {
		log.Printf("pg: found `Gopher` contact %v\n", contact.UUID)
	}

	// Mongo search + update
	// find batch with no group - it will be `Пинок Карлович Кио`
	batch, err = mongoBook.FindByGroup("")
	if err != nil {
		log.Println("mongo: search failed:", err)
		panic(err)
	}
	// from mongoFoundData update 1rst contact's `group` field
	contactPinocchio = *(mongoBook.AssignContactToGroup(batch[0], types.Gopher))
	log.Printf("mongo: updated contact: %v", contactPinocchio)

	// find both from batch by group
	batch, _ = mongoBook.FindByGroup(types.Gopher)
	for _, record := range batch {
		log.Printf("mongo: found `Gopher` contact %v\n", record.UUID)
	}

	// shutdownApp(pgBook, mongoBook)
}

func shutdownApp(pgBook book.ContactsBookDataSource, mongoBook book.ContactsBookDataSource) {
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
