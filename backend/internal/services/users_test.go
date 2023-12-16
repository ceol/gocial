package services

import (
	"testing"

	"github.com/ceol/gocial/internal/database"
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
