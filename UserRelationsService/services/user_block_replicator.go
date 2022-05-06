package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/UserRelationsService/dto"
)

type UserBlockReplicator struct {
	sender *clients.AMQPSender
}

func (userBlockReplicator *UserBlockReplicator) Initialize() {
	userBlockReplicator.sender = &clients.AMQPSender{}
	settings := settings.GetDefaultAMQPSettings()
	userBlockReplicator.sender.Initialize(&settings, "user_block_replications")
}

func (userBlockReplicator *UserBlockReplicator) Replicate(replication *dto.UserBlockReplication) {
	body, _ := json.Marshal(replication)

	userBlockReplicator.sender.Send(body)
}

func (userBlockReplicator *UserBlockReplicator) Deinitialize() {
	userBlockReplicator.sender.Deinitialize()
}
