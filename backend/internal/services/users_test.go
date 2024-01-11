package services

import (
	"fmt"
	"testing"

	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/models"
	"github.com/ceol/gocial/internal/repositories"
)

func TestUserCreate(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	tests := []struct {
		name     string
		email    string
		password string
	}{
		{"test", "test@test.com", "test"},
		{"", "", ""},
		{"👏🥳🍰", "🫥🎉🐶", "🙏🥱🫡"},
	}
	for _, tt := range tests {
		user, err := serv.Create(tt.name, tt.email, tt.password)
		if err != nil || user.ID == 0 {
			t.Errorf("Create(%v) failed: %v [%v]", tt, user, err)
		}
	}
}

func TestUserSetPassword(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	tests := []string{
		"test password",
		"😂👍👾🦄",
	}
	for _, password := range tests {
		user := &models.User{}
		err := serv.SetPassword(user, password)
		if err != nil || !CheckPasswordHash(password, user.PasswordHash) {
			t.Errorf("SetPassword(%v, %v) failed: %v", user, password, err)
		}
	}
}

func TestUserFindByID(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	name := "testuserFindByID"
	user, _ := serv.Create(name, fmt.Sprintf("%v@test.com", name), "test password")
	userID := user.ID

	user, err := serv.FindByID(userID)
	if err != nil || user.ID != userID {
		t.Errorf("FindByID(%v) failed: got %v [%v]", userID, user, err)
	}
}

func TestUserFindByName(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	name := "testuserfindbyusername"
	user, _ := serv.Create(name, fmt.Sprintf("%v@test.com", name), "test password")

	user, err := serv.FindByName(name)
	if err != nil || user.Name != name {
		t.Errorf("FindByName(%v) failed: got %v [%v]", name, user, err)
	}
}

func TestUserSave(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	user, _ := serv.Create("testusersave", "testusersave@test.com", "test password")

	testChange := "testusersave2"
	user.Name = testChange
	user, err := serv.Save(user)

	user, _ = serv.FindByID(user.ID)
	if err != nil || user.Name != testChange {
		t.Errorf("Save(%v) failed: [%v]", user, err)
	}
}

func TestUserDeleteByID(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	name := "testuserdelete"
	user, _ := serv.Create(name, fmt.Sprintf("%v@test.com", name), "test password")

	err := serv.DeleteByID(user.ID)

	_, findErr := serv.FindByID(user.ID)
	if err != nil || findErr == nil {
		t.Errorf("DeleteByID(%v) failed: [%v]", user, err)
	}
}
