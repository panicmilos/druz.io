package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/ChatService/dto"
	"github.com/panicmilos/druz.io/ChatService/repository"
	ravendb "github.com/ravendb/ravendb-go-client"
)

type UsersReplicator struct {
	receiver *clients.AMQPReceiver
	store    *ravendb.DocumentStore
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

		session, _ := usersReplicator.store.OpenSession("")
		usersCollection := &repository.UsersCollection{
			Session: session,
		}

		switch userReplication.ReplicationType {
		case "Create":
			usersCollection.Create(user)
		case "Update":
			usersCollection.Update(user)
		case "Delete":
			usersCollection.Delete(user.ID)
		case "Disable":
			usersCollection.Disable(user.ID)
		case "Reactivate":
			usersCollection.Reactivate(user.ID)
		}

		usersCollection.Session.Close()
	})
}

func (usersReplicator *UsersReplicator) Deinitialize() {
	usersReplicator.receiver.Deinitialize()
}
