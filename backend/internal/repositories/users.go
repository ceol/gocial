package repositories

import (
	"github.com/ceol/gocial/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*models.User) error
	Update(*models.User) error
	Delete(*models.User) error
	Find(*models.User) error
	FindAll(*models.User) ([]models.User, error)
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

func (repo GormUserRepository) Find(user *models.User) error {
	return repo.DB.Where(user).First(user).Error
}

func (repo GormUserRepository) FindAll(user *models.User) ([]models.User, error) {
	var users []models.User
	return users, repo.DB.Where(user).Find(&users).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	var repo UserRepository = GormUserRepository{DB: db}
	return repo
}
