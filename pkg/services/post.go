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
func (s *PostService) GetById(id string) (models.Post, error) {

}
func (s *PostService) GetList(page int, limit int) (models.OutputPostList, error) {

}
func (s *PostService) Create(post models.InputPost) (models.OutPost, error) {

}
func (s *PostService) Update(post models.InputUpdatesPost) error {

}
func (s *PostService) Delete(id string) error {

}
