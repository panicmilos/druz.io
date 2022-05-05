package main

import (
	"fmt"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/panicmilos/druz.io/UserRelationsService/server"
	"github.com/panicmilos/druz.io/UserRelationsService/services"
	"github.com/panicmilos/druz.io/UserRelationsService/setup"
)

func main() {

	setup.SetupEnviroment()

	receiver := &clients.AMQPReceiver{}
	settings := settings.GetDefaultAMQPSettings()
	receiver.Initialize(&settings, "user_replication")

	receiver.Consume(func(b []byte) { fmt.Printf("%s", b) })

	server.New().Start()

	defer receiver.Deinitialize()
	defer services.Provider.Delete()
}
