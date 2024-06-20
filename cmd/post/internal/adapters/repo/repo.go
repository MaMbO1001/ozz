package repo

import (
	"context"
	"database/sql"
)

type (
	Repo struct {
		DB *sql.DB
	}
)

func New(ctx context.Context, namespace string) (*Repo, error) {
	db, err := sql.Open("postgres", namespace)
	if err != nil {
		panic(err)
	}
	return &Repo{db}, nil
}
func (r *Repo) Close() error {
	if r.DB != nil {
		return r.DB.Close()
	}
	return nil
}

func (r *Repo) CreatePost(ctx context.Context, p app.Post) (post *app.Post, err error {

}
