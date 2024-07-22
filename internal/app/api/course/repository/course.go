package repository

import (
	"context"

	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"gorm.io/gorm"
)

type CourseRepositoryClient struct {
	db *gorm.DB
}

type ICourseRepositoryClient interface {
	Commit() error
	Rollback() error
	CreateCourse(ctx context.Context, course entity.Course) error
	GetCourseByClassID(ctx context.Context, classID int64) (entity.Course, error)
	UpdateCourse(ctx context.Context, classID int64, course entity.Course) error
}

func (r *CourseRepositoryClient) Commit() error {
	if err := r.db.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *CourseRepositoryClient) Rollback() error {
	if err := r.db.Rollback().Error; err != nil {
		return err
	}
	return nil

}

func (r *CourseRepositoryClient) CreateCourse(ctx context.Context, course entity.Course) error {
	if err := r.db.WithContext(ctx).Model(&entity.Course{}).Create(&course).Error; err != nil {
		return err
	}
	return nil
}

func (r *CourseRepositoryClient) GetCourseByClassID(ctx context.Context, classID int64) (entity.Course, error) {
	var courses entity.Course
	if err := r.db.WithContext(ctx).Model(&entity.Course{}).Where("class_id = ?", classID).First(&courses).Error; err != nil {
		return entity.Course{}, err
	}
	return courses, nil
}

func (r *CourseRepositoryClient) UpdateCourse(ctx context.Context, classID int64, course entity.Course) error {
	if err := r.db.WithContext(ctx).Model(&entity.Course{}).Where("class_id = ?", classID).Updates(&course).Error; err != nil {
		return err
	}
	return nil
}
