package repository

import (
	"context"

	"github.com/Ndraaa15/IQuest/internal/pkg/entity"
	"gorm.io/gorm"
)

type QuizRepositoryClient struct {
	db *gorm.DB
}

type IQuizRepositoryClient interface {
	Commit() error
	Rollback() error
	GetQuizByClassID(ctx context.Context, classID int64) ([]entity.Quiz, error)
	GetQuizByID(ctx context.Context, quizID int64) (entity.Quiz, error)
	CreateResultQuiz(ctx context.Context, result entity.ResultQuiz) error
	GetQuizResultByQuizID(ctx context.Context, quizID int64) ([]entity.ResultQuiz, error)
}

func (r *QuizRepositoryClient) Commit() error {
	if err := r.db.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r *QuizRepositoryClient) Rollback() error {
	if err := r.db.Rollback().Error; err != nil {
		return err
	}
	return nil

}

func (r *QuizRepositoryClient) GetQuizByClassID(ctx context.Context, classID int64) ([]entity.Quiz, error) {
	var quizzes []entity.Quiz
	if err := r.db.WithContext(ctx).Where("class_id = ?", classID).Find(&quizzes).Error; err != nil {
		return []entity.Quiz{}, err
	}
	return quizzes, nil
}

func (r *QuizRepositoryClient) GetQuizByID(ctx context.Context, quizID int64) (entity.Quiz, error) {
	var quiz entity.Quiz
	if err := r.db.WithContext(ctx).Where("id = ?", quizID).Preload("Questions").First(&quiz).Error; err != nil {
		return entity.Quiz{}, err
	}
	return quiz, nil
}

func (r *QuizRepositoryClient) CreateResultQuiz(ctx context.Context, result entity.ResultQuiz) error {
	if err := r.db.WithContext(ctx).Model(&entity.ResultQuiz{}).Create(&result).Error; err != nil {
		return err
	}
	return nil
}

func (r *QuizRepositoryClient) GetQuizResultByQuizID(ctx context.Context, quizID int64) ([]entity.ResultQuiz, error) {
	var results []entity.ResultQuiz
	if err := r.db.WithContext(ctx).Where("quiz_id = ?", quizID).Preload("User").Find(&results).Error; err != nil {
		return []entity.ResultQuiz{}, err
	}
	return results, nil
}
