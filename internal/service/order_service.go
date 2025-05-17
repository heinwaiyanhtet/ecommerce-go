package services

import (
	"context"
	"log"
	"github.com/google/uuid"
	"github.com/heinwaiyanhtet/ecommerce-go/internal/model"
	"github.com/heinwaiyanhtet/ecommerce-go/internal/repository"
)

type OrderService struct {
	orderRepo      *repositories.OrderRepository
	orderPublisher *OrderPublisher
}

func NewOrderService(repo *repositories.OrderRepository, publisher *OrderPublisher) *OrderService {
	return &OrderService{
		orderRepo:      repo,
		orderPublisher: publisher,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) error {

	order.ID = uuid.New().String()
	order.Status = "created"

	err := s.orderRepo.Create(ctx, order)
	if err != nil {
		return err
	}

	go func(orderID string) {
		err := s.orderPublisher.PublishOrderCreated(orderID)
		if err != nil {
			log.Printf("Failed to publish order created event: %v", err)
		}
	}(order.ID)

	return nil
}
