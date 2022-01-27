package users

import (
	"github.com/AntonioTrupac/hannaWebshop/model"
	"gorm.io/gorm"
)

type Service interface {
	GetUsers() ([]*model.User, error)
}

type users struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Service {
	return &users{
		DB: db,
	}
}

func (u *users) GetUsers() ([]*model.User, error) {
	var users []*model.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
		//fmt.Errorf GOOGLE ME CHUCK
	}

	return users, nil
}
