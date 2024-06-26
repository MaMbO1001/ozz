package repo

import (
	"github.com/MaMbO1001/cmd/post/internal/app/cmd/post/internal/app"
	"math/rand"
	"testing"
	"time"
)

func TestRepo_Smoke(t *testing.T) {
	t.Parallel()

	post := app.Post{
		ID:         rand.Int(),
		UserID:     rand.Int(),
		Text:       "test post",
		CreatedAt:  time.Now(),
		BanComment: bool(false),
	}

	comment := app.Comment{
		PostID:    post.ID,
		UserID:    rand.Int(),
		ID:        rand.Int(),
		Text:      "test comment",
		CreatedAt: time.Now(),
	}
	createPost, err := Repo
}
