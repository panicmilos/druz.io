package setup

import (
	"github.com/joho/godotenv"
)

func SetupEnviroment() {
	godotenv.Load(".env")
}
