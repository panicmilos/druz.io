package main

import (
	"UserService/server"
	"UserService/services"
	"UserService/setup"
)

func main() {
	setup.SetupEnviroment()
	// setup.SetupDatabase()

	server.New().Start()

	defer services.Provider.Delete()
}
