package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	models "github.com/ecommerce-go/internal/model"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	FetchUser() (*models.User, error)
	GetAll() ([]*models.User, error)
	GetByUserName(username string) (*models.User, error)
	Create(u *models.User) error
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

func (r *userRepo) Create(u *models.User) error {

	query := "INSERT INTO users (name,PasswordHash) VALUES (?,?)"
	res, err := r.db.Exec(query, u.Name, u.PasswordHash)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}
	id, _ := res.LastInsertId()
	u.ID = int(id)
	return nil

}

	func (r *userRepo) GetByUserName(name string) (*models.User, error) {
		u := &models.User{}

		log.Printf("Looking up user by name: %s", name) // Debug input

		query := "SELECT id, name, PasswordHash FROM users WHERE name = ?"
		err := r.db.QueryRow(query, name).Scan(&u.ID, &u.Name, &u.PasswordHash)

		if err != nil {
			log.Printf("Error fetching user %s: %v", name, err)
			return nil, fmt.Errorf("get user: %w", err)
		}

		log.Printf("User found: ID=%d, Name=%s, PasswordHash length=%d", u.ID, u.Name, len(u.PasswordHash))

		return u, nil
	}



func (r *userRepo) FetchUser() (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = 1").Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return &user, nil
}

func (r *userRepo) GetAll() ([]*models.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
