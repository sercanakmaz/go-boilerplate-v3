package users

import (
	"errors"
	"fmt"
	"github.com/sercanakmaz/go-boilerplate-v3/contexts/user/core-api/aggregates"
	userEvents "github.com/sercanakmaz/go-boilerplate-v3/events/user/users"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/string-helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrAlreadyExistRole = errors.New("ErrAlreadyExistRole")

type User struct {
	Id                primitive.ObjectID            `json:"id" bson:"_id"`
	FirstName         string                        `json:"first_name"`
	LastName          string                        `json:"last_name"`
	UserName          string                        `json:"user_name"`
	EncryptedPassword *aggregates.EncryptedPassword `json:"encrypted_password"`
	Roles             []*UserRole                   `json:"roles"`
	domainEvents      []ddd.IBaseEvent              `json:"domain_events"`
}

func (u *User) ClearDomainEvents() {
	u.domainEvents = nil
}

func (u *User) GetDomainEvents() []ddd.IBaseEvent {
	return u.domainEvents
}

func (u *User) RaiseEvent(event ddd.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

func NewUser(firstName, lastName, username, password string) *User {

	var user *User

	if string_helper.IsNullOrEmpty(username) {
		panic(fmt.Errorf("%v %w", "username", string_helper.ErrIsNullOrEmpty))
	}

	user = &User{
		Id:                primitive.NewObjectID(),
		FirstName:         firstName,
		LastName:          lastName,
		UserName:          username,
		EncryptedPassword: aggregates.NewEncryptedPassword(password),
	}

	user.RaiseEvent(&userEvents.UserCreated{
		Id:        user.Id,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  username,
	})

	return user
}

func NewGuestUser() *User {

	user := NewUser("", "", "Guest", "12345")
	user.AddUserRole(UserRole_Guest)

	return user
}

func NewAdminUser(firstName, lastName, username, password string) *User {

	user := NewUser(firstName, lastName, username, password)
	user.AddUserRole(UserRole_Admin)

	return user
}

func (u *User) AddUserRole(role *UserRole) {

	if role == nil {
		panic(fmt.Errorf("%v %w", "role", string_helper.ErrIsNullOrEmpty))
	}

	for _, roleItem := range u.Roles {
		if roleItem.Name == role.Name {
			panic(fmt.Errorf("%v %w", "role", ErrAlreadyExistRole))
		}
	}

	u.Roles = append(u.Roles, role)
}

func (u *User) ChangePassword(oldPassword, newPassword string) {

	if !u.EncryptedPassword.VerifyPassword(oldPassword) {
		panic("")
	}

	u.EncryptedPassword = aggregates.NewEncryptedPassword(newPassword)
}
