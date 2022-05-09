package repository

import (
	"Reddit/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db *sqlx.DB
}

const (
	AccountRoleAdmin     = "admin"
	AccountRoleModerator = "moderator"
	AccountRoleUser      = "user"
)

func (r *AuthPostgres) SignIn(input *models.InputSignIn) (*models.OutPutSignIn, error) {
	var output models.OutPutSignIn
	var account models.Account

	if err := r.db.Get(&account, `select * from "Account" where login=$1 or email=$1`,
		input.Identifier); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(input.Password)); err != nil {
		return nil, err
	}

	sessionHash, err := r.session.Generate(account.Login)
	if err != nil {
		return nil, err
	}

	output.Session = sessionHash
	output.Account = account

	return &output, nil
}

func (r *AuthPostgres) SignUp(input *models.InputSignUp) (*models.OutPutSignIn, error) {

	passwordHash, err := getPasswordHash(input.Password)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Query(`insert into "Account" (login, password, name, email,role,create_date) 
		values ($1, $2, $3, $4,5,current_timestamp )`, input.Login, passwordHash, input.Name, input.Email, AccountRoleUser)
	if err != nil {
		return err
	}

	return nil

}
func getPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}
