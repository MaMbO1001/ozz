package app

import (
	"context"
	"fmt"
)

func (a *App) CreatePost(ctx context.Context, text string) (*Post, error) {
	pst, err := a.repo.CreatePost(ctx, Post{Text: text})
	if err != nil {
		return nil, fmt.Errorf("a.repo.CreatePost: %w", err)
	}
	return pst, nil
}

func (a *App) CreateComment(ctx context.Context, text string) (*Comment, error) {
	comm, err := a.repo.CreateComment(ctx, Comment{Text: text})
	if err != nil {
		return nil, fmt.Errorf("a.repo.CreateComment: %w", err)
	}
	return comm, nil
}

func (a *App) GetPost(ctx context.Context, postID int) (*Post, error) {
	pst, err := a.repo.GetOnePost(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetOnePost: %w", err)
	}
	return pst, nil
}

func (a *App) GetAllPosts(ctx context.Context) ([]Post, error) {
	allPosts, err := a.repo.GetAllPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetAllPosts: %w", err)
	}
	return allPosts, nil
}
