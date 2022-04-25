package services

import (
	"Reddit/models"
	"Reddit/pkg/repositories"
)

type Post interface {
	GetById(id string) (models.Post, error)
	GetList(page int, limit int) (models.OutputPostList, error)
	Create(post models.InputPost) (models.OutPost, error)
	Update(post models.InputUpdatesPost) error
	Delete(id string) error
}
type Service struct {
	Post Post
}

func NewService(repos *repositories.Repositories) *Service {
	return &Service{
		Post: NewPostService(repos.Post),
	}
}
