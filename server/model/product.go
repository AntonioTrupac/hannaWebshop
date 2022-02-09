package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	price       float64
	name        string
	description string
	stock       int
	Image       *Image
	Tag         []*Tag
}
