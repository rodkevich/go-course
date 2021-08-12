package book

import "github.com/rodkevich/go-course/homework/hw009/book/types"

// ContactBookDataSource represent the repositories for Contact-Books
type ContactBookDataSource interface {
	Up() error
	Close()
	Drop() error
	Truncate() error

	Create(contact *types.Contact) (string, error)
	AssignContactToGroup(contact *types.Contact, group types.Group) (n *types.Contact)
	FindByGroup(group types.Group) ([]*types.Contact, error)
	String() string
}
