package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Video
}

type Video interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Video: NewVideoPostgres(db),
	}
}
