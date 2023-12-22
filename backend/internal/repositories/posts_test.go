package repositories

import (
	"testing"

	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/models"
)

func TestPostSave(t *testing.T) {
	repo := NewPostRepository(database.DB)

	post, err := repo.Save(models.Post{})
	if err != nil || post.ID == 0 {
		t.Errorf("Post not created: %v [%v]", post, err)
	}
}

func TestPostFindByID(t *testing.T) {
	repo := NewPostRepository(database.DB)

	post, _ := repo.Save(models.Post{})
	postID := post.ID

	post, err := repo.FindByID(postID)
	if err != nil || post.ID != postID {
		t.Errorf("Post not found: %v [%v]", post, err)
	}
}

func TestPostFindAllByUserID(t *testing.T) {
	repo := NewPostRepository(database.DB)

	var userID uint = 99
	posts := []models.Post{
		{UserID: userID},
		{UserID: userID},
		{UserID: userID},
		{UserID: 0},
	}
	for _, post := range posts {
		repo.Save(post)
	}

	results, err := repo.FindAllByUserID(userID)
	if err != nil || len(results) != 3 {
		t.Errorf("Posts not found: %v [%v]", results, err)
	}
}

func TestPostDeleteByID(t *testing.T) {
	repo := NewPostRepository(database.DB)

	post, _ := repo.Save(models.Post{})

	err := repo.DeleteByID(post.ID)

	_, findErr := repo.FindByID(post.ID)
	if err != nil || findErr == nil {
		t.Errorf("Post not deleted: %v [%v]", post, err)
	}
}
