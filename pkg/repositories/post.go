package repositories

import (
	"Reddit/models"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (s *PostPostgres) GetById(id string) (models.Post, error) {

}
func (s *PostPostgres) GetList(page int, limit int) (models.OutputPostList, error) {

}
func (s *PostPostgres) Create(post models.InputPost) (models.OutPost, error) {

}
func (s *PostPostgres) Update(post models.InputUpdatesPost) error {

}
func (s *PostPostgres) Delete(id string) error {

}

type Service struct {
	Post Post
}

func NewService(repos *Repositories) *Service {
	return &Service{Post: post}
}
