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

func (receiver *AMQPReceiver) Initialize(settings *settings.AMQPSettings, exchangeName string) {
	conn, err := amqp.Dial(settings.ToConnectionString())
	errors.FailOnError(err, "Failed to connect to RabbitMQ")
	receiver.connection = conn

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

	receiver.channel = ch

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	errors.FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	errors.FailOnError(err, "Failed to bind a queue")

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
