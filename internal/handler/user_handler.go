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

// internal/handlers/user_handler.go
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

    users, err := h.service.GetAllUsers()
    if err != nil {
        http.Error(w, "failed to get users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
	
}

