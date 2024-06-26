package app

import "context"

type (
	Repo interface {
		CreatePost(context.Context, Post) (*Post, error)
		CreateComment(context.Context, Comment) (*Comment, error)
		GetAllPosts(context.Context, int, int) ([]Post, error)
		GetOnePost(context.Context, int) (*Post, error)
		GetAllComments(context.Context, int, int, int) ([]Comment, error)
		GetComment(context.Context, int) (*Comment, error)
		UpdatePost(context.Context, Post) (*Post, error)
	}
)
