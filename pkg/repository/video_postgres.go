package repository

import (
	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/jmoiron/sqlx"
)

const (
	insertVideo = `INSERT INTO videos (title, description) VALUES ($1, $2) RETURNING id`
)

type VideoPostgres struct {
	db *sqlx.DB
}

func NewVideoPostgres(db *sqlx.DB) *VideoPostgres {
	return &VideoPostgres{db: db}
}

func (r *VideoPostgres) AddVideo(video models.Video) (int, error) {
	var id int
	row := r.db.QueryRow(insertVideo, video.Title, video.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
