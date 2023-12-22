package repositories

import (
	"github.com/ceol/gocial/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(models.User) (models.User, error)
	DeleteByID(uint) error
	FindByID(uint) (models.User, error)
	FindByUserName(string) (models.User, error)
}

type GormUserRepository struct {
	DB *gorm.DB
}

func (repo GormUserRepository) Save(user models.User) (models.User, error) {
	return user, repo.DB.Save(&user).Error
}

func (repo GormUserRepository) DeleteByID(id uint) error {
	return repo.DB.Delete(&models.User{ID: id}).Error
}

func (repo GormUserRepository) FindByID(id uint) (models.User, error) {
	user := models.User{}
	return user, repo.DB.Where("id = ?", id).Take(&user).Error
}

func (repo GormUserRepository) FindByUserName(userName string) (models.User, error) {
	user := models.User{}
	return user, repo.DB.Where("user_name = ?", userName).Take(&user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	var repo UserRepository = GormUserRepository{DB: db}
	return repo
}
