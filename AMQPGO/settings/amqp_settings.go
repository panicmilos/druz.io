package settings

import (
	"fmt"
	"os"
)

type AMQPSettings struct {
	Username string
	Password string
	Host     string
	Port     string
}

func (aqmpSettings *AMQPSettings) ToConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", aqmpSettings.Username, aqmpSettings.Password, aqmpSettings.Host, aqmpSettings.Port)
}

func GetDefaultAMQPSettings() AMQPSettings {
	return AMQPSettings{
		Username: os.Getenv("AMQP_USERNAME"),
		Password: os.Getenv("AMQP_PASSWORD"),
		Host:     os.Getenv("AMQP_HOST"),
		Port:     os.Getenv("AMQP_PORT"),
	}
}
