package services

import (
	"log"

	"github.com/ceol/gocial/internal/models"
	"github.com/ceol/gocial/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepository
}

func (serv UserService) Create(
	userName string,
	email string,
	password string,
) (models.User, error) {
	user := models.User{
		UserName:     userName,
		Email:        email,
		PasswordHash: HashPassword(password),
	}
	return user, serv.repo.Create(&user)
}

func (serv UserService) Update(user models.User) error {
	return serv.repo.Update(&user)
}

func (serv UserService) SetPassword(id uint, newPassword string) error {
	return serv.repo.Update(&models.User{
		ID:           id,
		PasswordHash: HashPassword(newPassword),
	})
}

func (serv UserService) Delete(id uint) error {
	return serv.repo.Delete(&models.User{ID: id})
}

func (serv UserService) FindById(id uint) (models.User, error) {
	user := models.User{ID: id}
	return user, serv.repo.Find(&user)
}

func (serv UserService) FindByUserName(userName string) (models.User, error) {
	user := models.User{UserName: userName}
	return user, serv.repo.Find(&user)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return UserService{repo}
}

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashed)
}

func CheckPasswordHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
