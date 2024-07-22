package entity

import (
	"time"
)

type Course struct {
	ClassID     int64     `gorm:"type:int;not null;foreignKey:ID;references:classes" json:"-"`
	Course      string    `gorm:"type:text;not null" json:"course"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	UpdateAt    time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"-"`
	CreateAt    time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"-"`
}
