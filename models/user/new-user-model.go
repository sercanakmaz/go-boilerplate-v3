package user

import (
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/user/core-api/aggregates/users"
)

type NewUserModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
}

func LoadFromUser(user *users.User) *NewUserModel {
	return &NewUserModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}
}
