package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/ChatService/dto"
	"github.com/panicmilos/druz.io/ChatService/repository"
)

type UsersReplicator struct {
	receiver *clients.AMQPReceiver
	Users    *repository.UsersCollection
}

func (usersReplicator *UsersReplicator) Initialize() {
	usersReplicator.receiver = &clients.AMQPReceiver{}
	settings := settings.GetDefaultAMQPSettings()
	usersReplicator.receiver.Initialize(&settings, "user_replications")
}

func (usersReplicator *UsersReplicator) StartReplicating() {
	usersReplicator.receiver.Consume(func(b []byte) {
		userReplication := &dto.UserReplication{}
		json.Unmarshal(b, userReplication)

		user := userReplication.User.ToModel()

		switch userReplication.ReplicationType {
		case "Create":
			usersReplicator.Users.Create(user)
		case "Update":
			usersReplicator.Users.Update(user)
		case "Delete":
			usersReplicator.Users.Delete(user.ID)
		case "Disable":
			usersReplicator.Users.Disable(user.ID)
		case "Reactivate":
			usersReplicator.Users.Reactivate(user.ID)
		}
	})
}

func (usersReplicator *UsersReplicator) Deinitialize() {
	usersReplicator.receiver.Deinitialize()
}
