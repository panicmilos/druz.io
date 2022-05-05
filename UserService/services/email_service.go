package services

import (
	"encoding/json"

	"github.com/panicmilos/druz.io/UserService/dto"

	"github.com/panicmilos/druz.io/AMQPGO/clients"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
)

type EmailService struct {
	sender *clients.AMQPSender
}

func (emailService *EmailService) Initialize() {
	emailService.sender = &clients.AMQPSender{}
	settings := settings.GetDefaultAMQPSettings()
	emailService.sender.Initialize(&settings, "emails")
}

func (emailService *EmailService) Send(email dto.Email) {
	body, _ := json.Marshal(email)

	emailService.sender.Send(body)
}

func (emailService *EmailService) Deinitialize() {
	emailService.sender.Deinitialize()
}
