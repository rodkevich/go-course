// Repository interface
// provides methods to act with postgres person table

package repository

import "time"

// PersonModel represent the user model
type PersonModel struct {
	UserID      string       `json:"userID"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Description string       `json:"description"`
	Privileged  bool         `json:"privileged"`
	Location    *Coordinates `json:"location"`
}

// Coordinates ...
type Coordinates struct {
	Lon *float64 `json:"lon"`
	Lat *float64 `json:"lat"`
}

// Repository represent the repositories
type Repository interface {
	Up() error
	Close()
	Drop() error
	Truncate() error

	Create(user *PersonModel) (string, error)
	Update(user *PersonModel) error
	Delete(id string) error
	// Find look for something
	Find() ([]*PersonModel, error)
	FindByID(id string) (*PersonModel, error)
}
