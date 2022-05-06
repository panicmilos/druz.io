package dto

import "github.com/panicmilos/druz.io/UserRelationsService/models"

type UserBlockReplication struct {
	ReplicationType string
	UserBlock       *models.UserBlock
}
