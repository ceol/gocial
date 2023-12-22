package services

import (
	"testing"

	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/models"
	"github.com/ceol/gocial/internal/repositories"
)

func TestPostCreate(t *testing.T) {
	repo := repositories.NewPostRepository(database.DB)
	serv := NewPostService(repo)

	post, err := serv.Create(1, "test content")
	if err != nil || post.ID == 0 {
		t.Errorf("Create failed: %v [%v]", post, err)
	}
}

func TestPostFindByID(t *testing.T) {
	repo := repositories.NewPostRepository(database.DB)
	serv := NewPostService(repo)

	created, _ := serv.Create(1, "test content")

	post, err := serv.FindByID(created.ID)
	if err != nil || post.ID != created.ID {
		t.Errorf("FindByID(%v) failed: %v [%v]", created.ID, post, err)
	}
}

func TestPostFindAllByUserID(t *testing.T) {
	repo := repositories.NewPostRepository(database.DB)
	serv := NewPostService(repo)

	var userID uint = 99
	posts := []models.Post{
		{UserID: userID},
		{UserID: userID},
		{UserID: userID},
		{UserID: 0},
	}
	for _, post := range posts {
		serv.Save(post)
	}

	results, err := serv.FindAllByUserID(userID)
	if err != nil || len(results) != 3 {
		t.Errorf("Posts not found: %v [%v]", results, err)
	}
}

func TestPostSave(t *testing.T) {
	repo := repositories.NewPostRepository(database.DB)
	serv := NewPostService(repo)

	post, _ := serv.Create(1, "test content")

	testChange := "test content 2"
	post.Content = testChange
	post, err := serv.Save(post)

	post, _ = serv.FindByID(post.ID)
	if err != nil || post.Content != testChange {
		t.Errorf("Save(%v) failed: [%v]", post, err)
	}
}

func TestPostDeleteByID(t *testing.T) {
	repo := repositories.NewPostRepository(database.DB)
	serv := NewPostService(repo)

	post, _ := serv.Create(1, "test content")

	err := serv.DeleteByID(post.ID)

	_, findErr := serv.FindByID(post.ID)
	if err != nil || findErr == nil {
		t.Errorf("DeleteByID(%v) failed: [%v]", post, err)
	}
}
