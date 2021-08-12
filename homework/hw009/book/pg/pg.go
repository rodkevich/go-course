package pg

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	cb "github.com/rodkevich/go-course/homework/hw009/book"
	"github.com/rodkevich/go-course/homework/hw009/book/types"
)

// Represents the contactsBook model
type contactsBook struct {
	db *pgxpool.Pool
}

var (
	stmt              string
	rows              pgx.Rows
	row               pgx.Row
	ctxDefault        = context.Background()
	operationsTimeOut = 3 * time.Second
)

// Up prepares database
func (b *contactsBook) Up() (err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	// create PG extension to generate UUID's
	stmt = `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		`
	rows, err = b.db.Query(ctx, stmt)
	if err != nil {
		log.Println("pg: error on creating for uuid generation extension")
		return
	}
	defer rows.Close()
	// create table for contacts
	stmt = `
		CREATE TABLE IF NOT EXISTS contact (
		contact_id uuid DEFAULT uuid_generate_v4 (),
		contact_group VARCHAR NOT NULL,
		contact_name VARCHAR NOT NULL,
		contact_phone VARCHAR,
		PRIMARY KEY (contact_id));
		`
	rows, err = b.db.Query(ctx, stmt)
	if err != nil {
		log.Printf("pg: error: create tables: %v", err)
		return
	}
	log.Println("pg: book UP operation done")
	return
}

// Close the connection
func (b *contactsBook) Close() {
	log.Println("pg: book disconnecting ...")
	defer log.Println("pg: book disconnecting - done")
	go b.db.Close()
}

// Drop ...
func (b *contactsBook) Drop() (err error) {
	stmt = `
		DROP TABLE IF EXISTS contact;
		`
	// statement to send to db
	rows, err = b.db.Query(ctxDefault, stmt)
	if err != nil {
		log.Printf("pg: error: cbd.Drop(): %v", err)
		return
	}
	defer rows.Close()
	log.Println("pg: database dropped")
	return
}

// Truncate ...
func (b *contactsBook) Truncate() (err error) {
	stmt = `
		TRUNCATE TABLE contact;
		`
	rows, err = b.db.Query(ctxDefault, stmt)
	if err != nil {
		log.Printf("pg: error: cbd.Truncate(): %v", err)
		return
	}
	defer rows.Close()
	return
}

// Create ...
func (b *contactsBook) Create(contact *types.Contact) (contactID string, err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	stmt = `
		INSERT INTO contact (contact_group, contact_name, contact_phone)
		VALUES ($1, $2, $3)
		RETURNING contact_id;
		`
	err = b.db.QueryRow(
		ctx, stmt,
		contact.Group,
		contact.Name,
		contact.Phone).Scan(&contactID)
	if err != nil {
		return
	}
	return
}

// AssignContactToGroup ...
func (b *contactsBook) AssignContactToGroup(contact *types.Contact, group types.Group) (newContact *types.Contact) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	stmt := `
		UPDATE contact
		SET contact_group = $1
		WHERE contact_id = $2
		RETURNING *;
		`
	newContact = new(types.Contact)
	err := b.db.QueryRow(ctx, stmt, group, contact.UUID).Scan(
		&newContact.UUID,
		&newContact.Name,
		&newContact.Group,
		&newContact.Phone,
	)
	if err != nil {
		log.Printf("pg: error: db.AssignContactToGroup(): %v", err)
		return
	}
	return
}

// FindByGroup ...
func (b *contactsBook) FindByGroup(group types.Group) (contacts []*types.Contact, err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	stmt := `
		SELECT contact_id, contact_name, contact_group, contact_phone
		FROM contact
		WHERE contact_group = $1
		`
	rows, err = b.db.Query(ctx, stmt, group)
	if err != nil {
		log.Printf("pg: find by group: stmt: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		c := new(types.Contact)
		err = rows.Scan(
			&c.UUID,
			&c.Name,
			&c.Group,
			&c.Phone,
		)
		if err != nil {
			log.Printf("pg: find by group: stmt.Next(): %v\n", err)
			return
		}
		contacts = append(contacts, c)
	}
	return
}

// NewContactsBook ...
func NewContactsBook() (ds cb.ContactBookDataSource, err error) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	var config = os.Getenv("DATABASE_URL")
	// create data-base connection pool:
	pool, err := pgxpool.Connect(ctx, config)
	if err != nil {
		log.Printf("pg: unable to connection to database: %v\n", err)
		return
	}
	log.Printf("pg: connected!")
	ds = &contactsBook{pool}
	return
}
