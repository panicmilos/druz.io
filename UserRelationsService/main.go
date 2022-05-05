package main

import (
	"github.com/panicmilos/druz.io/UserRelationsService/server"
	"github.com/panicmilos/druz.io/UserRelationsService/services"
	"github.com/panicmilos/druz.io/UserRelationsService/setup"
)

func main() {

	setup.SetupEnviroment()
	setup.SetupDatabase()
	setup.SetupReplicators()

	server.New().Start()

	defer services.Provider.Delete()
}
