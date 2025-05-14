package services

import (
	"github.com/ecommerce-go/internal/repository"
	"github.com/ecommerce-go/internal/model"
)

type UserService interface {
	GetUser() (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser() (*models.User, error) {
	return s.repo.FetchUser()
}
