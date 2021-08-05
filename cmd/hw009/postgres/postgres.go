package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/pgxpool"
	repo "github.com/rodkevich/go-course/homework/hw009/repository"
)

// Represents the contactBook model
type contactBook struct {
	db *pgxpool.Pool
}

// Up prepares database
func (r *contactBook) Up() error {
	ctx := context.Background()
	extension := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	_, err := r.db.Query(ctx, extension)
	if err != nil {
		return err
	}
	// create table for persons
	query := `CREATE TABLE IF NOT EXISTS person (
              person_id uuid DEFAULT uuid_generate_v4 (),
              group VARCHAR NOT NULL,
              name VARCHAR NOT NULL,
              phone VARCHAR,
              PRIMARY KEY (person_id)
          )`

	_, err = r.db.Query(ctx, query)
	fmt.Println("Query was made")

	if err != nil {
		return err
	}
	return nil
}

func (r *contactBook) Close() {
	panic("implement me")
}

func (r *contactBook) Drop() error {
	panic("implement me")
}

func (r *contactBook) Truncate() error {
	panic("implement me")
}

func (r *contactBook) Create(contact *repo.Contact) (string, error) {
	panic("implement me")
}

func (r *contactBook) UpdateContactGroup(contact *repo.Contact) error {
	panic("implement me")
}

func (r *contactBook) Find() ([]*repo.Contact, error) {
	panic("implement me")
}

func (r *contactBook) FindByGroup(GroupID string) (*repo.Contact, error) {
	panic("implement me")
}

// NewRepository will create a variable that represent the Repository struct
func NewRepository() (repo.ContactsBook, error) {
	var config = os.Getenv("DATABASE_URL")
	pool, poolErr := pgxpool.Connect(context.Background(), config)
	if poolErr != nil {
		log.Fatalf("Unable to connection to database: %v\n", poolErr)
	}
	log.Printf("Connected!")
	return &contactBook{pool}, nil
}
