package service

import (
	"context"
	"errors"

	"github.com/Ndraaa15/Educode/internal/app/api/course/repository"
	"github.com/Ndraaa15/Educode/internal/pkg/dto"
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"gorm.io/gorm"
)

type ICourseService interface {
	CreateCourse(ctx context.Context, req dto.CreateCourseRequest) error
	GetCourseByClassID(ctx context.Context, classID int64) (entity.Course, error)
}

type CourseService struct {
	course repository.ICourseRepository
}

func NewCourseService(course repository.ICourseRepository) ICourseService {
	return &CourseService{course: course}
}

func (s *CourseService) CreateCourse(ctx context.Context, req dto.CreateCourseRequest) error {
	db := s.course.NewClient(true)

	var (
		errx error
	)

	defer func() {
		if errx != nil {
			if errTx := db.Rollback(); errTx != nil {
				errx = errTx
				return
			}
		}
	}()

	course, err := db.GetCourseByClassID(ctx, req.ClassID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	parseCourseRequest(&req, &course)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.CreateCourse(ctx, course); err != nil {
			return err
		}
	} else {
		if err := db.UpdateCourse(ctx, req.ClassID, course); err != nil {
			return err
		}
	}

	if err := db.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *CourseService) GetCourseByClassID(ctx context.Context, classID int64) (entity.Course, error) {
	db := s.course.NewClient(false)

	courses, err := db.GetCourseByClassID(ctx, classID)
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func parseCourseRequest(req *dto.CreateCourseRequest, course *entity.Course) {
	if req.Course != "" {
		course.Course = req.Course
	}

	if req.Title != "" {
		course.Title = req.Title
	}

	if req.Description != "" {
		course.Description = req.Description
	}
}
