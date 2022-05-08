module github.com/panicmilos/druz.io/ChatService

go 1.18

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/ravendb/ravendb-go-client v0.0.0-20220329095225-8c32f0ab1fe3 // indirect
	github.com/sarulabs/di v2.0.0+incompatible // indirect
	github.com/streadway/amqp v1.0.0 // indirect
)

require github.com/panicmilos/druz.io/AMQPGO v0.0.0

replace github.com/panicmilos/druz.io/AMQPGO v0.0.0 => ../AMQPGO
