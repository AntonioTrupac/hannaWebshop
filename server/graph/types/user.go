package types

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

func ModelUser(ctx context.Context, in generated.UserInput) *model.User {
	return &model.User{
		Email:     in.Email,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Age:       uint(in.Age),
	}
}
