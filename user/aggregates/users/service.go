package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserService interface {
		AddNewUser(ctx context.Context, firstName string, lastName string, userName string, password string) (*User, error)
		AddNewAdminUser(ctx context.Context, firstName string, lastName string, userName string, password string) (*User, error)
		AddNewGuestUser(ctx context.Context) (*User, error)
		GetUserById(ctx context.Context, id string) (*User, error)
		AuthUser(ctx context.Context, username, password string) (bool, error)
	}
	userService struct {
		Repository IUserRepository
	}
)

func NewUserService(repository IUserRepository) UserService {
	return &userService{Repository: repository}
}

func (service userService) GetUserById(ctx context.Context, id string) (*User, error) {

	var (
		user     *User
		objectId primitive.ObjectID
		err      error
	)

	if objectId, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	if user, err = service.Repository.FindOneById(ctx, objectId); err != nil {
		return nil, err
	}

	return user, nil
}

func (service userService) AddNewUser(ctx context.Context,
	firstName string,
	lastName string,
	userName string,
	password string) (*User, error) {

	user := NewUser(firstName, lastName, userName, password)

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (service userService) AddNewAdminUser(ctx context.Context,
	firstName string,
	lastName string,
	userName string,
	password string) (*User, error) {

	user := NewAdminUser(firstName, lastName, userName, password)

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (service userService) AddNewGuestUser(ctx context.Context) (*User, error) {
	user := NewGuestUser()

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (service userService) AuthUser(ctx context.Context,
	username string,
	password string) (bool, error) {

	var (
		user *User
		err  error
	)

	if user, err = service.Repository.FindOneByUsername(ctx, username); err != nil {
		return false, err
	}

	return user.EncryptedPassword.VerifyPassword(password), nil
}
