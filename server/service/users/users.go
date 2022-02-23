package service

import (
	"github.com/AntonioTrupac/hannaWebshop/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService interface {
	GetUsers() ([]*model.User, error)
	CreateAUser(input *model.User) error
}

type users struct {
	DB *gorm.DB
}

func NewUsers(db *gorm.DB) UserService {
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

func (u *users) CreateAUser(input *model.User) error {
	return u.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(input).Error; err != nil {
			return err
		}

		for _, value := range input.Address {
			value.UserId = int(input.ID)
		}

		if err := tx.CreateInBatches(input.Address, 100).Error; err != nil {
			return err
		}

		return nil
	})
}
