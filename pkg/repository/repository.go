package repository

import (
	"Reddit/models"

	"github.com/jmoiron/sqlx"
)

type Post interface {
	GetById(id string) (*models.Post, error)
	GetList(page int, limit int) (*models.OutputPostList, error)
	Create(post *models.InputPost) (*models.OutputPost, error)
	Update(post *models.InputUpdatesPost) error
	Delete(id string) error
}
type Auth interface {
	SignIn(input *models.InputSignIn) (*models.OutPutSignIn, error)
	SignUp(input *models.InputSignUp) (*models.OutPutSignUp, error)
}
type Session interface {
	Generate(login string) (string, error)
	GetAccount(hash string) (*models.Account, error)
}
type RecoverAccess interface {
	GenerateLink(input *models.InputRecoverAccessLink) (string, string, error)
}

type Repositories struct {
	Post          Post
	Auth          Auth
	Session       Session
	RecoverAccess RecoverAccess
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		NewPostPostgres(db),
		NewAuthPostgres(db),
		Session:       NewSessionPostgres(db),
		RecoverAccess: NewRecoverAccessPostgres(db),
	}
}
