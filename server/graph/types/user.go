package types

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

func mapAddressesToUser(addressInput []*generated.AddressInput) []model.Address {
	var addresses []model.Address

	for _, addressInput := range addressInput {
		addresses = append(addresses, model.Address{
			AddressLine: addressInput.AddressLine,
			City:        addressInput.City,
			Country:     addressInput.Country,
			PostalCode:  addressInput.PostalCode,
			UserId:      addressInput.UserID,
		})
	}

	return addresses
}

func ModelUser(ctx context.Context, in generated.UserInput) *model.User {
	return &model.User{
		Email:     in.Email,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Age:       uint(in.Age),
		Phone:     in.Phone,
		Address:   mapAddressesToUser(in.Address),
	}
}
