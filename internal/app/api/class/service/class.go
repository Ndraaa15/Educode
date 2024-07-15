package service

import (
	"context"

	"github.com/Ndraaa15/Educode/internal/app/api/class/repository"
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
)

type IClassService interface {
	GetClassByID(ctx context.Context, id int64) (entity.Class, error)
	UpdateClass(ctx context.Context, class entity.Class) error
}

type ClassService struct {
	class repository.IClassRepository
}

func NewClassService(class repository.IClassRepository) IClassService {
	return &ClassService{class: class}
}

func (s *ClassService) GetClassByID(ctx context.Context, id int64) (entity.Class, error) {
	db := s.class.NewClient(false)

	class, err := db.GetClassByID(id)
	if err != nil {
		return class, err
	}

	return class, nil
}

func (s *ClassService) UpdateClass(ctx context.Context, class entity.Class) error {
	db := s.class.NewClient(false)

	err := db.UpdateClass(ctx, class)
	if err != nil {
		return err
	}

	return nil
}
