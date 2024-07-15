package service

import (
	"context"

	"github.com/Ndraaa15/Educode/internal/app/api/course/repository"
	"github.com/Ndraaa15/Educode/internal/pkg/dto"
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
)

type ICourseService interface {
	CreateCourse(ctx context.Context, req dto.CreateCourseRequest) error
	GetCourseByID(ctx context.Context, id int64) (entity.Course, error)
	GetCourseByClassID(ctx context.Context, classID int64) ([]entity.Course, error)
}

type CourseService struct {
	course repository.ICourseRepository
}

func NewCourseService(course repository.ICourseRepository) ICourseService {
	return &CourseService{course: course}
}

func (s *CourseService) CreateCourse(ctx context.Context, req dto.CreateCourseRequest) error {
	db := s.course.NewClient(false)

	course := entity.Course{
		Title:       req.Title,
		Description: req.Description,
		Course:      "https://res.cloudinary.com/dcvnwpyd9/image/upload/v1720244441/iquest/i3l195tfithqpdgolfqd.pdf",
		ClassID:     req.ClassID,
	}

	if err := db.CreateCourse(ctx, course); err != nil {
		return err
	}

	return nil
}

func (s *CourseService) GetCourseByID(ctx context.Context, id int64) (entity.Course, error) {
	db := s.course.NewClient(false)

	course, err := db.GetCourseByID(id)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (s *CourseService) GetCourseByClassID(ctx context.Context, classID int64) ([]entity.Course, error) {
	db := s.course.NewClient(false)

	courses, err := db.GetCourseByClassID(ctx, classID)
	if err != nil {
		return courses, err
	}

	return courses, nil
}
