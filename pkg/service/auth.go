package service

import (
	"Reddit/models"
	"Reddit/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SigIn(input *models.InputSingIn) (*models.OutPutIn, error) {
	return s.repo.SignIn(input)
}

func (s *AuthService) SigUp(input *models.InputSinUp) (*models.OutPutUp, error) {
	return s.repo.SignUp(input)
}
