package main

// import "database/sql"

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"

)

// "fmt"
// "log"


const (
	dbDriver = "mysql"
	dbUser   = "mysql"
	dbPass   = "hZOalxYnBTUO92sRysZqgoj6y4wvw7hBS6Mv9t3RhpxZDa573xLSX1IzhxGDW4an"
	dbHost   = "109.123.234.186"
	dbPort   = "3307"
	dbName   = "default"
)


type User struct {
	ID       int    `json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
}



func dbConn() (*sql.DB, error) {
	return sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
}


func createUserHandler(w http.ResponseWriter, r *http.Request) {

	db, err := dbConn()

	if err != nil {
		 panic(err)
	}

	defer db.Close()

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err = db.Exec(query, user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")

}




func getUserHandler(w http.ResponseWriter, r *http.Request) {		

	db,err := dbConn()
	if err != nil {
		panic(err)
	}	
	defer db.Close()
	
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)


}




func updateUserHandler(w http.ResponseWriter, r *http.Request) {

	db, err := dbConn()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err = db.Exec(query, user.Name, user.Email, id)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	fmt.Fprintln(w, "User updated successfully")
}	

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {	
	
	db, err := dbConn()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	query := "DELETE FROM users WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User deleted successfully")
}


func main() {	
	r := mux.NewRouter()

	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}