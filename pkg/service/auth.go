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

func (s *AuthService) SignIn(input *models.InputSignIn) (*models.OutPutSignIn, error) {
	return s.repo.SignIn(input)
}

func (s *AuthService) SignUp(input *models.InputSignUp) (*models.OutPutSignUp, error) {
	return s.repo.SignUp(input)
}
