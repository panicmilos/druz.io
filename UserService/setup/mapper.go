package setup

import (
	"UserService/models"

	"github.com/devfeel/mapper"
)

func SetupMapper() {
	mapper.Register(&models.Account{})
	mapper.Register(&models.Profile{})
}
