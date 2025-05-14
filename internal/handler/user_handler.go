package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/ecommerce-go/internal/service"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUser is a handler function that fetches a user
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.GetUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
