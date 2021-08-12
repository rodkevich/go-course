package types

import "fmt"

const (
	NoGroup        Group = ""
	Gopher         Group = "trainee"
	Pythonist      Group = "active"
	Sishneg        Group = "pending" // invalid if checked with IsValid()
	Javascriptizer Group = "blocked" // invalid if checked with IsValid()
)

// Group ...
type Group string

// IsValid ...
func (g *Group) IsValid() bool {
	switch *g {
	case Gopher, Pythonist, NoGroup:
		return true
	}
	return false
}

func (g Group) String() string {
	return string(g)
}

// CheckValidity ...
func (g *Group) CheckValidity(v interface{}) error {
	str, ok := v.(Group)
	if !ok {
		return fmt.Errorf("arg is not of Group type")
	}
	*g = str
	if !g.IsValid() {
		return fmt.Errorf("%s is not a valid Group", str)
	}
	return nil
}
