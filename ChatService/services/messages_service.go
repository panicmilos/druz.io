package services

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ambelovsky/gosf"
	"github.com/jellydator/ttlcache/v3"
	"github.com/panicmilos/druz.io/ChatService/dto"
	"github.com/panicmilos/druz.io/ChatService/errors"
	"github.com/panicmilos/druz.io/ChatService/helpers"
	"github.com/panicmilos/druz.io/ChatService/models"
	"github.com/panicmilos/druz.io/ChatService/repository"
)

type MessagesService struct {
	repository *repository.Repository

	UsersService   *UsersService
	Clients        *ttlcache.Cache[string, *gosf.Client]
	SessionStorage *helpers.SessionStorage
}

func (messagesService *MessagesService) ChatsWith(forId string) *[]dto.Chat {
	return messagesService.repository.Messages.ChatsWith(forId)
}

func (messagesService *MessagesService) ReadMessage(chat string, messageId string) (*models.Message, error) {
	message := messagesService.repository.Messages.ReadMessage(chat, messageId)
	if message == nil {
		return nil, errors.NewErrNotFound("Message does not exist.")
	}

	return message, nil
}

func (messagesService *MessagesService) ReadChat(chat string) ([]*models.Message, error) {
	messages := messagesService.repository.Messages.ReadChat(chat)
	if messages == nil {
		return nil, errors.NewErrNotFound("Chat does not exist.")
	}

	return messages, nil
}

func (messagesService *MessagesService) SearchChat(chat string, searchParams *dto.ChatSearchParams) []*models.Message {
	return messagesService.repository.Messages.SearchChat(chat, searchParams)
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

	message.CreatedAt = time.Now()
	message.DeletedBy1 = ""
	message.DeletedBy2 = ""

	createdMessage, err := messagesService.repository.Messages.Create(message), nil

	if err := messagesService.NotifyMessagedUser(message); err != nil {
		return nil, err
	}

	return createdMessage, nil
}

func (messagesService *MessagesService) NotifyMessagedUser(message *models.Message) error {
	if messagesService.Clients.Get(message.ToId) == nil {
		return nil
	}

	from, err := messagesService.UsersService.ReadById(message.FromId)
	if err != nil {
		return err
	}

	messageNotification := &dto.MessageNotification{
		Message: message,
		From:    from,
	}

	serializedNotification, _ := json.Marshal(messageNotification)
	print(string(serializedNotification))
	gosf.Broadcast(message.ToId, "messages_chat", gosf.NewSuccessMessage(string(serializedNotification)))
	gosf.Broadcast(message.ToId, "messages_sidebar", gosf.NewSuccessMessage(string(serializedNotification)))
	gosf.Broadcast(message.ToId, "messages_global", gosf.NewSuccessMessage(string(serializedNotification)))

	return nil
}

func (messagesService *MessagesService) DeleteMessage(chat string, messageId string, mode string) (*models.Message, error) {
	message, err := messagesService.ReadMessage(chat, messageId)
	if err != nil {
		return nil, err
	}

	deleteFor := []string{}
	if mode == "for_me" {
		deleteFor = append(deleteFor, messagesService.SessionStorage.AuthenticatedUserId)
	} else {
		deleteFor = append(deleteFor, message.FromId, message.ToId)

		if err := messagesService.NotifyAboutDeletedMesage(message); err != nil {
			return nil, err
		}
	}

	return messagesService.repository.Messages.DeleteMessage(chat, messageId, deleteFor), nil
}

func (messagesService *MessagesService) NotifyAboutDeletedMesage(message *models.Message) error {

	from, err := messagesService.UsersService.ReadById(message.FromId)
	if err != nil {
		return err
	}

	messageNotification := &dto.MessageNotification{
		Message: message,
		From:    from,
	}

	serializedNotification, _ := json.Marshal(messageNotification)
	if messagesService.SessionStorage.AuthenticatedUserId != message.ToId {
		if messagesService.Clients.Get(message.ToId) != nil {
			gosf.Broadcast(message.ToId, "messages_delete", gosf.NewSuccessMessage(string(serializedNotification)))
		}
	} else {
		if messagesService.Clients.Get(message.FromId) != nil {
			gosf.Broadcast(message.FromId, "messages_delete", gosf.NewSuccessMessage(string(serializedNotification)))
		}
	}

	return nil
}

func (messagesService *MessagesService) DeleteChat(chat string, mode string) ([]*models.Message, error) {
	_, err := messagesService.ReadChat(chat)
	if err != nil {
		return nil, err
	}

	deleteFor := []string{}
	if mode == "for_me" {
		deleteFor = append(deleteFor, messagesService.SessionStorage.AuthenticatedUserId)
	} else {
		deleteFor = append(deleteFor, strings.Split(chat, "-")...)

		if err := messagesService.NotifyAboutDeletedChat(chat, strings.Split(chat, "-")...); err != nil {
			return nil, err
		}
	}

	return messagesService.repository.Messages.DeleteChat(chat, deleteFor), nil
}

func (messagesService *MessagesService) NotifyAboutDeletedChat(chat string, users ...string) error {

	deletedChat := &dto.DeletedChatNotification{
		ChatId: chat,
	}

	serializedNotification, _ := json.Marshal(deletedChat)

	if messagesService.SessionStorage.AuthenticatedUserId != users[0] {
		if messagesService.Clients.Get(users[0]) != nil {
			gosf.Broadcast(users[0], "chat_delete", gosf.NewSuccessMessage(string(serializedNotification)))
		}
	} else {
		if messagesService.Clients.Get(users[1]) != nil {
			gosf.Broadcast(users[1], "chat_delete", gosf.NewSuccessMessage(string(serializedNotification)))
		}
	}

	return nil
}
