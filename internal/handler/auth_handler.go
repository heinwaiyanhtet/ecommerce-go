package handlers

import "github.com/ecommerce-go/internal/service"

type AuthHandler struct{
	auth services.AuthService
}

func NewAuthHandler(a services.AuthService) *AuthHandler {
    return &AuthHandler{auth: a}
}
