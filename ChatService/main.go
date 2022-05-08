package main

import (
	"github.com/panicmilos/druz.io/ChatService/setup"
)

func main() {

	setup.SetupEnviroment()
	setup.SetupReplicators()

	forever := make(chan bool)
	<-forever
}
