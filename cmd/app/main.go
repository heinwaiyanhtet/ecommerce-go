package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	handlers "github.com/heinwaiyanhtet/ecommerce-go/internal/handler"
	repositories "github.com/heinwaiyanhtet/ecommerce-go/internal/repository"
	services "github.com/heinwaiyanhtet/ecommerce-go/internal/service"
	"github.com/heinwaiyanhtet/ecommerce-go/pkg/rabbitmq"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))

	// Build DSN and connect to DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

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
	// Connect to RabbitMQ
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/"
	}

	rabbitConn, rabbitChannel, err := rabbitmq.Connect(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	defer rabbitConn.Close()
	defer rabbitChannel.Close()

	orderPublisher, err := services.NewOrderPublisher(rabbitChannel)
	if err != nil {
		log.Fatalf("Failed to create order publisher: %v", err)
	}

	// Initialize services
	authSvc := services.NewAuthService(userRepo, os.Getenv("JWT_SECRET"), 24*time.Hour)
	userService := services.NewUserService(userRepo)

	orderService := services.NewOrderService(orderRepo, orderPublisher)

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
