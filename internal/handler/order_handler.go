package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/heinwaiyanhtet/ecommerce-go/internal/model"
	"github.com/heinwaiyanhtet/ecommerce-go/internal/service"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(s *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: s,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.orderService.CreateOrder(r.Context(), &order)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(order)
}
