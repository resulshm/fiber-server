package service

import (
	"context"

	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/ResulShamuhammedov/fiber-server/pkg/repository"
)

type VideoService struct {
	repo repository.Video
}

func NewVideoService(repo repository.Video) *VideoService {
	return &VideoService{repo: repo}
}

func (s *VideoService) AddVideo(ctx context.Context, video models.Video) (int, error) {
	return s.repo.AddVideo(ctx, video)
}

func (s *VideoService) GetVideoByID(ctx context.Context, id int) (models.Video, error) {
	return s.repo.GetVideoByID(ctx, id)
}

func (s *VideoService) GetVideoByTitle(ctx context.Context, title string) (models.Video, error) {
	return s.repo.GetVideoByTitle(ctx, title)
}
