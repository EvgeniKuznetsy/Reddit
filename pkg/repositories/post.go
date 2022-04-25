package repositories

import (
	"Reddit/models"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func (r *PostPostgres) GetById(id string) (*models.Post, error) {
	var post *models.Post

	if err := r.db.Get(&post, `select * from "Post" where id=$1`, id); err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostPostgres) GetList(page int, limit int) (*models.OutputPostList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostPostgres) Create(post *models.InputPost) (*models.OutPost, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostPostgres) Update(post *models.InputUpdatesPost) error {
	//TODO implement me
	panic("implement me")
}

func (r *PostPostgres) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}
