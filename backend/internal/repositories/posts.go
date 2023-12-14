package repositories

import (
	"github.com/ceol/gocial/internal/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(*models.Post) error
	Update(*models.Post) error
	Delete(*models.Post) error
	Find(*models.Post) error
	FindAllByUser(uint) ([]models.Post, error)
}

type GormPostRepository struct {
	DB *gorm.DB
}

func (repo GormPostRepository) Create(post *models.Post) error {
	return repo.DB.Create(post).Error
}

func (repo GormPostRepository) Update(post *models.Post) error {
	return repo.DB.Save(post).Error
}

func (repo GormPostRepository) Delete(post *models.Post) error {
	return repo.DB.Delete(post).Error
}

func (repo GormPostRepository) Find(post *models.Post) error {
	return repo.DB.First(post).Error
}

func (repo GormPostRepository) FindAllByUser(userId uint) ([]models.Post, error) {
	posts := []models.Post{}
	return posts, repo.DB.Where("user_id = ?", userId).Error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return GormPostRepository{DB: db}
}
