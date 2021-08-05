package repository

import "fmt"

// Contact represent the user model
type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"first_name"`
	Phone string `json:"phone"`
	Group Group  `json:"group"`
}

// Group ...
type Group string

const (
	Gopher         Group = "trainee"
	Pythonist      Group = "active"
	Sishnik        Group = "pending"
	Javascriptizer Group = "blocked"
)

// AllUserGroups ...
var AllUserGroups = []Group{
	Gopher,
	Pythonist,
	Sishnik,
	Javascriptizer,
}

func (g Group) IsValid() bool {
	switch g {
	case Gopher, Pythonist:
		return true
	}
	return false
}

func (g Group) String() string {
	return string(g)
}

// CheckValidity ...
func (g *Group) CheckValidity(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*g = Group(str)
	if !g.IsValid() {
		return fmt.Errorf("%s is not a valid Group", str)
	}
	return nil
}

// ContactsBook represent the repositories
type ContactsBook interface {
	Up() error
	Close()
	Drop() error
	Truncate() error

	Create(contact *Contact) (string, error)
	UpdateContactGroup(contact *Contact) error
	Find() ([]*Contact, error)
	FindByGroup(Group string) (*Contact, error)
}
