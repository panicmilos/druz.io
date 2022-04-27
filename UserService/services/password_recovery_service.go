package services

import (
	"UserService/dto"
	"UserService/errors"
	"UserService/helpers"
	"UserService/models"
	"UserService/repository"
	"time"

	"github.com/google/uuid"
)

type PasswordRecoveryService struct {
	repository *repository.Repository

	emailDispatcher *EmailService
}

func (passwordRecoveryService *PasswordRecoveryService) Request(email string) (*models.PasswordRecovery, error) {
	account := passwordRecoveryService.repository.Users.ReadAccountByEmail(email)
	if account == nil {
		return nil, errors.NewErrNotFound("Account is not found.")
	}

	passwordRecoveryService.repository.PasswordRecoveriesCollection.DeleteByProfileId(account.Profile.ID)

	passwordRecovery := &models.PasswordRecovery{
		ProfileId: account.Profile.ID,
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	passwordRecoveryService.emailDispatcher.Send(dto.Email{
		Subject: "Password recovery",
		From:    "panic.milos99@gmail.com",
		To:      account.Email,
		Message: dto.EmailMessage{
			Template: "password_recovery",
			Params: map[string]interface{}{
				"name":  account.Profile.FirstName + " " + account.Profile.LastName,
				"id":    account.Profile.ID,
				"token": passwordRecovery.Token,
			},
		},
	})

	return passwordRecoveryService.repository.PasswordRecoveriesCollection.Create(passwordRecovery), nil
}

func (passwordRecoveryService *PasswordRecoveryService) Recover(id uint, token string, newPassword string) (*models.Profile, error) {
	passwordRecovery := passwordRecoveryService.repository.PasswordRecoveriesCollection.ReadByProfileId(id)
	if passwordRecovery == nil || passwordRecovery.Token != token {
		return nil, errors.NewErrBadRequest("Token is not valid.")
	}

	if time.Now().After(passwordRecovery.ExpiresAt) {
		return nil, errors.NewErrBadRequest("Token has expired.")
	}

	passwordRecoveryService.repository.PasswordRecoveriesCollection.DeleteByProfileId(id)

	account := passwordRecoveryService.repository.Users.ReadAccountByProfileId(id)
	account.Salt = helpers.GetRandomToken(16)
	account.Password = helpers.GetSaltedAndHashedPassword(newPassword, account.Salt)

	return &passwordRecoveryService.repository.Users.UpdateAccount(account).Profile, nil
}
