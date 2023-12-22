package repositories

import (
	"github.com/ceol/gocial/internal/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	DeleteByID(uint) error
	FindByID(uint) (models.Post, error)
	FindAllByUserID(uint) ([]models.Post, error)
}

type GormPostRepository struct {
	DB *gorm.DB
}

func (repo GormPostRepository) Save(post models.Post) (models.Post, error) {
	return post, repo.DB.Save(&post).Error
}

func (repo GormPostRepository) DeleteByID(id uint) error {
	return repo.DB.Delete(&models.Post{ID: id}).Error
}

func (repo GormPostRepository) FindByID(id uint) (models.Post, error) {
	post := models.Post{}
	return post, repo.DB.Where("id = ?", id).Take(&post).Error
}

func (repo GormPostRepository) FindAllByUserID(userID uint) ([]models.Post, error) {
	posts := []models.Post{}
	return posts, repo.DB.Where("user_id = ?", userID).Find(&posts).Error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return GormPostRepository{DB: db}
}
