package models

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Senha      string `json:"senha"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"user_status"`
}

// var Personalidades []Personalidade
