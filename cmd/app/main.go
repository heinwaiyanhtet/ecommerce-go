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

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))

	repo, err := repositories.NewUserRepo()

	if err != nil {
		log.Fatalf("Error initializing user repo: %v", err)
	}
	
	userService := services.NewUserService(repo)
	userHandler := handlers.NewUserHandler(userService)

	http.HandleFunc("/user", userHandler.GetUser)
	port := ":8080" 
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
