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

}

func TestPostFindAllByUser(t *testing.T) {

}

func TestPostUpdate(t *testing.T) {

}

func TestPostDelete(t *testing.T) {

}
