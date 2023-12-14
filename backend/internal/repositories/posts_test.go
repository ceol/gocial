package repositories

import (
	"testing"

	"github.com/ceol/gocial/internal/database"
	"github.com/ceol/gocial/internal/models"
)

func TestPostCreate(t *testing.T) {
	repo := NewPostRepository(database.DB)

	post := &models.Post{}
	repo.Create(post)
	if post.ID == 0 {
		t.Errorf("Post not created: %v", post)
	}
}

func TestPostFind(t *testing.T) {
	repo := NewPostRepository(database.DB)

	testContent := "test find"

	post := &models.Post{Content: testContent}
	repo.Create(post)

	post = &models.Post{Content: testContent}
	err := repo.Find(post)
	if err != nil || post.Content != testContent {
		t.Errorf("Post not found: %v [%v]", post, err)
	}
}

func TestPostFindAllByUser(t *testing.T) {
	userRepo := NewUserRepository(database.DB)
	user := &models.User{}
	userRepo.Create(user)

	postRepo := NewPostRepository(database.DB)
	posts := []*models.Post{
		{UserID: user.ID},
		{UserID: user.ID},
		{UserID: user.ID},
		{UserID: 0},
	}
	for _, post := range posts {
		postRepo.Create(post)
	}

	results, err := postRepo.FindAllByUser(user.ID)
	if err != nil || len(results) != 3 {
		t.Errorf("Posts not found: %v [%v]", results, err)
	}
}

func TestPostUpdate(t *testing.T) {
	repo := NewPostRepository(database.DB)

	post := &models.Post{}
	repo.Create(post)

	testContent := "test update"
	post = &models.Post{ID: post.ID, Content: testContent}
	err := repo.Update(post)
	if err != nil {
		t.Errorf("Error updating post: %v [%v]", post, err)
	}

	post = &models.Post{ID: post.ID}
	repo.Find(post)
	if post.Content != testContent {
		t.Errorf("Post not updated: %v", post)
	}
}

func TestPostDelete(t *testing.T) {
	repo := NewPostRepository(database.DB)

	post := &models.Post{}
	repo.Create(post)

	err := repo.Delete(post)
	if err != nil || !post.DeletedAt.Valid {
		t.Errorf("Post not deleted: %v [%v]", post, err)
	}
}
