package services

import (
	"errors"
	"log"
	"time"

	models "github.com/heinwaiyanhtet/ecommerce-go/internal/model"
	repositories "github.com/heinwaiyanhtet/ecommerce-go/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, password string) (*models.User, error)
	Login(name, password string) (string, error)
}

type authService struct {
	repo      repositories.UserRepository
	jwtSecret []byte
	tokenTTL  time.Duration
}

func NewAuthService(repo repositories.UserRepository, jwtSecret string, ttl time.Duration) AuthService {
	return &authService{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
		tokenTTL:  ttl,
	}
}

func (s *authService) Login(name string, password string) (string, error) {

	user, err := s.repo.GetByUserName(name)

    if err != nil {
		log.Printf("Login failed - user lookup error for name '%s': %v", name, err)
		return "", err
	}

	log.Printf("Login attempt - user found: ID=%d, Name=%s", user.ID, user.Name)

	if err != nil {
		return "", err
	}

	u := user

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(s.tokenTTL).Unix(),
	})
	return token.SignedString(s.jwtSecret)

}

func (s *authService) Register(username string, password string) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &models.User{Name: username, PasswordHash: string(hash)}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	user.PasswordHash = "" // never return hash

	return user, nil
}
