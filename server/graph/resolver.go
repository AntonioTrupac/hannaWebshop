package graph

import (
	"github.com/AntonioTrupac/hannaWebshop/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	users []*model.User
	DB    *gorm.DB
}