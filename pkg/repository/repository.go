package repository

import (
	"Reddit/models"

	"github.com/jmoiron/sqlx"
)

type Post interface {
	GetById(id string) (*models.Post, error)
	GetList(page int, limit int) (*models.OutputPostList, error)
	Create(post *models.InputPost) (*models.OutPost, error)
	Update(post *models.InputUpdatesPost) error
	Delete(id string) error
}
type Auth interface {
	SignIn(input *models.InputSignIn) (*models.OutPutSignIn, error)
	SignUp(input *models.InputSignUp) (*models.OutPutSignUp, error)
}

type Repositories struct {
	Post Post
	Auth Auth
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		NewPostPostgres(db),
		NewAuthPostgres(db),
	}
}
