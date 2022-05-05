package dto

import "github.com/panicmilos/druz.io/UserService/models"

type UserReplication struct {
	ReplicationType string
	User            *models.Profile
}
