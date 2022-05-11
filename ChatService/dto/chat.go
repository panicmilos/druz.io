package dto

import "github.com/panicmilos/druz.io/ChatService/models"

type Chat struct {
	Chat string
	User *models.User
}
