package types

import (
	"context"
	"fmt"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

func mapImagesToProducts(images []*model.Image) []*generated.Image {
	var out []*generated.Image

	for _, i := range images {
		out = append(out, &generated.Image{
			URL:       i.Url,
			ID:        int(i.ID),
			ProductID: i.ProductId,
		})
	}

	return out
}

// Products map for query get products
func Products(products []*model.Product) []*generated.Product {
	var out []*generated.Product

	for _, p := range products {
		out = append(out, &generated.Product{
			ID:          int(p.ID),
			Price:       p.Price,
			Name:        p.Name,
			Description: p.Description,
			Rating:      p.Rating,
			Stock:       p.Stock,
			Image:       mapImagesToProducts(p.Image),
		})
	}

	fmt.Println("OUT OUT OUT", out)

	return out
}

// ModelProducts => mapping gql input to gorm model fields
func ModelProducts(ctx context.Context, in generated.ProductInput) *model.Product {
	return &model.Product{
		Price:       in.Price,
		Name:        in.Name,
		Description: in.Description,
		Rating:      in.Rating,
		Stock:       in.Stock,
		Image:       mapImagesToProduct(in.Images),
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

func mapImagesToProduct(imageInput []*generated.ImageInput) []*model.Image {
	var images []*model.Image

	for _, input := range imageInput {
		images = append(images, &model.Image{
			Url: input.URL,
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
