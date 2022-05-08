package dto

import (
	"strconv"

	"github.com/panicmilos/druz.io/ChatService/models"
)

type UserFriendReplication struct {
	ReplicationType string
	UserFriend      *UserFriendDTO
}

type UserFriendDTO struct {
	ID       uint
	UserId   uint
	FriendId uint
}

func (userFriendDTO *UserFriendDTO) ToModel() *models.UserFriend {
	return &models.UserFriend{
		ID:       strconv.Itoa(int(userFriendDTO.ID)),
		UserId:   strconv.Itoa(int(userFriendDTO.UserId)),
		FriendId: strconv.Itoa(int(userFriendDTO.FriendId)),
	}
}
