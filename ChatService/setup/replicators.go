package setup

import "github.com/panicmilos/druz.io/ChatService/services"

func SetupReplicators() {
	usersReplicator := services.Provider.Get(services.UserReplicator).(*services.UsersReplicator)

	usersReplicator.StartReplicating()
}
