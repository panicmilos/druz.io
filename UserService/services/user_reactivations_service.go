package services

import (
	"time"

	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/models"
	"github.com/panicmilos/druz.io/UserService/repository"

	"github.com/google/uuid"
)

type UserReactivationsService struct {
	repository *repository.Repository

	emailDispatcher *EmailService
	userReplicator  *UserReplicator
}

func (userReactivationsService *UserReactivationsService) Request(email string) (*models.UserReactivation, error) {
	profile := userReactivationsService.repository.Users.ReadDeactivatedByEmail(email)
	if profile == nil {
		return nil, errors.NewErrNotFound("Profile is given email is not deactivated.")
	}

	userReactivationsService.repository.UserReactivationsCollection.DeleteByProfileId(profile.ID)

	userReactivation := &models.UserReactivation{
		ProfileId: profile.ID,
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	userReactivationsService.emailDispatcher.Send(&dto.Email{
		Subject: "User Reactivation",
		From:    "panic.milos99@gmail.com",
		To:      email,
		Message: dto.EmailMessage{
			Template: "user_reactivation",
			Params: map[string]interface{}{
				"name":  profile.FirstName + " " + profile.LastName,
				"id":    profile.ID,
				"token": userReactivation.Token,
			},
		},
	})

	return userReactivationsService.repository.UserReactivationsCollection.Create(userReactivation), nil
}

func (userReactivationsService *UserReactivationsService) Reactivate(id uint, token string) (*models.Profile, error) {
	userReactivation := userReactivationsService.repository.UserReactivationsCollection.ReadByProfileId(id)
	if userReactivation == nil || userReactivation.Token != token {
		return nil, errors.NewErrBadRequest("Token is not valid.")
	}

	if time.Now().After(userReactivation.ExpiresAt) {
		return nil, errors.NewErrBadRequest("Token has expired.")
	}

	profile := userReactivationsService.repository.Users.ReadDeactivatedById(id)
	if profile == nil {
		return nil, errors.NewErrNotFound("Profile does not exist.")
	}

	userReactivationsService.repository.UserReactivationsCollection.DeleteByProfileId(id)

	profile.Disabled = false

	reactivatedUser := userReactivationsService.repository.Users.UpdateProfile(profile)

	userReactivationsService.userReplicator.Replicate(&dto.UserReplication{
		ReplicationType: "Reactivated",
		User:            reactivatedUser,
	})

	return reactivatedUser, nil
}
