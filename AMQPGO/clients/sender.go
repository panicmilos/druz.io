package clients

import (
	"github.com/panicmilos/druz.io/AMQPGO/errors"
	"github.com/panicmilos/druz.io/AMQPGO/settings"

	"github.com/streadway/amqp"
)

type AMQPSender struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func (sender *AMQPSender) Initialize(settings settings.AMQPSettings, queueName string) {
	conn, err := amqp.Dial(settings.ToConnectionString())

	errors.FailOnError(err, "Failed to connect to RabbitMQ")
	sender.connection = conn

	ch, err := conn.Channel()
	errors.FailOnError(err, "Failed to open a channel")
	sender.channel = ch

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	errors.FailOnError(err, "Failed to declare a queue")
	sender.queue = &q
}

func (sender *AMQPSender) Send(body []byte) {

	sender.channel.Publish(
		"",                // exchange
		sender.queue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func (sender *AMQPSender) Deinitialize() {
	defer sender.connection.Close()
	defer sender.channel.Close()
}
