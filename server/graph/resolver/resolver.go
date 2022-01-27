package resolver

import (
	"github.com/AntonioTrupac/hannaWebshop/service/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users users.Service
}

func NewResolver(users users.Service) *Resolver {
	return &Resolver{
		users: users,
	}
}
