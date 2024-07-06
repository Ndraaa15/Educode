package repository

import (
	"context"

	"github.com/Ndraaa15/IQuest/internal/pkg/entity"
	"gorm.io/gorm"
)

type CourseRepositoryClient struct {
	db *gorm.DB
}

type ICourseRepositoryClient interface {
	Commit() error
	Rollback() error
	CreateCourse(ctx context.Context, course entity.Course) error
	GetCourseByClassID(ctx context.Context, classID int64) ([]entity.Course, error)
	GetCourseByID(id int64) (entity.Course, error)
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

func (r *CourseRepositoryClient) GetCourseByID(id int64) (entity.Course, error) {
	var course entity.Course
	if err := r.db.Model(&entity.Course{}).Where("id = ?", id).First(&course).Error; err != nil {
		return course, err
	}
	return course, nil
}

func (r *CourseRepositoryClient) GetCourseByClassID(ctx context.Context, classID int64) ([]entity.Course, error) {
	var courses []entity.Course
	if err := r.db.WithContext(ctx).Model(&entity.Course{}).Where("class_id = ?", classID).Find(&courses).Error; err != nil {
		return courses, err
	}
	return courses, nil
}
