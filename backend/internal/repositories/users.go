package repositories

import (
	"github.com/ceol/gocial/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
	Find(id uint) (*models.User, error)
}

type GormUserRepository struct {
	DB *gorm.DB
}

func (repo GormUserRepository) Create(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo GormUserRepository) Update(user *models.User) error {
	return repo.DB.Save(user).Error
}

func (repo GormUserRepository) Delete(user *models.User) error {
	return repo.DB.Delete(user).Error
}

func (repo GormUserRepository) Find(id uint) (*models.User, error) {
	user := &models.User{}
	return user, repo.DB.First(user, "id = ?", id).Error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	var repo UserRepository = GormUserRepository{DB: db}
	return &repo
}
