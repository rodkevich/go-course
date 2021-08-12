package types

import "github.com/google/uuid"

// Contact represents any persons model
// 	@param UUID  *uuid.UUID
// 	@param Name  string
// 	@param Phone types.Phone
// 	@param Group types.Group
type Contact struct {
	UUID  *uuid.UUID `json:"uuid" bson:"uuid"`
	Group Group      `json:"group" bson:"group"`
	Name  string     `json:"name" bson:"name"`
	Phone Phone      `json:"phone" bson:"phone"`
}

// NewContact ...
func NewContact(name string, phone Phone) (c *Contact, err error) {
	err = phone.CheckValidity(phone)
	if err != nil {
		return
	}
	c = &Contact{UUID: nil, Name: name, Phone: phone, Group: NoGroup}
	return
}
