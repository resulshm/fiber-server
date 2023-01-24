package repository

import "github.com/jmoiron/sqlx"

type VideoPostgres struct {
	db *sqlx.DB
}

func NewVideoPostgres(db *sqlx.DB) *VideoPostgres {
	return &VideoPostgres{db: db}
}
