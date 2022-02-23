package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/graph/mapper"
	"github.com/AntonioTrupac/hannaWebshop/graph/types"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.UserInput) (*generated.User, error) {
	user := types.ModelUser(ctx, input)

	err := r.users.CreateAUser(user)

	if err != nil {
		return nil, err
	}

	return mapper.GeneratedUser(user), nil
}

func (r *mutationResolver) CreateProducts(ctx context.Context, input generated.ProductInput) (*generated.Product, error) {
	product := types.ModelProducts(ctx, input)

	err := r.products.CreateAProduct(product)

	if err != nil {
		return nil, err
	}

	return types.GeneratedProduct(product), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*generated.User, error) {
	users, err := r.users.GetUsers()

	if err != nil {
		return nil, err
	}

	return mapper.Users(users), nil
}

func (r *queryResolver) GetProducts(ctx context.Context) ([]*generated.Product, error) {
	products, err := r.products.GetProducts()

	if err != nil {
		return nil, err
	}

	return types.GetProductsFromDb(products), nil
}

func (r *queryResolver) GetProductByID(ctx context.Context, id int) (*generated.Product, error) {
	product, err := r.products.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return types.GetProductFromDb(product), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
