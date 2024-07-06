package service

import (
	"errors"

	"github.com/Ndraaa15/IQuest/internal/pkg/entity"
)

func ParseRole(role string) (entity.Role, error) {
	switch role {
	case "teacher":
		return entity.RoleTeacher, nil
	case "student":
		return entity.RoleStudent, nil
	default:
		return entity.RoleUnknown, errors.New("INVALID_ROLE")
	}
}
