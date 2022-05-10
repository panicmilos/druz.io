package main

import (
	"github.com/panicmilos/druz.io/ChatService/server"
	"github.com/panicmilos/druz.io/ChatService/services"
	"github.com/panicmilos/druz.io/ChatService/setup"
)

func main() {

	setup.SetupEnviroment()
	setup.SetupReplicators()

	server.New().Start()

	defer services.Provider.Delete()
}
