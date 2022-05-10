package service

import (
	"Reddit/models"
	"Reddit/pkg/repository"
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
type Service struct {
	Post          Post
	Auth          Auth
	Session       Session
	RecoverAccess RecoverAccess
}

type RecoverAccess interface {
	GenerateLink(link *models.InputRecoverAccessLink) error
}

func NewService(repos *repository.Repositories) *Service {
	return &Service{
		Post:          NewPostService(repos.Post),
		Auth:          NewAuthService(repos.Auth),
		Session:       NewSessionService(repos.Session),
		RecoverAccess: NewRecoverAccessService(repos.RecoverAccess),
	}
}
