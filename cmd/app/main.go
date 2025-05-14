package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/ecommerce-go/internal/handler"
	"github.com/ecommerce-go/internal/repository"
	"github.com/ecommerce-go/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Print the environment variables (optional)
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))

	// Initialize repository and services
	repo, err := repositories.NewUserRepo()
	if err != nil {
		log.Fatalf("Error initializing user repo: %v", err)
	}
	userService := services.NewUserService(repo)
	userHandler := handlers.NewUserHandler(userService)

	// Set up HTTP routing
	http.HandleFunc("/user", userHandler.GetUser)
	port := ":8080"  // Example port
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
