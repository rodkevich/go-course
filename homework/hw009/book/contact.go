package book

import (
	"github.com/google/uuid"
	"github.com/rodkevich/go-course/homework/hw009/book/types"
)

// Contact represents any persons model
// 	@param UUID  *uuid.UUID
// 	@param Name  string
// 	@param Phone types.Phone
// 	@param Group types.Group
type Contact struct {
	UUID  *uuid.UUID  `json:"uuid" bson:"uuid"`
	Name  string      `json:"name" bson:"name"`
	Phone types.Phone `json:"phone" bson:"phone"`
	Group types.Group `json:"group" bson:"group"`
}

// UnsafeEmptyContact ....
func UnsafeEmptyContact() *Contact {
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
	return &Contact{UUID: nil, Name: name, Phone: phone, Group: group}, nil
}
