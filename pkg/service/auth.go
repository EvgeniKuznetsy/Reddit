package service

import (
	"Reddit/models"
	"Reddit/pkg/repository"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService struct {
	repo repository.Auth
}

const jwtSecret = "69738360beaec81e9c06bd7f785a1a8615522a8db1f50bc0d8ba3438880fc2cf"

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SigIn(input *models.InputSingIn) (*models.OutPutIn, error) {
	output, err := s.repo.SignIn(input)
	if err != nil {
		return nil, err
	}

	token, err := generateJwtToken(output.Account.Permissions)
	output.Token = token

	return output, nil
}
func generateJwtToken(permissions string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		models.JwtClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			},
			Permissions: permissions,
		})
	signingToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signingToken, nil
}

func (s *AuthService) SigUp(input *models.InputSinUp) (*models.OutPutUp, error) {
	return s.repo.SignUp(input)
}
