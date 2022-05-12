package main

import (
	"github.com/panicmilos/druz.io/ChatService/http_server"
	"github.com/panicmilos/druz.io/ChatService/services"
	"github.com/panicmilos/druz.io/ChatService/setup"
	"github.com/panicmilos/druz.io/ChatService/sockets_server"
)

func main() {

	setup.SetupEnviroment()
	setup.SetupReplicators()

	sockets_server.New().Start()
	http_server.New().Start()

	defer services.Provider.Delete()
}
