package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/UserRelationsService/dto"
)

type UserFriendReplicator struct {
	sender *clients.AMQPSender
}

func (userFriendReplicator *UserFriendReplicator) Initialize() {
	userFriendReplicator.sender = &clients.AMQPSender{}
	settings := settings.GetDefaultAMQPSettings()
	userFriendReplicator.sender.Initialize(&settings, "user_friend_replications")
}

func (userFriendReplicator *UserFriendReplicator) Replicate(replication *dto.UserFriendReplication) {
	body, _ := json.Marshal(replication)

	userFriendReplicator.sender.Send(body)
}

func (userFriendReplicator *UserFriendReplicator) Deinitialize() {
	userFriendReplicator.sender.Deinitialize()
}
