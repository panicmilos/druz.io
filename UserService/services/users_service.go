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

func (userService *UserService) ReadById(id uint) (*models.Profile, error) {
	profile := userService.repository.Users.ReadById(id)
	if profile == nil {
		return nil, errors.NewErrNotFound("Profile is not found")
	}

	return profile, nil
}

func (userService *UserService) Create(user *models.Account) (*models.Profile, error) {
	if userService.repository.Users.ReadAccountByEmail(user.Email) != nil {
		return nil, errors.NewErrBadRequest("Email is already in use.")
	}

	user.Salt = helpers.GetRandomToken(16)
	user.Password = helpers.GetSaltedAndHashedPassword(user.Password, user.Salt)

	return userService.repository.Users.Create(user), nil
}

func (userService *UserService) Update(profile *models.Profile) (*models.Profile, error) {
	existingProfile, err := userService.ReadById(profile.ID)
	if err != nil {
		return nil, err
	}

	userService.repository.LivePlaces.DeleteByProfileId(profile.ID)
	userService.repository.WorkPlaces.DeleteByProfileId(profile.ID)
	userService.repository.Educations.DeleteByProfileId(profile.ID)
	userService.repository.Intereses.DeleteByProfileId(profile.ID)

	existingProfile.FirstName = profile.FirstName
	existingProfile.LastName = profile.LastName
	existingProfile.Birthday = profile.Birthday
	existingProfile.Gender = profile.Gender
	existingProfile.About = profile.About
	existingProfile.PhoneNumber = profile.PhoneNumber
	existingProfile.LivePlaces = profile.LivePlaces
	existingProfile.WorkPlaces = profile.WorkPlaces
	existingProfile.Educations = profile.Educations
	existingProfile.Intereses = profile.Intereses

	return userService.repository.Users.UpdateProfile(existingProfile), nil
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

func (userService *UserService) Delete(id uint) (*models.Profile, error) {
	existingProfile, err := userService.ReadById(id)
	if err != nil {
		return nil, err
	}

	return userService.repository.Users.Delete(existingProfile.ID), nil
}
