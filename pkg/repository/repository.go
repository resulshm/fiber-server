package repository

import (
	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Video
}

type Video interface {
	AddVideo(video models.Video) (int, error)
	GetVideoByID(id int) (models.Video, error)
	GetVideoByTitle(title string) (models.Video, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Video: NewVideoPostgres(db),
	}
}
