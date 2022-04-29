package models

type Account struct {
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	Password string `json:"-" db:"password"`
}
