package repository

import (
	"context"
	"database/sql"

	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	insertVideo      = `INSERT INTO videos (title, description) VALUES ($1, $2) RETURNING id`
	getVideoByID     = `SELECT * FROM videos WHERE id = $1`
	getVideoByTitile = `SELECT * FROM videos WHERE title = $1`
)

type VideoPostgres struct {
	db *pgxpool.Pool
}

func NewVideoPostgres(db *pgxpool.Pool) *VideoPostgres {
	return &VideoPostgres{db: db}
}

func (r *VideoPostgres) AddVideo(ctx context.Context, video models.Video) (int, error) {
	var id int
	row := r.db.QueryRow(ctx, insertVideo, video.Title, video.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *VideoPostgres) GetVideoByID(ctx context.Context, id int) (models.Video, error) {
	var video models.Video
	row := r.db.QueryRow(ctx, getVideoByID, id)
	if err := row.Scan(&video); err != nil {
		if err == sql.ErrNoRows {
			return models.Video{}, nil
		}
		return video, err
	}

	return video, nil
}

func (r *VideoPostgres) GetVideoByTitle(ctx context.Context, title string) (models.Video, error) {
	var video models.Video
	row := r.db.QueryRow(ctx, getVideoByTitile, title)
	if err := row.Scan(&video); err != nil {
		if err == sql.ErrNoRows {
			return models.Video{}, nil
		}
		return video, err
	}

	return video, nil
}
