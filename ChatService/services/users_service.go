package services

import (
	"github.com/panicmilos/druz.io/ChatService/errors"
	"github.com/panicmilos/druz.io/ChatService/models"
	"github.com/panicmilos/druz.io/ChatService/repository"
)

type UsersService struct {
	repository *repository.Repository
}

func (userService *UsersService) ReadById(id string) (*models.User, error) {
	user := userService.repository.Users.ReadById(id)
	if user == nil {
		return nil, errors.NewErrNotFound("User is not found")
	}

	return user, nil
}
