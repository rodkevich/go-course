package book

import (
	"github.com/google/uuid"
	"github.com/rodkevich/go-course/homework/hw009/book/types"
)

// Contact represents any persons model
type Contact struct {
	ID    *uuid.UUID  `json:"id"`
	Name  string      `json:"first_name"`
	Phone types.Phone `json:"phone"`
	Group types.Group `json:"group"`
}

// EmptyContact ....
func EmptyContact() *Contact {
	return &Contact{}
}

// NewContact ...
func NewContact(name string, phone types.Phone, group types.Group) (*Contact, error) {
	err := phone.CheckValidity(phone)
	if err != nil {
		return nil, err
	}
	err = group.CheckValidity(group)
	if err != nil {
		return nil, err
	}
	return &Contact{ID: nil, Name: name, Phone: phone, Group: group}, nil
}
