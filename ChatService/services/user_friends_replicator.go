package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/ChatService/dto"
	"github.com/panicmilos/druz.io/ChatService/repository"
	ravendb "github.com/ravendb/ravendb-go-client"
)

type UserFriendsReplicator struct {
	receiver *clients.AMQPReceiver
	store    *ravendb.DocumentStore
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

		session, _ := userFriendsReplicator.store.OpenSession("")
		userFriendsCollection := &repository.UserFriendsCollection{
			Session: session,
		}

		switch userFriendReplication.ReplicationType {
		case "Add":
			userFriendsCollection.Create(userFriend)
		case "Remove":
			userFriendsCollection.Delete(userFriend.UserId, userFriend.FriendId)
		}

		userFriendsCollection.Session.Close()
	})
}

func (userFriendsReplicator *UserFriendsReplicator) Deinitialize() {
	userFriendsReplicator.receiver.Deinitialize()
}
