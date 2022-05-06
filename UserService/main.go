package main

import (
	"github.com/panicmilos/druz.io/UserService/server"
	"github.com/panicmilos/druz.io/UserService/services"
	"github.com/panicmilos/druz.io/UserService/setup"
)

func main() {
	setup.SetupEnviroment()
	setup.SetupDatabase()
	setup.SetupCronTasks()
	setup.SetupReplicators()

	server.New().Start()

	defer services.Provider.Delete()
}
