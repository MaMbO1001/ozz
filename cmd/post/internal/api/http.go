package api

import (
	"context"
	"github.com/MaMbO1001/cmd/post/internal/app/cmd/post/internal/app"
)

type Application interface {
	CreatePost(context.Context, string) (*app.Post, error)
	CreateComment(context.Context, Comment) (*Comment, error)
	GetAllPosts(context.Context, int, int) ([]Post, error)
	GetOnePost(context.Context, int) (*Post, error)
	GetAllComments(context.Context, int, int, int) ([]Comment, error)
	GetComment(context.Context, int) (*Comment, error)
	UpdatePost(context.Context, Post) (*Post, error)
}
