package app

import (
	"context"
	"fmt"
)

func (a *App) CreatePost(ctx context.Context, text string, id int) (*Post, error) {
	pst, err := a.repo.CreatePost(ctx, Post{Text: text, UserID: id})
	if err != nil {
		return nil, fmt.Errorf("a.repo.CreatePost: %w", err)
	}
	return pst, nil
}

func (a *App) CreateComment(ctx context.Context, text string, id int) (*Comment, error) {
	comm, err := a.repo.CreateComment(ctx, Comment{Text: text, UserID: id})
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

func (a *App) GetAllPosts(ctx context.Context, limit, offset int) ([]Post, error) {
	allPosts, err := a.repo.GetAllPosts(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetAllPosts: %w", err)
	}
	return allPosts, nil
}

func (a *App) GetComment(ctx context.Context, commentID int) (*Comment, error) {
	com, err := a.repo.GetComment(ctx, commentID)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetComment: %w", err)
	}
	return com, nil
}

func (a *App) GetAllComments(ctx context.Context, postID int, limit, offset int) ([]Comment, error) {
	coms, err := a.repo.GetAllComments(ctx, postID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetAllComments: %w", err)
	}
	return coms, nil
}

func (a *App) UpdatePost(ctx context.Context, postID int, banComment bool) (*Post, error) {
	pst, err := a.repo.GetOnePost(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetOnePost: %w", err)
	}
	pst.BanComment = banComment
	post, err := a.repo.UpdatePost(ctx, *pst)
	if err != nil {
		return post, fmt.Errorf("a.repo.UpdatePost: %w", err)
	}
	return pst, nil
}
