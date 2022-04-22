package services

import (
	"UserService/errors"
	"UserService/helpers"
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
	if userService.repository.Users.ReadAccountByEmail(user.Email) != nil {
		return nil, errors.NewErrBadRequest("Email is already in use.")
	}

	user.Salt = helpers.GetRandomToken(16)
	user.Password = helpers.GetSaltedAndHashedPassword(user.Password, user.Salt)

	return userService.repository.Users.Create(user), nil
}

func (userService *UserService) ChangePassword(id uint, currentPassword string, newPassword string) (*models.Profile, error) {
	account := userService.repository.Users.ReadAccountByProfileId(id)
	if account == nil {
		return nil, errors.NewErrNotFound("Account is not found.")
	}

	if account.Password != helpers.GetSaltedAndHashedPassword(currentPassword, account.Salt) {
		return nil, errors.NewErrBadRequest("Current password does not match.")
	}

	account.Salt = helpers.GetRandomToken(16)
	account.Password = helpers.GetSaltedAndHashedPassword(newPassword, account.Salt)

	return &userService.repository.Users.UpdateAccount(account).Profile, nil

}
