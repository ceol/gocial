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
		userName string
		email    string
		password string
	}{
		{"test", "test@test.com", "test"},
		{"", "", ""},
		{"👏🥳🍰", "🫥🎉🐶", "🙏🥱🫡"},
	}
	for _, tt := range tests {
		user, err := serv.Create(tt.userName, tt.email, tt.password)
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

func TestUserFind(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	userName := "testuserfind"
	user, _ := serv.Create(userName, fmt.Sprintf("%v@test.com", userName), "test password")

	user = models.User{UserName: user.UserName, Email: user.Email}
	err := serv.Find(&user)
	if err != nil || user.ID == 0 {
		t.Errorf("Find(%v) failed: [%v]", user, err)
	}
}

func TestUserFindByID(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	userName := "testuserFindByID"
	user, _ := serv.Create(userName, fmt.Sprintf("%v@test.com", userName), "test password")
	userID := user.ID

	user, err := serv.FindByID(userID)
	if err != nil || user.ID != userID {
		t.Errorf("FindByID(%v) failed: got %v [%v]", userID, user, err)
	}
}

func TestUserFindByUserName(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	userName := "testuserfindbyusername"
	user, _ := serv.Create(userName, fmt.Sprintf("%v@test.com", userName), "test password")

	user, err := serv.FindByUserName(userName)
	if err != nil || user.UserName != userName {
		t.Errorf("FindByUserName(%v) failed: got %v [%v]", userName, user, err)
	}
}

func TestUserSave(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	user, _ := serv.Create("testusersave", "testusersave@test.com", "test password")

	testChange := "testusersave2"
	user.UserName = testChange
	err := serv.Save(&user)

	user, _ = serv.FindByID(user.ID)
	if err != nil || user.UserName != testChange {
		t.Errorf("Save(%v) failed: [%v]", user, err)
	}
}

func TestUserDelete(t *testing.T) {
	repo := repositories.NewUserRepository(database.DB)
	serv := NewUserService(repo)

	userName := "testuserdelete"
	user, _ := serv.Create(userName, fmt.Sprintf("%v@test.com", userName), "test password")

	err := serv.Delete(&user)

	if err != nil || !user.DeletedAt.Valid {
		t.Errorf("Delete(%v) failed: [%v]", user, err)
	}
}
