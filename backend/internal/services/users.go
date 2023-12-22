package services

import (
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
		UserName: userName,
		Email:    email,
	}
	serv.SetPassword(&user, password)
	return serv.repo.Save(user)
}

func (serv UserService) Save(user models.User) (models.User, error) {
	return serv.repo.Save(user)
}

func (serv UserService) SetPassword(user *models.User, newPassword string) error {
	hashed, err := HashPassword(newPassword)
	if hashed != "" {
		user.PasswordHash = hashed
	}
	return err
}

func (serv UserService) DeleteByID(id uint) error {
	return serv.repo.DeleteByID(id)
}

func (serv UserService) FindByID(id uint) (models.User, error) {
	return serv.repo.FindByID(id)
}

func (serv UserService) FindByUserName(userName string) (models.User, error) {
	return serv.repo.FindByUserName(userName)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return UserService{repo}
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func CheckPasswordHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
