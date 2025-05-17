package repositories

import (
	"context"
	"database/sql"
	"github.com/heinwaiyanhtet/ecommerce-go/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, order *models.Order) error {
	query := `INSERT INTO orders (id, user_id, amount, status) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, order.ID, order.UserID, order.Amount, order.Status)
	return err
}
