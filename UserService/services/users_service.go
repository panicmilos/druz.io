package services

import (
	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/helpers"
	"github.com/panicmilos/druz.io/UserService/models"
	"github.com/panicmilos/druz.io/UserService/repository"
)

type UserService struct {
	repository *repository.Repository

	userReplicator *UserReplicator
}

func (userService *UserService) Search(params *dto.UsersSearchParams) *[]models.Profile {

	return userService.repository.Users.Search(params)
}

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

	createdUser := userService.repository.Users.Create(user)

	userService.userReplicator.Replicate(&dto.UserReplication{
		ReplicationType: "Create",
		User:            createdUser,
	})

	return createdUser, nil
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

	updatedUser := userService.repository.Users.UpdateProfile(existingProfile)

	userService.userReplicator.Replicate(&dto.UserReplication{
		ReplicationType: "Update",
		User:            updatedUser,
	})

	return updatedUser, nil
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

	deletedUser := userService.repository.Users.Delete(existingProfile.ID)

	userService.userReplicator.Replicate(&dto.UserReplication{
		ReplicationType: "Delete",
		User:            deletedUser,
	})

	return deletedUser, nil
}

func (userService *UserService) Disable(id uint) (*models.Profile, error) {
	existingProfile, err := userService.ReadById(id)
	if err != nil {
		return nil, err
	}

	existingProfile.Disabled = true

	disabledUser := userService.repository.Users.UpdateProfile(existingProfile)

	userService.userReplicator.Replicate(&dto.UserReplication{
		ReplicationType: "Disable",
		User:            disabledUser,
	})

	return disabledUser, nil
}
