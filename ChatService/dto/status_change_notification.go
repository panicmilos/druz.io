package dto

import "github.com/panicmilos/druz.io/ChatService/models"

type StatusChangeNotification struct {
	Status string
	User   *models.User
}
