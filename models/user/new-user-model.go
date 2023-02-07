package user

import "go-boilerplate-v3/user/aggregates/users"

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
