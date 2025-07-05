package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	handlers "github.com/heinwaiyanhtet/ecommerce-go/internal/handler"
	repositories "github.com/heinwaiyanhtet/ecommerce-go/internal/repository"
	services "github.com/heinwaiyanhtet/ecommerce-go/internal/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))

	// // Build DSN and connect to DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("DB_HOST:", dsn)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	userRepo, err := repositories.NewUserRepo()
	if err != nil {
		log.Fatalf("Error creating user repository: %v", err)
	}

	orderRepo := repositories.NewOrderRepository(db)

	// Initialize services
	authSvc := services.NewAuthService(userRepo, os.Getenv("JWT_SECRET"), 24*time.Hour)
	userService := services.NewUserService(userRepo)

	orderService := services.NewOrderService(orderRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authSvc)
	userHandler := handlers.NewUserHandler(userService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", authHandler.Signup)
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/users", userHandler.GetAllUsers)
	mux.HandleFunc("/orders", orderHandler.CreateOrder)

	// Protected route example
	protected := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Secret data"))
	})

	mux.Handle("/secret", handlers.JWTMiddleware([]byte(os.Getenv("JWT_SECRET")))(protected))

	port := ":8080"
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))

}
