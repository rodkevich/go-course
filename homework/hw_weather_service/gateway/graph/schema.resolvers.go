package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rodkevich/go-course/homework/hw_weather_service/gateway/graph/generated"
	"github.com/rodkevich/go-course/homework/hw_weather_service/gateway/graph/model"
	"github.com/rodkevich/go-course/homework/hw_weather_service/history"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	hs := history.Record{
		TraceID:   "",
		LogString: nil,
	}
	println(hs)
	return nil, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, changes map[string]interface{}) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCityWeatherByName(ctx context.Context, name string, country *string, config *model.ConfigInput) (*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
