package repository

import "gorm.io/gorm"

type CourseRepository struct {
	db *gorm.DB
}

type ICourseRepository interface {
	NewClient(tx bool) ICourseRepositoryClient
}

func NewCourseRepository(db *gorm.DB) ICourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) NewClient(tx bool) ICourseRepositoryClient {
	if tx {
		return &CourseRepositoryClient{db: r.db.Begin()}
	}
	return &CourseRepositoryClient{db: r.db}
}
