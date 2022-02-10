package resolver

import (
	productsService "github.com/AntonioTrupac/hannaWebshop/service/products"
	usersService "github.com/AntonioTrupac/hannaWebshop/service/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users    usersService.UserService
	products productsService.ProductService
}

func NewResolver(users usersService.UserService, products productsService.ProductService) *Resolver {
	return &Resolver{
		users:    users,
		products: products,
	}
}
