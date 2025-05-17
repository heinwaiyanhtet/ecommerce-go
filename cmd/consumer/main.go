// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	_ "github.com/go-sql-driver/mysql"

// 	repository "github.com/heinwaiyanhtet/ecommerce-go/internal/repository"
// 	service "github.com/heinwaiyanhtet/ecommerce-go/internal/service"
// )

// func main() {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Failed to load .env: %v", err)
// 	}

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to DB: %v", err)
// 	}
// 	defer db.Close()

// 	orderRepo := repository.NewOrderRepository(db)

// 	rabbitURL := os.Getenv("RABBITMQ_URL")
// 	if rabbitURL == "" {
// 		rabbitURL = "amqp://guest:guest@localhost:5672/"
// 	}

// 	consumer, err := service.NewOrderConsumer(rabbitURL, orderRepo)
// 	if err != nil {
// 		log.Fatalf("Failed to create order consumer: %v", err)
// 	}

// 	err = consumer.StartConsuming()
// 	if err != nil {
// 		log.Fatalf("Consumer failed to start: %v", err)
// 	}

// 	select {} // block forever to keep consuming
// }
