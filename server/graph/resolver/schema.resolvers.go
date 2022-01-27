package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/graph/mapper"

	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.UserInput) (*generated.User, error) {
	user := generated.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Age:       input.Age,
		Password:  input.Password,
	}

	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*generated.User, error) {

	users, err := r.users.GetUsers()

	if err != nil {
		return nil, err
	}

	return mapper.Users(users), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
