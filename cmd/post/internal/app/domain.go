package app

import "time"

type Post struct {
	ID         int
	UserID     int
	Text       string
	CreatedAt  time.Time
	Comments   []Comment
	BanComment bool
}

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	Text      string
	CreatedAt time.Time
}
