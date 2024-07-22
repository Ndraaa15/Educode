package entity

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID        int64     `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	TeacherID uuid.UUID `gorm:"not null;constraint:OnDelete:CASCADE;foreignKey:ID;references:users" json:"teacherID"`
	Goal      string    `gorm:"type:text;not null" json:"goal"`
	Courses   []Course  `gorm:"foreignKey:ClassID;references:ID" json:"courses"`
	Quizzes   []Quiz    `gorm:"foreignKey:ClassID;references:ID" json:"quizzes"`
	Videos    []Video   `gorm:"foreignKey:ClassID;references:ID" json:"videos"`
	UpdateAt  time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"-"`
	CreateAt  time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"-"`
}
