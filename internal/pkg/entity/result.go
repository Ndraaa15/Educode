package entity

import (
	"time"

	"github.com/google/uuid"
)

type ResultQuiz struct {
	ID          int64     `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	QuizID      int64     `json:"quizID" gorm:"type:int;not null;foreignKey:ID;references:quizzes"`
	UserID      uuid.UUID `json:"userID" gorm:"type:varchar(36);not null;foreignKey:ID;references:users"`
	User        User      `json:"user,omitempty"`
	Result      float64   `json:"result" gorm:"type:numeric;not null"`
	RightAnswer int       `json:"rightAnswer" gorm:"type:int;not null"`
	WrongAnswer int       `json:"wrongAnswer" gorm:"type:int;not null"`
	UpdateAt    time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"-"`
	CreateAt    time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"-"`
}
