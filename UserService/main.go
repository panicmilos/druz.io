package main

import (
	"UserService/server"
	"UserService/services"
	"UserService/setup"
)

func main() {
	setup.SetupEnviroment()
	setup.SetupDatabase()
	setup.SetupCronTasks()

	server.New().Start()

	defer services.Provider.Delete()
}
