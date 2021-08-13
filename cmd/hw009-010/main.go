package main

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/rodkevich/go-course/homework/hw009/book"
	"github.com/rodkevich/go-course/homework/hw009/book/pg"
	"github.com/rodkevich/go-course/homework/hw009/book/types"
	"github.com/rodkevich/go-course/homework/hw010/mongodb"
)

var (
	uuID                string
	err                 error
	peterPan, pinocchio *types.Contact
	batchContacts       []*types.Contact
	pgBook, mongoBook   book.ContactBookDataSource
	booksAvailable      []book.ContactBookDataSource
)

func init() {
	// setup env
	err = os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// init books
	mongoBook, err = mongodb.NewContactsBook()
	if err != nil {
		panic(err)
	}
	pgBook, err = pg.NewContactsBook()
	if err != nil {
		panic(err)
	}

	booksAvailable = []book.ContactBookDataSource{pgBook, mongoBook}
}

func main() {
	// schedule app turn off
	defer cleanOnAppShutdown(false)

	// prepare the contacts:
	// create contacts with/without group and uuid
	peterPan, err = types.NewContact("Питер Джеймсович Пэн", "+91 (123) 456-7890")
	if err != nil {
		panic(err)
	}
	// unsafe method
	pinocchio = &types.Contact{
		UUID:  nil,
		Name:  "Пинок Карлович Кио",
		Phone: "123.456.7890",
		Group: types.Gopher,
	}

	invalid := &types.Contact{
		UUID:  nil,
		Name:  "Не будет создан",
		Phone: "111.222.3333",
		Group: types.Javascriptizer,
	}

	// validate then add to batch
	var contactsToValidate = []*types.Contact{peterPan, pinocchio, invalid}
	for _, contact := range contactsToValidate {
		var hasErrors bool
		err = contact.Phone.CheckValidity(contact.Phone)
		if err != nil {
			log.Printf("warn: %v invalid. reason: %v\n", contact.Name, contact.Phone)
			hasErrors = true
		}
		err = contact.Group.CheckValidity(contact.Group)
		if err != nil {
			log.Printf("warn: %v invalid. reason: %v\n", contact.Name, contact.Group)
			hasErrors = true
		}
		// skip corrupted contacts
		if hasErrors == false {
			batchContacts = append(batchContacts, contact)
		}
	}

	// setup data-sources
	for _, dataSource := range booksAvailable {
		if err = dataSource.Up(); err != nil {
			panic(err)
		}
	}

	// create records from batch in both data-sources
	for _, contact := range batchContacts {
		// add to postgres to get generated uuid
		uuID, err = pgBook.Create(contact)
		if err != nil {
			panic(err)
		}

		// convert from returned str to valid uuid type
		decodedUUID, err := uuid.Parse(uuID)
		if err != nil {
			panic(err)
		}

		// set uuid to a contact field
		contact.UUID = &decodedUUID
		// add contact to mongo
		uuID, err = mongoBook.Create(contact)
		if err != nil {
			panic(err)
		}
		// show in logs
		log.Printf("pg: contact created: %v", contact.UUID)
		log.Printf("mongo: contact created: %v", uuID)
	}

	// find + update operations
	for _, dataSource := range booksAvailable {
		log.Printf("using data-source: %v", dataSource.String())

		// find batchContacts with no group - it will be `Питер Джеймсович Пэн`
		batchContacts, err = dataSource.FindByGroup(types.NoGroup) // "" can be used
		toBeUpdatedContact := batchContacts[0]
		log.Printf("wanna update `NoGroup` contact %v\n", toBeUpdatedContact)
		if err != nil {
			panic(err)
		}

		// from found update 1-rst contact's `group` field
		newPeterPan := dataSource.AssignContactToGroup(toBeUpdatedContact, types.Gopher)
		if newPeterPan.UUID.String() != peterPan.UUID.String() {
			panic("OMG it's not Peter")
		}
		log.Printf("updated contact: %v", newPeterPan)

		// find & print both contacts to ensure: now they are in a same group
		batchContacts, err = dataSource.FindByGroup(types.Gopher)
		if err != nil {
			panic(err)
		}
		for _, contact := range batchContacts {
			log.Printf("found `Gopher` contact %v\n", contact)
		}
	}
}

func cleanOnAppShutdown(deleteAll bool) {
	if deleteAll {
		for _, dataSource := range booksAvailable {
			// delete records
			err := dataSource.Truncate()
			if err != nil {
				panic(err)
			}
			// drop databases
			err = dataSource.Drop()
			if err != nil {
				panic(err)
			}
			// close connections
			dataSource.Close()
		}
		return
	}
	for _, dataSource := range booksAvailable {
		dataSource.Close()
	}
	return
}
