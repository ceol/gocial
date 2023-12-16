package repositories

import (
	"testing"

	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/models"
)

func TestUserCreate(t *testing.T) {
	repo := NewUserRepository(database.DB)

	user := &models.User{}
	repo.Create(user)
	if user.ID == 0 {
		t.Errorf("User not created: %v", user)
	}
}

func TestUserFind(t *testing.T) {
	repo := NewUserRepository(database.DB)

	user := &models.User{UserName: "test_find"}
	repo.Create(user)
	originalID := user.ID

	user = &models.User{ID: originalID}
	err := repo.Find(user)
	if err != nil || user.UserName != "test_find" {
		t.Errorf("User not found: %v [%v]", user, err)
	}
}

func TestUserFindAll(t *testing.T) {
	repo := NewUserRepository(database.DB)

	users := []*models.User{
		{UserName: "test_findall"},
		{UserName: "test_findall"},
		{UserName: "test_findall"},
		{UserName: "test_findall2"},
	}
	for _, user := range users {
		repo.Create(user)
	}

	results, err := repo.FindAll(&models.User{UserName: "test_findall"})
	if err != nil || len(results) != 3 {
		t.Errorf("Users not found: %v [%v]", results, err)
	}
}

func TestUserSave(t *testing.T) {
	repo := NewUserRepository(database.DB)

	originalName := "test_user"
	user := &models.User{UserName: originalName}
	repo.Create(user)
	originalID := user.ID
	originalUpdated := user.UpdatedAt

	user.UserName = "test_changed"
	repo.Save(user)

	user = &models.User{ID: originalID}
	repo.Find(user)
	if user.UserName == originalName {
		t.Errorf("User not updated: %v", user)
	}
	if user.UpdatedAt == originalUpdated {
		t.Errorf("User.UpdatedAt not updated: %v", user)
	}
}

func TestUserDelete(t *testing.T) {
	repo := NewUserRepository(database.DB)

	user := &models.User{}
	repo.Create(user)

	err := repo.Delete(user)
	if err != nil || !user.DeletedAt.Valid {
		t.Errorf("User not deleted: %v [%v]", user, err)
	}
}
