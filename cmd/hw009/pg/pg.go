package pg

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rodkevich/go-course/homework/hw009/book/types"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
	cb "github.com/rodkevich/go-course/homework/hw009/book"
)

// Represents the contactsBook model
type contactsBook struct {
	db *pgxpool.Pool
}

// Up prepares database
func (b *contactsBook) Up() error {
	var (
		ctx          = context.Background()
		strExtension = "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
		err          error
	)
	// create str_extension for uuid generation by db
	if _, err = b.db.Query(ctx, strExtension); err != nil {
		log.Println("error on creating str_extension for uuid generation")
		return err
	}
	// create table for persons
	str := `CREATE TABLE IF NOT EXISTS contact (
              contact_id uuid DEFAULT uuid_generate_v4 (),
              contact_group VARCHAR NOT NULL,
              contact_name VARCHAR NOT NULL,
              contact_phone VARCHAR,
              PRIMARY KEY (contact_id)
          )`

	_, err = b.db.Query(ctx, str)
	if err != nil {
		log.Println("CREATE TABLE")
		return err
	}
	log.Println("UP operation done")
	return nil
}

// Close the connection
func (b *contactsBook) Close() {
	b.db.Close()
}

// Drop ...
func (b *contactsBook) Drop() error {
	var (
		ctx  = context.Background()
		str  = "DROP TABLE IF EXISTS contact"
		stmt pgx.Rows
		err  error
	)
	// statement to send to db
	stmt, err = b.db.Query(ctx, str)
	if err != nil {
		log.Printf("error :cbd.Drop(): %v", err)
		return err
	}
	log.Println("Database dropped")
	defer stmt.Close()
	return nil
}

// Truncate ...
func (b *contactsBook) Truncate() error {
	ctx := context.Background()
	str := "TRUNCATE TABLE contact;"
	stmt, err := b.db.Query(ctx, str)
	if err != nil {
		log.Printf("error :cbd.Truncate(): %v", err)
		return err
	}
	defer stmt.Close()
	return nil
}

// Create ...
func (b *contactsBook) Create(contact *cb.Contact) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var contactId string
	stmt := "INSERT INTO contact (" +
		"contact_group, " +
		"contact_name, " +
		"contact_phone" +
		") VALUES ($1, $2, $3) " +
		"returning contact_id;"
	_ = b.db.QueryRow(
		ctx, stmt,
		contact.Group,
		contact.Name,
		contact.Phone,
	).Scan(&contactId)
	return contactId, nil
}

// AssignContactToGroup ...
func (b *contactsBook) AssignContactToGroup(contact *cb.Contact, group types.Group) (n *cb.Contact) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	n = new(cb.Contact)
	stmt := "UPDATE contact " +
		"SET contact_group = $1 " +
		"WHERE contact_id = $2 " +
		"RETURNING *;"

	err := b.db.QueryRow(ctx, stmt, group, contact.ID).Scan(
		&n.ID,
		&n.Name,
		&n.Group,
		&n.Phone,
	)
	if err != nil {
		log.Printf("error :db.AssignContactToGroup(): %v", err)
		return nil
	}
	return n
}

// FindByGroup ...
func (b *contactsBook) FindByGroup(group types.Group) ([]*cb.Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	persons := make([]*cb.Contact, 0)
	query := "SELECT " +
		"contact_id, " +
		"contact_name, " +
		"contact_group, " +
		"contact_phone " +
		"FROM contact WHERE contact_group = $1"

	rows, err := b.db.Query(ctx, query, group.String())
	if err != nil {
		log.Printf("find by group :rows: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		// c := cb.EmptyContact()
		c := new(cb.Contact)
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.Group,
			&c.Phone,
		)
		if err != nil {
			log.Printf("find by group :rows.Next(): %v\n", err)
			return nil, err
		}
		persons = append(persons, c)
	}
	return persons, nil
}

// NewContactsBook ...
func NewContactsBook() (cb.ContactsBookDataSource, error) {
	var config = os.Getenv("DATABASE_URL")
	// create data-base connection pool:
	pool, poolErr := pgxpool.Connect(context.Background(), config)
	if poolErr != nil {
		log.Fatalf("Unable to connection to database: %v\n", poolErr)
	}
	log.Printf("Connected!")
	return &contactsBook{pool}, nil
}
