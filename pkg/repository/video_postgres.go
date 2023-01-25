package repository

import (
	"database/sql"

	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/jmoiron/sqlx"
)

const (
	insertVideo      = `INSERT INTO videos (title, description) VALUES ($1, $2) RETURNING id`
	getVideoByID     = `SELECT * FROM videos WHERE id = $1`
	getVideoByTitile = `SELECT * FROM videos WHERE title = $1`
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

func (r *VideoPostgres) GetVideoByID(id int) (models.Video, error) {
	var video models.Video
	if err := r.db.Get(&video, getVideoByID, id); err != nil {
		if err == sql.ErrNoRows {
			return models.Video{}, nil
		}
		return video, err
	}

	return video, nil
}

func (r *VideoPostgres) GetVideoByTitle(title string) (models.Video, error) {
	var video models.Video
	if err := r.db.Get(&video, getVideoByTitile, title); err != nil {
		if err == sql.ErrNoRows {
			return models.Video{}, nil
		}
		return video, err
	}

	return video, nil
}
