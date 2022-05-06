package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/repository"
)

type UserBlockReplicator struct {
	receiver   *clients.AMQPReceiver
	UserBlocks *repository.UserBlocksCollection
}

func (userBlockReplicator *UserBlockReplicator) Initialize() {
	userBlockReplicator.receiver = &clients.AMQPReceiver{}
	settings := settings.GetDefaultAMQPSettings()
	userBlockReplicator.receiver.Initialize(&settings, "user_block_replications")
}

func (userBlockReplicator *UserBlockReplicator) StartReplicating() {
	userBlockReplicator.receiver.Consume(func(b []byte) {
		userBlockReplication := &dto.UserBlockReplication{}
		json.Unmarshal(b, userBlockReplication)

		switch userBlockReplication.ReplicationType {
		case "Block":
			userBlockReplicator.UserBlocks.Create(userBlockReplication.UserBlock)
		case "Unblock":
			userBlockReplicator.UserBlocks.Delete(userBlockReplication.UserBlock.ID)
		}
	})
}

func (userBlockReplicator *UserBlockReplicator) Deinitialize() {
	userBlockReplicator.receiver.Deinitialize()
}
