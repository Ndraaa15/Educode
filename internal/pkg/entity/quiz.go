package entity

import (
	"time"
)

type Quiz struct {
	ID          int64          `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:text;not null" json:"description"`
	ClassID     int64          `gorm:"type:int;not null;foreignKey:ID;references:classes" json:"-"`
	Questions   []QuestionQuiz `gorm:"foreignKey:QuizID;references:ID" json:"questions,omitempty"`
	UpdateAt    time.Time      `gorm:"type:timestamp;not null;autoUpdateTime" json:"-"`
	CreateAt    time.Time      `gorm:"type:timestamp;not null;autoCreateTime" json:"-"`
}

type QuestionQuiz struct {
	ID       int64     `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	QuizID   int64     `gorm:"type:int;not null;foreignKey:ID;references:quizzes;primaryKey" json:"-"`
	Question string    `gorm:"type:text;not null" json:"question"`
	OptionA  string    `gorm:"type:text;not null" json:"optionA"`
	OptionB  string    `gorm:"type:text;not null" json:"optionB"`
	OptionC  string    `gorm:"type:text;not null" json:"optionC"`
	OptionD  string    `gorm:"type:text;not null" json:"optionD"`
	OptionE  string    `gorm:"type:text;not null" json:"optionE,omitempty"`
	Answer   string    `gorm:"type:text;not null" json:"-"`
	UpdateAt time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"-"`
	CreateAt time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"-"`
}
