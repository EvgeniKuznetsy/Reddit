package service

import (
	"Reddit/models"
	"Reddit/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func (a *AuthService) SigIn(input *models.InputSingIn) *models.OutPutIn {

}

func (a *AuthService) SigUp(input *models.InputSinUp) (*models.OutPutUp, *models.OutPutUp) {

}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}
