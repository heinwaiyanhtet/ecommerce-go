package services

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderPublisher struct {
	channel *amqp.Channel
}

func NewOrderPublisher(ch *amqp.Channel) (*OrderPublisher, error) {
	err := ch.ExchangeDeclare(
		"order_exchange",
		"fanout",
		true, false, false, false, nil,
	)
	if err != nil {
		return nil, err
	}

	return &OrderPublisher{channel: ch}, nil
}

func (p *OrderPublisher) PublishOrderCreated(orderID string) error {
	err := p.channel.Publish(
		"order_exchange",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(orderID),
		},
	)
	if err != nil {
		log.Printf("Failed to publish order created: %v", err)
		return err
	}
	log.Printf("Published order created: %s", orderID)
	return nil
}
