package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/panicmilos/druz.io/UserService/dto"

	"github.com/streadway/amqp"
)

type EmailService struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (emailService *EmailService) Initialize() {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", os.Getenv("AMQP_USERNAME"), os.Getenv("AMQP_PASSWORD"), os.Getenv("AMQP_HOST"), os.Getenv("AMQP_PORT"))
	conn, err := amqp.Dial(connectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	emailService.connection = conn

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	emailService.channel = ch

	q, err := ch.QueueDeclare(
		"emails", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")
	emailService.queue = &q
}

func (emailService *EmailService) Send(email dto.Email) {
	body, _ := json.Marshal(email)

	emailService.channel.Publish(
		"",                      // exchange
		emailService.queue.Name, // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func (emailService *EmailService) Deinitialize() {
	defer emailService.connection.Close()
	defer emailService.channel.Close()
}
