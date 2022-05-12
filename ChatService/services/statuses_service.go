package services

import (
	"encoding/json"
	"fmt"

	"github.com/ambelovsky/gosf"
	"github.com/jellydator/ttlcache/v3"
	"github.com/panicmilos/druz.io/ChatService/dto"
	"github.com/panicmilos/druz.io/ChatService/repository"
)

type StatusesService struct {
	repository *repository.Repository

	UsersService *UsersService
	Clients      *ttlcache.Cache[string, *gosf.Client]
}

func (statusesService *StatusesService) ReadStatuses(id string) *[]dto.Status {
	userFriends := statusesService.repository.UserFriends.ReadByUserId(id)

	statuses := []dto.Status{}

	for _, userFriend := range userFriends {
		var status string

		if statusesService.Clients.Get(userFriend.FriendId) != nil {
			status = "online"
		} else {
			status = "offline"
		}

		statuses = append(statuses, dto.Status{
			Status: status,
			UserId: userFriend.FriendId,
		})
	}

	return &statuses
}

func (statusesService *StatusesService) NotifyCameOnline(id string) {
	statusesService.notifyStatusChange(id, "online")
}

func (statusesService *StatusesService) NotifyWentOffline(id string) {
	statusesService.notifyStatusChange(id, "offline")
}

func (statusesService *StatusesService) notifyStatusChange(id string, status string) {
	user, _ := statusesService.UsersService.ReadById(id)
	userFriends := statusesService.repository.UserFriends.ReadByUserId(id)

	statusChangeNotification := &dto.StatusChangeNotification{
		Status: status,
		User:   user,
	}
	serializedNotification, _ := json.Marshal(statusChangeNotification)

	for _, userFriend := range userFriends {
		if statusesService.Clients.Get(userFriend.FriendId) != nil {
			fmt.Println(userFriend.FriendId)
			gosf.Broadcast(userFriend.FriendId, "statuses", gosf.NewSuccessMessage(string(serializedNotification)))
		}
	}
}
