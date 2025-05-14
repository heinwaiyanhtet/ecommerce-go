package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	handlers "github.com/ecommerce-go/internal/handler"
	repositories "github.com/ecommerce-go/internal/repository"
	services "github.com/ecommerce-go/internal/service"
	"github.com/joho/godotenv"
)

func main() {

		err := godotenv.Load()

		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
		fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))

		repo, err := repositories.NewUserRepo();

		authSvc := services.NewAuthService(repo, os.Getenv("JWT_SECRET"), 24*time.Hour)
		authHandler := handlers.NewAuthHandler(authSvc)	


		mux := http.NewServeMux()
		mux.HandleFunc("/signup", authHandler.Signup)
		mux.HandleFunc("/login", authHandler.Login)

		protected := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Secret data"))
		})

		mux.Handle("/secret", handlers.JWTMiddleware([]byte(os.Getenv("JWT_SECRET")))(protected))

		userService := services.NewUserService(repo)
		userHandler := handlers.NewUserHandler(userService)

		mux.HandleFunc("/users", userHandler.GetAllUsers)

		port := ":8080"
		fmt.Printf("Starting server on %s\n", port)
		log.Fatal(http.ListenAndServe(port, mux))
}
