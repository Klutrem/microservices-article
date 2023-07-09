package lib

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitbroker struct {
	Connection amqp.Connection
}

func NewRabbitbroker(env Env) Rabbitbroker {
	rabbitHost := env.BrokerHost
	rabbitPort := env.BrokerPort
	rabbitUser := env.BrokerUser
	rabbitPass := env.BrokerPass

	connectionString := fmt.Sprintf("amqp://",
		rabbitUser, ":", rabbitPass,
		"@", rabbitHost, ":", rabbitPort)

	conn, err := amqp.DialConfig(connectionString, amqp.Config{})
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	defer func() {
		_ = conn.Close()
	}()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close()
	}()
}
