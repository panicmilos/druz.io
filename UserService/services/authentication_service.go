package services

import (
	"strconv"

	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/helpers"
	"github.com/panicmilos/druz.io/UserService/repository"
)

type AuthenticationService struct {
	repository *repository.Repository
}

func (authenticationService *AuthenticationService) Auth(email string, password string) (*dto.AuthenticatedUser, error) {
	account := authenticationService.repository.Users.ReadAccountByEmail(email)
	if account == nil {
		return nil, errors.NewErrUnauthorized("The Combination of username and password doesn't match any account.")
	}

	if helpers.GetSaltedAndHashedPassword(password, account.Salt) != account.Password {
		return nil, errors.NewErrUnauthorized("The Combination of username and password doesn't match any account.")
	}

	claims := helpers.Claims{
		Name: strconv.FormatUint(uint64(account.Profile.ID), 10),
		Role: strconv.Itoa((int(account.Role))),
	}
	jwtToken, _ := helpers.GetJwtToken(claims)

	return &dto.AuthenticatedUser{
		Jwt:     jwtToken,
		Profile: &account.Profile,
	}, nil
}
