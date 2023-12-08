package main

import (
	"log"

	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/models"
	"github.com/ceol/gocial/internal/repositories"
	"github.com/ceol/gocial/internal/services"
)

func main() {
	// server.Start()
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	repo := repositories.NewUserRepository(database.DB)
	serv := services.NewUserService(repo)

	user, _ := serv.Create("test", "test", "test")
	log.Println(user)
}
