package repo

import (
	"context"
	"fmt"
	"github.com/MaMbO1001/cmd/post/internal/app/cmd/post/internal/app"
	"github.com/jmoiron/sqlx"
)

var _ app.Repo = &Repo{}

type (
	Repo struct {
		sql *sqlx.DB
	}
)

func New(ctx context.Context, dsn string) (*Repo, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	return &Repo{db}, nil
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) CreatePost(ctx context.Context, p app.Post) (post *app.Post, err error) {
	newPost := convert(p)
	const query = `
insert into 
post
(text, user_id)
values ($1, $2) returning *`
	pst := Post{}
	err = r.sql.GetContext(ctx, &pst, query, newPost.Text, newPost.UserID)
	if err != nil {
		return nil, fmt.Errorf("db.GetContext: %w", err)
	}
	post = pst.convert()
	return post, nil
}

func (r *Repo) CreateComment(ctx context.Context, c app.Comment) (comment *app.Comment, err error) {
	newCom := conv(c)
	const query = `
insert into
comment
(text, user_id)
values ($1, $2) returning *`
	com := Comment{}
	err = r.sql.GetContext(ctx, &com, query, newCom.Text, newCom.UserID)
	if err != nil {
		return nil, fmt.Errorf("db.GetContext: %w", err)
	}
	comment = com.convert()
	return comment, nil
}

func (r *Repo) GetOnePost(ctx context.Context, id int) (post *app.Post, err error) {
	pst := Post{}
	const query = `
select * from post where id = $1`
	err = r.sql.GetContext(ctx, &pst, query, id)
	if err != nil {
		return nil, fmt.Errorf("db.GetContext: %w", err)
	}
	return pst.convert(), nil
}

func (r *Repo) GetComment(ctx context.Context, id int) (comment *app.Comment, err error) {
	com := Comment{}
	const query = `
select * from comment where id = $1`
	err = r.sql.GetContext(ctx, &com, query, id)
	if err != nil {
		return nil, fmt.Errorf("db.GetContext: %w", err)
	}
	return com.convert(), nil
}

func (r *Repo) GetAllPosts(ctx context.Context, limit, offset int) (posts []app.Post, err error) {
	const query = `
select * from post ORDER BY created_at ASC LIMIT $1 OFFSET $2`
	pos := []Post{}
	err = r.sql.SelectContext(ctx, &pos, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("db.SelectContext: %w", err)
	}
	return posts, nil
}

func (r *Repo) GetAllComments(ctx context.Context, postID int, limit, offset int) (comments []app.Comment, err error) {
	const query = `
select * from comment where post_id = $1
limit $2 offset $3`
	coms := []Comment{}
	err = r.sql.SelectContext(ctx, &coms, query, postID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("db.SelectContext: %w", err)
	}
	return comments, nil
}

func (r *Repo) UpdatePost(ctx context.Context, post app.Post) (upPost *app.Post, err error) {
	updPost := convert(post)
	const query = `
update post
set
ban_comment = $1
where id = $2
returning *`
	var pst Post
	err = r.sql.GetContext(ctx, &pst, query, updPost.Text, updPost.UserID)
	if err != nil {
		return nil, fmt.Errorf("db.GetContext: %w", err)
	}
	upPost = pst.convert()
	return upPost, nil
}
