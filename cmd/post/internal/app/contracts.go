package app

import "context"

type (
	Repo interface {
		CreatePost(context.Context, Post) (*Post, error)
		CreateComment(context.Context, Comment) (*Comment, error)
		GetAllPosts(context.Context) ([]Post, error)
		GetOnePost(context.Context, int) (*Post, error)
	}
)
