package mapper

import (
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

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

func GeneratedUser(user *model.User) *generated.User {
	return &generated.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       int(user.Age),
	}
}
