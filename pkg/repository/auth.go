package repository

import (
	"Reddit/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (a *AuthPostgres) SignIn(input *models.InputSingIn) (*models.OutPutIn, error) {

}

func (a *AuthPostgres) SignUp(input *models.InputSinUp) (*models.OutPutIn, error) {

}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
