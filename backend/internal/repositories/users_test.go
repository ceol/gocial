package repositories

import (
	"testing"

	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/models"
)

func TestUserSave(t *testing.T) {
	repo := NewUserRepository(database.DB)

	user, err := repo.Save(models.User{})
	if err != nil || user.ID == 0 {
		t.Errorf("User not created: %v [%v]", user, err)
	}
}

func TestUserFindByID(t *testing.T) {
	repo := NewUserRepository(database.DB)

	user, _ := repo.Save(models.User{})
	userID := user.ID

	user, err := repo.FindByID(userID)
	if err != nil || user.ID != userID {
		t.Errorf("User not found: %v [%v]", user, err)
	}
}

func TestUserFindByName(t *testing.T) {
	repo := NewUserRepository(database.DB)

	user, _ := repo.Save(models.User{Name: "testusername"})
	name := user.Name

	user, err := repo.FindByName(name)
	if err != nil || user.Name != name {
		t.Errorf("User not found: %v [%v]", user, err)
	}
}

func TestUserDeleteByID(t *testing.T) {
	repo := NewUserRepository(database.DB)

	user, _ := repo.Save(models.User{})

	err := repo.DeleteByID(user.ID)

	_, findErr := repo.FindByID(user.ID)
	if err != nil || findErr == nil {
		t.Errorf("User not deleted: %v [%v]", user, err)
	}
}
