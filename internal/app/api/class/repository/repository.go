package repository

import "gorm.io/gorm"

type ClassRepository struct {
	db *gorm.DB
}

type IClassRepository interface {
	NewClient(tx bool) IClassRepositoryClient
}

func NewClassRepository(db *gorm.DB) IClassRepository {
	return &ClassRepository{db: db}
}

func (r *ClassRepository) NewClient(tx bool) IClassRepositoryClient {
	if tx {
		return &ClassRepositoryClient{db: r.db.Begin()}
	}
	return &ClassRepositoryClient{db: r.db}
}
