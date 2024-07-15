package repository

import (
	"context"

	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"gorm.io/gorm"
)

type AuthRepositoryClient struct {
	db *gorm.DB
}

type IAuthRepositoryClient interface {
	Commit() error
	Rollback() error
	CreateUser(ctx context.Context, user entity.User) error
	CreateClass(ctx context.Context, class entity.Class) error
	GetUserByName(ctx context.Context, name string) (entity.User, error)
	GetClassByName(ctx context.Context, name string) (entity.Class, error)
	GetClassByUserID(ctx context.Context, userID string) (entity.Class, error)
}

func (r *AuthRepositoryClient) Commit() error {
	if err := r.db.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryClient) Rollback() error {
	if err := r.db.Rollback().Error; err != nil {
		return err
	}
	return nil

}

func (r *AuthRepositoryClient) CreateUser(ctx context.Context, user entity.User) error {
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryClient) CreateClass(ctx context.Context, class entity.Class) error {
	if err := r.db.WithContext(ctx).Model(&entity.Class{}).Create(&class).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryClient) GetUserByName(ctx context.Context, name string) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("name = ?", name).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthRepositoryClient) GetClassByName(ctx context.Context, name string) (entity.Class, error) {
	var class entity.Class
	if err := r.db.WithContext(ctx).Model(&entity.Class{}).Where("name = ?", name).First(&class).Error; err != nil {
		return class, err
	}
	return class, nil
}

func (r *AuthRepositoryClient) GetClassByUserID(ctx context.Context, userID string) (entity.Class, error) {
	var class entity.Class
	if err := r.db.WithContext(ctx).Model(&entity.Class{}).Where("teacher_id = ?", userID).First(&class).Error; err != nil {
		return class, err
	}
	return class, nil
}
