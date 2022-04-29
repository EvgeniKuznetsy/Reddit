package models

import "github.com/golang-jwt/jwt"

type Account struct {
	Email       string `json:"email" db:"email"`
	Name        string `json:"name" db:"name"`
	Login       string `json:"login" db:"login"`
	Password    string `json:"-" db:"password"`
	Permissions string
}
type JwtClaims struct {
	jwt.StandardClaims
	Permissions string
}
