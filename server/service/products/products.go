package service

import (
	"github.com/AntonioTrupac/hannaWebshop/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductService interface {
	GetProducts() ([]*model.Product, error)
	GetProductById(id int) (*model.Product, error)
	CreateAProduct(input *model.Product) error
}

type products struct {
	DB *gorm.DB
}

func NewProducts(db *gorm.DB) ProductService {
	return &products{
		DB: db,
	}
}

func (p products) GetProducts() ([]*model.Product, error) {
	var products []*model.Product

	if err := p.DB.Preload("Category").Preload("Image").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p products) GetProductById(id int) (*model.Product, error) {
	var product *model.Product

	if err := p.DB.Preload("Category").Preload("Image").Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p products) CreateAProduct(input *model.Product) error {

	return p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(input).Error; err != nil {
			return err
		}

		for _, value := range input.Image {
			value.ProductId = int(input.ID)
		}

		if err := tx.CreateInBatches(input.Image, 100).Error; err != nil {
			return err
		}

		for _, categoryValue := range input.Category {
			categoryValue.ProductId = int(input.ID)
		}

		if err := tx.CreateInBatches(input.Category, 100).Error; err != nil {
			return err
		}

		return nil
	})
}
