package setup

import (
	"github.com/panicmilos/druz.io/UserRelationsService/services"
)

func SetupReplicators() {
	userReplicator := services.Provider.Get(services.UsersReplicator).(*services.UserReplicator)

	userReplicator.StartReplicating()
}
