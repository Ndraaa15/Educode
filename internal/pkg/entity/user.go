package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	Password string    `gorm:"type:varchar(255);not null" json:"password"`
	Class    Class     `gorm:"foreignKey:TeacherID;references:ID" json:"class,omitempty"`
	Role     Role      `gorm:"type:int;not null" json:"role"`
	UpdateAt time.Time `gorm:"type:timestamp;autoUpdateTime" json:"-"`
	CreateAt time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"-"`
}

type JWTClaims struct {
	ID   uuid.UUID
	Role string
	jwt.RegisteredClaims
}

type Role int64

const (
	RoleUnknown Role = 0
	RoleStudent Role = 1
	RoleTeacher Role = 2
)

var (
	StatusName = map[Role]string{
		RoleStudent: "student",
		RoleTeacher: "teacher",
	}
)

func (s Role) String() string {
	return StatusName[s]
}

func (s Role) Value() int {
	return int(s)
}
