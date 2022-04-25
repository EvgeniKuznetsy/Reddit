package repositories

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

type Repositories struct {
	Post Post
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Post: NewPostPostgres(db),
	}
}
