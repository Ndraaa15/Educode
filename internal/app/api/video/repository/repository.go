package repository

import "gorm.io/gorm"

type VideoRepository struct {
	db *gorm.DB
}

type IVideoRepository interface {
	NewClient(tx bool) IVideoRepositoryClient
}

func NewVideoRepository(db *gorm.DB) IVideoRepository {
	return &VideoRepository{db: db}
}

func (r *VideoRepository) NewClient(tx bool) IVideoRepositoryClient {
	if tx {
		return &VideoRepositoryClient{db: r.db.Begin()}
	}
	return &VideoRepositoryClient{db: r.db}
}
