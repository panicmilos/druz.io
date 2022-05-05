package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/UserRelationsService/dto"
	"github.com/panicmilos/druz.io/UserRelationsService/repository"
)

type UserReplicator struct {
	receiver *clients.AMQPReceiver
	Users    *repository.UsersCollection
}

func (userReplicator *UserReplicator) Initialize() {
	userReplicator.receiver = &clients.AMQPReceiver{}
	settings := settings.GetDefaultAMQPSettings()
	userReplicator.receiver.Initialize(&settings, "user_replications")
}

func (userReplicator *UserReplicator) StartReplicating() {
	userReplicator.receiver.Consume(func(b []byte) {
		userReplication := &dto.UserReplication{}
		json.Unmarshal(b, userReplication)

		switch userReplication.ReplicationType {
		case "Create":
			userReplicator.Users.Create(userReplication.User)
		case "Update":
			userReplicator.Users.Update(userReplication.User)
		case "Delete":
			userReplicator.Users.Delete(userReplication.User.ID)
		case "Disable":
			userReplicator.Users.Disable(userReplication.User.ID)
		case "Reactivate":
			userReplicator.Users.Reactivate(userReplication.User.ID)
		}
	})
}

func (userReplicator *UserReplicator) Deinitialize() {
	userReplicator.receiver.Deinitialize()
}
