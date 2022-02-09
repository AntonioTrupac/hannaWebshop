package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	url       string
	productId int
}
