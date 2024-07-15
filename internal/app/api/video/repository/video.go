package repository

import (
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"gorm.io/gorm"
)

type VideoRepositoryClient struct {
	db *gorm.DB
}

type IVideoRepositoryClient interface {
	Commit() error
	Rollback() error
	GetVideoByClassID(classID int64) ([]entity.Video, error)
	GetVideoByID(videoID int64) (entity.Video, error)
}

func (r *VideoRepositoryClient) Commit() error {
	if err := r.db.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *VideoRepositoryClient) Rollback() error {
	if err := r.db.Rollback().Error; err != nil {
		return err
	}
	return nil

}

func (r *VideoRepositoryClient) GetVideoByClassID(classID int64) ([]entity.Video, error) {
	var videos []entity.Video
	if err := r.db.Where("class_id = ?", classID).Find(&videos).Error; err != nil {
		return []entity.Video{}, err
	}
	return videos, nil
}

func (r *VideoRepositoryClient) GetVideoByID(videoID int64) (entity.Video, error) {
	var video entity.Video
	if err := r.db.Where("id = ?", videoID).First(&video).Error; err != nil {
		return entity.Video{}, err
	}
	return video, nil
}
