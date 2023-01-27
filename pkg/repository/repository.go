package repository

import (
	"context"

	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Video
}

type Video interface {
	AddVideo(ctx context.Context, video models.Video) (int, error)
	GetVideoByID(ctx context.Context, id int) (models.Video, error)
	GetVideoByTitle(ctx context.Context, title string) (models.Video, error)
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Video: NewVideoPostgres(db),
	}
}
