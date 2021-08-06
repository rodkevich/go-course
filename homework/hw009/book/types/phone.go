package types

import (
	"fmt"
	"regexp"
)

// Phone ...
type Phone string

// IsValid ..
// 	123-456-7890
// 	(123) 456-7890
// 	123 456 7890
// 	123.456.7890
// 	+91 (123) 456-7890
func (p *Phone) IsValid() bool {
	re := regexp.MustCompile(`(^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$)`)
	// submatch := re.FindStringSubmatch("Phone Number: 15817452367;")
	submatch := re.FindStringSubmatch(p.String())
	if len(submatch) < 2 {
		return false
	}
	// match := submatch[1]
	return true
}

func (p Phone) String() string {
	return string(p)
}

// CheckValidity ...
func (p *Phone) CheckValidity(v interface{}) error {
	str, ok := v.(Phone)
	if !ok {
		return fmt.Errorf("phone is of invalid type")
	}

	*p = Phone(str)
	if !p.IsValid() {
		return fmt.Errorf("%s is not a valid Phone", str)
	}
	return nil
}
