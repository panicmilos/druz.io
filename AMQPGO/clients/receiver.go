package clients

import (
	"github.com/panicmilos/druz.io/AMQPGO/errors"
	"github.com/panicmilos/druz.io/AMQPGO/settings"
	"github.com/streadway/amqp"
)

type AMQPConsumerFunction func([]byte)

type AMQPReceiver struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func (receiver *AMQPReceiver) Initialize(settings *settings.AMQPSettings, queueName string) {
	conn, err := amqp.Dial(settings.ToConnectionString())
	errors.FailOnError(err, "Failed to connect to RabbitMQ")
	receiver.connection = conn

	ch, err := conn.Channel()
	errors.FailOnError(err, "Failed to open a channel")
	receiver.channel = ch

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	errors.FailOnError(err, "Failed to declare a queue")
	receiver.queue = &q
}

func (receiver *AMQPReceiver) Consume(consumerFunc AMQPConsumerFunction) {

	msgs, err := receiver.channel.Consume(
		receiver.queue.Name, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	errors.FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			consumerFunc(d.Body)
		}
	}()
}

func (receiver *AMQPReceiver) Deinitialize() {
	defer receiver.connection.Close()
	defer receiver.channel.Close()
}
