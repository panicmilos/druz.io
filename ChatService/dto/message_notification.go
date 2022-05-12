package dto

import "github.com/panicmilos/druz.io/ChatService/models"

type MessageNotification struct {
	Message *models.Message
	From    *models.User
}
