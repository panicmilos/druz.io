package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/ChatService/dto"
	"github.com/panicmilos/druz.io/ChatService/repository"
)

type UserFriendsReplicator struct {
	receiver    *clients.AMQPReceiver
	UserFriends *repository.UserFriendsCollection
}

func (userFriendsReplicator *UserFriendsReplicator) Initialize() {
	userFriendsReplicator.receiver = &clients.AMQPReceiver{}
	settings := settings.GetDefaultAMQPSettings()
	userFriendsReplicator.receiver.Initialize(&settings, "user_friend_replications")
}

func (userFriendsReplicator *UserFriendsReplicator) StartReplicating() {
	userFriendsReplicator.receiver.Consume(func(b []byte) {
		userFriendReplication := &dto.UserFriendReplication{}
		json.Unmarshal(b, userFriendReplication)

		userFriend := userFriendReplication.UserFriend.ToModel()

		switch userFriendReplication.ReplicationType {
		case "Add":
			userFriendsReplicator.UserFriends.Create(userFriend)
		case "Remove":
			userFriendsReplicator.UserFriends.Delete(userFriend.UserId, userFriend.FriendId)
		}
	})
}

func (userFriendsReplicator *UserFriendsReplicator) Deinitialize() {
	userFriendsReplicator.receiver.Deinitialize()
}
