package clients

import (
	"github.com/panicmilos/druz.io/AMQPGO/errors"
	"github.com/panicmilos/druz.io/AMQPGO/settings"

	"github.com/streadway/amqp"
)

type AMQPSender struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	exchangeName string
}

func (sender *AMQPSender) Initialize(settings *settings.AMQPSettings, exchangeName string) {
	conn, err := amqp.Dial(settings.ToConnectionString())
	errors.FailOnError(err, "Failed to connect to RabbitMQ")
	sender.connection = conn

	ch, err := conn.Channel()
	errors.FailOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	errors.FailOnError(err, "Failed to declare an exchange")

	sender.channel = ch
	sender.exchangeName = exchangeName
}

func (sender *AMQPSender) Send(body []byte) {

	sender.channel.Publish(
		sender.exchangeName, // exchange
		"",                  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func (sender *AMQPSender) Deinitialize() {
	defer sender.connection.Close()
	defer sender.channel.Close()
}
