package setup

import "github.com/panicmilos/druz.io/UserService/services"

func SetupReplicators() {
	userBlockReplicator := services.Provider.Get(services.UserBlocksReplicator).(*services.UserBlockReplicator)

	userBlockReplicator.StartReplicating()
}
