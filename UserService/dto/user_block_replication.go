package dto

import "github.com/panicmilos/druz.io/UserService/models"

type UserBlockReplication struct {
	ReplicationType string
	UserBlock       *models.UserBlock
}
