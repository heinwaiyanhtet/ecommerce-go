package services

import (
	"log"

	repository "github.com/heinwaiyanhtet/ecommerce-go/internal/repository"
	"github.com/heinwaiyanhtet/ecommerce-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderConsumer struct {
	channel   *amqp.Channel
	orderRepo repository.OrderRepository
	queueName string
}

func NewOrderConsumer(rabbitURL string, orderRepo repository.OrderRepository) (*OrderConsumer, error) {
	_, ch, err := rabbitmq.Connect(rabbitURL)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"order_exchange",
		"fanout",
		true, false, false, false, nil,
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		"",
		"order_exchange",
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &OrderConsumer{
		channel:   ch,
		orderRepo: orderRepo,
		queueName: q.Name,
	}, nil
}

func (c *OrderConsumer) StartConsuming() error {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			
			orderID := string(d.Body)
			log.Printf("Received order ID: %s", orderID)

			// Process order (example)
			// err := c.orderRepo.MarkOrderProcessed(orderID)
			// if err != nil {
			// 	log.Printf("Failed to process order %s: %v", orderID, err)
			// } else {
			// 	log.Printf("Order %s marked as processed", orderID)
			// }
		}
	}()

	log.Println("Order consumer started")
	return nil
}
