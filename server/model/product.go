package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Price       float64
	Name        string
	Description string
	Rating      float64
	Stock       int
	Image       []*Image
	Category    []*Category
}
