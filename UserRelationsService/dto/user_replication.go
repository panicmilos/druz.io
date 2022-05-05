package dto

import "github.com/panicmilos/druz.io/UserRelationsService/models"

type UserReplication struct {
	ReplicationType string
	User            *models.User
}
