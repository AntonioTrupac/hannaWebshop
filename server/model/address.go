package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	AddressLine string
	City        string
	PostalCode  int
	Country     string
	UserId      int
}

func (a *Address) TableName() string {
	return "addresses"
}
