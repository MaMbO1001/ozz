package repo

import "time"

type Post struct {
	ID         int       `db:"id" json:"id"`
	UserID     int       `db:"user_id" json:"user_id"`
	Text       string    `db:"text" json:"text"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	Comments   []Comment `db:"comments" json:"comments"`
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
		ID: p.
	}
}
