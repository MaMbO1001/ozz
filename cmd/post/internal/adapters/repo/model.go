package repo

import (
	"github.com/MaMbO1001/cmd/post/internal/app/cmd/post/internal/app"
	"time"
)

type Post struct {
	ID         int       `db:"id" json:"id"`
	UserID     int       `db:"user_id" json:"user_id"`
	Text       string    `db:"text" json:"text"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	BanComment bool      `db:"ban_comment" json:"ban_comment"`
}

type Comment struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	PostID    int       `db:"post_id" json:"post_id"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func convert(p app.Post) *Post {
	return &Post{
		ID:         p.ID,
		Text:       p.Text,
		CreatedAt:  p.CreatedAt,
		BanComment: p.BanComment,
		UserID:     p.UserID,
	}
}

func (p *Post) convert() *app.Post {
	return &app.Post{
		ID:         p.ID,
		Text:       p.Text,
		CreatedAt:  p.CreatedAt,
		BanComment: p.BanComment,
		UserID:     p.UserID,
	}
}

func conv(c app.Comment) *Comment {
	return &Comment{
		ID:        c.ID,
		UserID:    c.UserID,
		PostID:    c.PostID,
		Text:      c.Text,
		CreatedAt: c.CreatedAt,
	}
}

func (c *Comment) convert() *app.Comment {
	return &app.Comment{
		ID:        c.ID,
		UserID:    c.UserID,
		PostID:    c.PostID,
		Text:      c.Text,
		CreatedAt: c.CreatedAt,
	}
}
