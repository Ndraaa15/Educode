package repository

import "gorm.io/gorm"

type QuizRepository struct {
	db *gorm.DB
}

type IQuizRepository interface {
	NewClient(tx bool) IQuizRepositoryClient
}

func NewQuizRepository(db *gorm.DB) IQuizRepository {
	return &QuizRepository{db: db}
}

func (r *QuizRepository) NewClient(tx bool) IQuizRepositoryClient {
	if tx {
		return &QuizRepositoryClient{db: r.db.Begin()}
	}
	return &QuizRepositoryClient{db: r.db}
}
