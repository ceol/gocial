package services

import (
	"github.com/ceol/gocial/internal/models"
	"github.com/ceol/gocial/internal/repositories"
)

type PostService struct {
	repo repositories.PostRepository
}

func (serv PostService) Create(userID uint, content string) (models.Post, error) {
	post := models.Post{UserID: userID, Content: content}
	return serv.repo.Save(post)
}

func (serv PostService) Save(post models.Post) (models.Post, error) {
	return serv.repo.Save(post)
}

func (serv PostService) DeleteByID(id uint) error {
	return serv.repo.DeleteByID(id)
}

func (serv PostService) FindByID(id uint) (models.Post, error) {
	return serv.repo.FindByID(id)
}

func (serv PostService) FindAllByUserID(userID uint) ([]models.Post, error) {
	return serv.repo.FindAllByUserID(userID)
}

func NewPostService(repo repositories.PostRepository) PostService {
	return PostService{repo: repo}
}
