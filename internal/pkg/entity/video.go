package entity

import (
	"time"
)

type Video struct {
	ID          int64     `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	ClassID     int64     `gorm:"type:int;not null;foreignKey:ID;references:classes" json:"-"`
	Image       string    `gorm:"type:text;not null" json:"image"`
	Video       string    `gorm:"type:text;not null" json:"video"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	UpdateAt    time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"-"`
	CreateAt    time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"-"`
}
