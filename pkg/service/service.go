package service

import (
	"context"

	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/ResulShamuhammedov/fiber-server/pkg/repository"
)

type Service struct {
	Video
}

type Video interface {
	AddVideo(ctx context.Context, video models.Video) (int, error)
	GetVideoByID(ctx context.Context, id int) (models.Video, error)
	GetVideoByTitle(ctx context.Context, title string) (models.Video, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Video: NewVideoService(repo.Video),
	}
}
