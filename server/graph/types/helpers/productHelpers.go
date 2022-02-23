package types

import (
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

// DB to GQL

func MapImagesToProducts(images []*model.Image) []*generated.Image {
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
