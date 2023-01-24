package service

import "github.com/ResulShamuhammedov/fiber-server/pkg/repository"

type VideoService struct {
	repo repository.Video
}

func NewVideoService(repo *repository.Video) *VideoService {
	return &VideoService{repo: repo}
}
