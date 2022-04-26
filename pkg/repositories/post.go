package repositories

import (
	"Reddit/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
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
	var output models.OutputPostList
	if err := r.db.Select(&output.Post, `select * from "Post" where deleted=false 
                     order by create_date desc limit $1 offset $2`,
		limit, (page-1)*limit); err != nil {
		return nil, err
	}

	if err := r.db.Get(&output.TotalCount, `select count(*) from "Post" where deleted=false`); err != nil {
		return nil, err
	}

	return &output, nil
}

func (r *PostPostgres) Create(post *models.InputPost) (*models.OutPost, error) {
	id := uuid.New().String()
	if id == "" {
		return nil, errors.New("generate void invalid")
	}
	timeNow := time.Now()
	_, err := r.db.Query(`insert into "Post"(id,author,caption,body,create_at,delete_at),
	values($1,$2,$3,$4,$5,$6)`, id, post.Author, post.Body, post.Caption, timeNow, false)
	if err != nil {
		return nil, err
	}
	return &models.OutPost{
		Id:         id,
		CreateDate: timeNow,
	}, nil

}

func (r *PostPostgres) Update(post *models.InputUpdatesPost) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1
	if post.Caption != "" {
		setValues = append(setValues, fmt.Sprintf("caption=$%d", argsId))
		args = append(args, post.Caption)
		argsId++
	}
	if post.Body != "" {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argsId))
		args = append(args, post.Body)
		argsId++
	}
	querySetPart := strings.Join(setValues, ",")
	query := fmt.Sprintf(`UPDATE "Post" SET %s WHERE id=%s`, querySetPart, post.Id)
	_, err := r.db.Query(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostPostgres) Delete(id string) error {

}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}
