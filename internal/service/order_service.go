package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/heinwaiyanhtet/ecommerce-go/internal/model"
	"github.com/heinwaiyanhtet/ecommerce-go/internal/repository"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) error {
	// generate ID for new order
	order.ID = uuid.New().String()
	order.Status = "created"

	// 1. Save order to DB
	err := s.orderRepo.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
