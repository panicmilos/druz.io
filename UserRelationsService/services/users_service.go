package services

import (
	"github.com/panicmilos/druz.io/UserRelationsService/errors"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
	"github.com/panicmilos/druz.io/UserRelationsService/repository"
)

type UsersService struct {
	repository *repository.Repository
}

func (userService *UsersService) ReadById(id uint) (*models.User, error) {
	user := userService.repository.Users.ReadById(id)
	if user == nil {
		return nil, errors.NewErrNotFound("User is not found")
	}

	return user, nil
}
