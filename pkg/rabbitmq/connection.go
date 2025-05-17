package rabbitmq

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
	once    sync.Once
)

func Connect(url string) (*amqp.Connection, *amqp.Channel, error) {
	var err error
	once.Do(func() {
		conn, err = amqp.Dial(url)
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %v", err)
			return
		}
		channel, err = conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open channel: %v", err)
			return
		}
	})
	return conn, channel, err
}

// Close closes channel and connection
func Close() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
