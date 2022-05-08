package dto

import "github.com/panicmilos/druz.io/UserRelationsService/models"

type UserFriendReplication struct {
	ReplicationType string
	UserFriend      *models.UserFriend
}
