package mapper

import (
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

// Users map for query get users
// TODO: need to add return address
func Users(users []*model.User) []*generated.User {
	var out []*generated.User

	for _, u := range users {
		out = append(out, &generated.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       int(u.Age),
			ID:        int(u.ID),
		})
	}

	return out
}

func mapGeneratedAddressesToUser(addressInput []*model.Address) []*generated.Address {
	var addresses []*generated.Address

	for _, addressInput := range addressInput {
		addresses = append(addresses, &generated.Address{
			PostalCode:  addressInput.PostalCode,
			Country:     addressInput.Country,
			City:        addressInput.City,
			AddressLine: addressInput.AddressLine,
		})
	}

	return addresses
}

func GeneratedUser(user *model.User) *generated.User {
	return &generated.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       int(user.Age),
		Phone:     user.Phone,
		Address:   mapGeneratedAddressesToUser(user.Address),
	}
}
