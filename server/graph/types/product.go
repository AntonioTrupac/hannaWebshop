package types

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	helpers "github.com/AntonioTrupac/hannaWebshop/graph/types/helpers"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

// Products map for query get products
func GetProductsFromDb(products []*model.Product) []*generated.Product {
	var out []*generated.Product

	for _, p := range products {
		out = append(out, &generated.Product{
			ID:          int(p.ID),
			Price:       p.Price,
			Name:        p.Name,
			Description: p.Description,
			Rating:      p.Rating,
			Stock:       p.Stock,
			Image:       helpers.MapImagesToProducts(p.Image),
			Category:    mapGeneratedCategoriesToProduct(p.Category),
		})
	}

	return out
}

func GetProductFromDb(product *model.Product) *generated.Product {
	var out *generated.Product

	out = &generated.Product{
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Description,
		Rating:      product.Rating,
		Stock:       product.Stock,
		Image:       helpers.MapImagesToProducts(product.Image),
		Category:    mapGeneratedCategoriesToProduct(product.Category),
	}

	return out
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

func mapCategoriesToProduct(categoryInput []*generated.CategoryInput) []*model.Category {
	var categories []*model.Category

	for _, input := range categoryInput {
		categories = append(categories, &model.Category{
			Name: input.Name,
		})
	}

	return categories
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
		Category:    mapCategoriesToProduct(in.Categories),
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

func mapGeneratedCategoriesToProduct(categoryInput []*model.Category) []*generated.Category {
	var categories []*generated.Category

	for _, input := range categoryInput {
		categories = append(categories, &generated.Category{
			Name: input.Name,
		})
	}

	return categories
}

// GeneratedProduct => returns a product after it is created via mutation CreateProducts
func GeneratedProduct(product *model.Product) *generated.Product {
	return &generated.Product{
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Description,
		Rating:      product.Rating,
		Stock:       product.Stock,
		Image:       mapGeneratedImagesToProduct(product.Image),
		Category:    mapGeneratedCategoriesToProduct(product.Category),
	}
}
