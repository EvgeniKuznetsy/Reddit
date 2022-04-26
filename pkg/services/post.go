package services

import (
	"Reddit/models"
	"Reddit/pkg/repositories"
)

type PostService struct {
	repo repositories.Post
}

func NewPostService(repo repositories.Post) *PostService {
	return &PostService{repo: repo}
}
func (s *PostService) GetById(id string) (*models.Post, error) {
	return s.repo.GetById(id)
}
func (s *PostService) GetList(page int, limit int) (*models.OutputPostList, error) {
	return s.repo.GetList(page, limit)
}
func (s *PostService) Create(post *models.InputPost) (*models.OutPost, error) {
	return s.repo.Create(post)
}
func (s *PostService) Update(post *models.InputUpdatesPost) error {
	return s.repo.Update(post)
}
func (s *PostService) Delete(id string) error {
	return s.repo.Delete(id)
}
