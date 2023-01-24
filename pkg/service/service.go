package service

import "github.com/ResulShamuhammedov/fiber-server/pkg/repository"

type Service struct {
	Video
}

type Video interface {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Video: NewVideoService(&repo.Video),
	}
}
