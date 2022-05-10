package services

import (
	"github.com/panicmilos/druz.io/ChatService/errors"
	"github.com/panicmilos/druz.io/ChatService/models"
	"github.com/panicmilos/druz.io/ChatService/repository"
)

type MessagesService struct {
	repository *repository.Repository

	UsersService *UsersService
}

func (messagesService *MessagesService) Create(message *models.Message) (*models.Message, error) {
	userFriend := messagesService.repository.UserFriends.ReadByIds(message.FromId, message.ToId)
	if userFriend == nil {
		return nil, errors.NewErrBadRequest("You are not friend with given user.")
	}

	_, err := messagesService.UsersService.ReadById(message.ToId)
	if err != nil {
		return nil, err
	}

	message.DeletedFor = []string{}

	return messagesService.repository.Messages.Create(message), nil
}
