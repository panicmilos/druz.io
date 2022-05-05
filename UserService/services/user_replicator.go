package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/UserService/dto"
)

type UserReplicator struct {
	sender *clients.AMQPSender
}

func (userReplicator *UserReplicator) Initialize() {
	userReplicator.sender = &clients.AMQPSender{}
	settings := settings.GetDefaultAMQPSettings()
	userReplicator.sender.Initialize(&settings, "user_replications")
}

func (userReplicator *UserReplicator) Replicate(replication *dto.UserReplication) {
	body, _ := json.Marshal(replication)

	userReplicator.sender.Send(body)
}

func (userReplicator *UserReplicator) Deinitialize() {
	userReplicator.sender.Deinitialize()
}
