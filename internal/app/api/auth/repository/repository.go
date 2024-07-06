package repository

import "gorm.io/gorm"

type AuthRepository struct {
	db *gorm.DB
}

type IAuthRepository interface {
	NewClient(tx bool) IAuthRepositoryClient
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) NewClient(tx bool) IAuthRepositoryClient {
	if tx {
		return &AuthRepositoryClient{db: r.db.Begin()}
	}
	return &AuthRepositoryClient{db: r.db}
}
