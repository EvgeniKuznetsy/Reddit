package repository

import (
	"Reddit/models"
	"github.com/jmoiron/sqlx"

	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db      *sqlx.DB
	session *SessionPostgres
}

func (a *AuthPostgres) SignIn(input *models.InputSingIn) (*models.OutPutIn, error) {
	var output models.OutPutIn
	var account models.Account

	if err := a.db.Get(&account, `select * from "Account" where login=$1 or email=$1`,
		input.Identifier); err != nil {
		return nil, err
	}
	account.Permissions = "admin, moderator, user"

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(input.Password)); err != nil {
		return nil, err
	}

	output.Account = account

	return &output, nil
}

func (a *AuthPostgres) SignUp(input *models.InputSinUp) (*models.OutPutUp, error) {
	var output models.OutPutUp

	passwordHash, err := getPasswordHash(input.Password)
	if err != nil {
		return nil, err
	}

	_, err = a.db.Query(`insert into "Account" (login, password, name, email) 
		values ($1, $2, $3, $4)`, input.Login, passwordHash, input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	sessionHash, err := a.session.Generate(input.Login)
	if err != nil {
		return nil, err
	}

	output.Session = sessionHash
	return &output, nil

}
func getPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db,
		session: NewSessionPostgres(db),
	}
}
