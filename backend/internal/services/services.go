package services

import (
	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/repositories"
)

var userRepo repositories.UserRepository
var User UserService

func init() {
	userRepo = repositories.NewUserRepository(database.DB)
	User = NewUserService(userRepo)
}
