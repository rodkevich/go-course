package book

import "github.com/rodkevich/go-course/homework/hw009/book/types"

// ContactsBookDataSource represent the repositories
type ContactsBookDataSource interface {
	Up() error
	Close()
	Drop() error
	Truncate() error

	Create(contact *Contact) (string, error)
	AssignContactToGroup(contact *Contact, group types.Group) (n *Contact)
	FindByGroup(group types.Group) ([]*Contact, error)
}
