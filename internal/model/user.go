package models


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	PasswordHash string `json:"-"`
}


type UsersResponse struct {
    Value []User `json:"value"`
}
