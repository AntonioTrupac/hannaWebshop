package types

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

func mapImagesToUser(imageInput []*generated.ImageInput) []*model.Image {
	var images []*model.Image

	for _, input := range imageInput {
		images = append(images, &model.Image{
			Url: input.URL,
		})
	}

	return images
}

// ModelProducts => mapping gql input to gorm model fields
func ModelProducts(ctx context.Context, in generated.ProductInput) *model.Product {
	return &model.Product{
		Price:       in.Price,
		Name:        in.Name,
		Description: in.Description,
		Rating:      in.Rating,
		Stock:       in.Stock,
		Image:       mapImagesToUser(in.Images),
	}
}

func mapGeneratedImagesToProduct(imageInput []*model.Image) []*generated.Image {
	var images []*generated.Image

	for _, input := range imageInput {
		images = append(images, &generated.Image{
			URL: input.Url,
		})
	}

	return images
}

// GeneratedProduct => mapping model fields to gql schema
func GeneratedProduct(product *model.Product) *generated.Product {
	return &generated.Product{
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Description,
		Rating:      product.Rating,
		Stock:       product.Stock,
		Image:       mapGeneratedImagesToProduct(product.Image),
	}
}
