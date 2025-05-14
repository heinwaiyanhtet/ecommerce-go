package repositories

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/ecommerce-go/internal/model"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	FetchUser() (*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo() (*userRepo, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return &userRepo{db: db}, nil

}

func (r *userRepo) FetchUser() (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = 1").Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return &user, nil
}