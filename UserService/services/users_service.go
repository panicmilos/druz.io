package services

import (
	"UserService/errors"
	"UserService/models"
	"UserService/repository"
)

type UserService struct {
	repository *repository.Repository
}

// func (userService *UserService) ReadUsers() *[]models.Profile {
// 	return userService.repository.Users.ReadUsers()
// }

func (userService *UserService) Create(user *models.Account) (*models.Profile, error) {
	if userService.repository.Users.ReadByEmail(user.Email) != nil {
		return nil, errors.NewErrBadRequest("Email is already in use.")
	}

	return userService.repository.Users.Create(user), nil
}
