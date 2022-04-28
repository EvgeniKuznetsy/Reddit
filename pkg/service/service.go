package service

import (
	"Reddit/models"
	"Reddit/pkg/repository"
)

type Post interface {
	GetById(id string) (*models.Post, error)
	GetList(page int, limit int) (*models.OutputPostList, error)
	Create(post *models.InputPost) (*models.OutPost, error)
	Update(post *models.InputUpdatesPost) error
	Delete(id string) error
}

type Auth interface {
	SigIn(input *models.InputSingIn) (*models.OutPutIn, error)
	SigUp(input *models.InputSinUp) (*models.OutPutUp, error)
}
type Service struct {
	Post Post
	Auth Auth
}

func NewService(repos *repository.Repositories) *Service {
	return &Service{
		Post: NewPostService(repos.Post),
		Auth: NewAuthService(repos.Auth),
	}
}
