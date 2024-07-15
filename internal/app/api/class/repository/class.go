package repository

import (
	"context"

	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"gorm.io/gorm"
)

type ClassRepositoryClient struct {
	db *gorm.DB
}

type IClassRepositoryClient interface {
	Commit() error
	Rollback() error
	GetClassByID(id int64) (entity.Class, error)
	UpdateClass(ctx context.Context, class entity.Class) error
}

func (r *ClassRepositoryClient) Commit() error {
	if err := r.db.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *ClassRepositoryClient) Rollback() error {
	if err := r.db.Rollback().Error; err != nil {
		return err
	}
	return nil

}

func (r *ClassRepositoryClient) GetClassByID(id int64) (entity.Class, error) {
	var class entity.Class
	if err := r.db.Model(&entity.Class{}).Preload("Courses").Preload("Quizzes").Preload("Videos").Where("id = ?", id).First(&class).Error; err != nil {
		return class, err
	}
	return class, nil
}

func (r *ClassRepositoryClient) UpdateClass(ctx context.Context, class entity.Class) error {
	if err := r.db.WithContext(ctx).Model(&entity.Class{}).Where("id = ?", class.ID).Save(&class).Error; err != nil {
		return err
	}
	return nil
}
