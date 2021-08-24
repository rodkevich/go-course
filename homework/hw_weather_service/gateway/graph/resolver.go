package graph

import "github.com/rodkevich/go-course/homework/hw_weather_service/gateway/graph/model"

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	users  []*model.User
	cities []*model.City
}
