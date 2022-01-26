package graph

import (
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *Resolver {
	if DB == nil {
		panic("DB IS BROKEN")
	}

	return &Resolver{
		DB,
	}
}
