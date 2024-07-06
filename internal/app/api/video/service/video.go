package service

import (
	"context"

	"github.com/Ndraaa15/IQuest/internal/app/api/video/repository"
	"github.com/Ndraaa15/IQuest/internal/pkg/entity"
)

type IVideoService interface {
	GetVideoByClassID(ctx context.Context, classID int64) ([]entity.Video, error)
	GetVideoByID(ctx context.Context, videoID int64) (entity.Video, error)
}

type VideoService struct {
	video repository.IVideoRepository
}

func NewVideoService(video repository.IVideoRepository) IVideoService {
	return &VideoService{video: video}
}

func (s *VideoService) GetVideoByClassID(ctx context.Context, classID int64) ([]entity.Video, error) {
	db := s.video.NewClient(false)

	videos, err := db.GetVideoByClassID(classID)
	if err != nil {
		return videos, err
	}

	return videos, nil
}

func (s *VideoService) GetVideoByID(ctx context.Context, videoID int64) (entity.Video, error) {
	db := s.video.NewClient(false)

	video, err := db.GetVideoByID(videoID)
	if err != nil {
		return video, err
	}

	return video, nil
}
