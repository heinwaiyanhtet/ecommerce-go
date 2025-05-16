package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/ecommerce-go/internal/service"
)

type AuthHandler struct{
	auth services.AuthService
}

func NewAuthHandler(a services.AuthService) *AuthHandler {
    return &AuthHandler{auth: a}
}



func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request)  {
	
	 var req struct{
			name string `json:"name"`
			Password string `json:"password"`
	 }

	 if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	 }

	 user, err := h.auth.Register(req.name, req.Password)

	 if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)

}


func (h *AuthHandler) Login (w http.ResponseWriter , r *http.Request){

	var req struct {
        name string `name:"username"`
        Password string `json:"password"`
    }

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	token, err := h.auth.Login(req.name, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
        "exp":   time.Now().Add(time.Hour * 24).Format(time.RFC3339),
    })
}