package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/Ndraaa15/IQuest/internal/app/api/auth/repository"
	"github.com/Ndraaa15/IQuest/internal/pkg/dto"
	"github.com/Ndraaa15/IQuest/internal/pkg/entity"
	"github.com/Ndraaa15/IQuest/pkg/errsx"
	"github.com/Ndraaa15/IQuest/pkg/utils/bcrypt"
	"github.com/Ndraaa15/IQuest/pkg/utils/jwt"
	"github.com/google/uuid"
)

type IAuthService interface {
	SignUp(ctx context.Context, req dto.SignUpRequest) error
	SignIn(ctx context.Context, req dto.SignInRequest) (string, int64, error)
}

type AuthService struct {
	auth repository.IAuthRepository
}

func NewAuthService(auth repository.IAuthRepository) IAuthService {
	return &AuthService{auth: auth}
}

func (s *AuthService) SignUp(ctx context.Context, req dto.SignUpRequest) error {
	db := s.auth.NewClient(true)

	var (
		errx error
	)

	defer func() {
		if errx != nil {
			if errTx := db.Rollback(); errTx != nil {
				errx = errTx
				return
			}
		}
	}()

	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		errx = err
		return err
	}

	role, err := ParseRole(req.Role)
	if err != nil {
		errx = err
		return errx
	}

	user := entity.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Password: hashedPassword,
		Role:     role,
	}

	if err := db.CreateUser(ctx, user); err != nil {
		errx = err
		return errx
	}

	if role == entity.RoleTeacher {
		if req.Class == "" {
			errx = errsx.NewCustomError(http.StatusBadRequest, "Service.Auth.SignUp", "Field class required", errors.New("INVALID_CLASS"))
			return errx
		}

		class := entity.Class{
			TeacherID: user.ID,
			Name:      req.Class,
			Goal:      "https://res.cloudinary.com/dcvnwpyd9/image/upload/v1720242761/iquest/jsal4xbga8q8cjhc74w3.pdf",
		}

		if err := db.CreateClass(ctx, class); err != nil {
			errx = err
			return errx
		}
	}

	db.Commit()

	return nil
}

func (s *AuthService) SignIn(ctx context.Context, req dto.SignInRequest) (string, int64, error) {
	db := s.auth.NewClient(false)

	user, err := db.GetUserByName(ctx, req.Name)
	if err != nil {
		return "", 0, err
	}

	if err := bcrypt.ComparePassword(user.Password, req.Password); err != nil {
		return "", 0, errsx.NewCustomError(http.StatusBadRequest, "Service.Auth.SignIn", "Invalid password", errors.New("INVALID_PASSWORD"))
	}

	var class entity.Class
	if user.Role == entity.RoleStudent {
		if req.Class == "" {
			return "", 0, errsx.NewCustomError(http.StatusBadRequest, "Service.Auth.SignIn", "Field class required", errors.New("INVALID_CLASS"))
		}

		class, err = db.GetClassByName(ctx, req.Class)
		if err != nil {
			return "", 0, err
		}
	} else if user.Role == entity.RoleTeacher {
		class, err = db.GetClassByUserID(ctx, user.ID.String())
		if err != nil {
			return "", 0, err
		}

	}

	token, err := jwt.EncodeToken(&user)
	if err != nil {
		return "", 0, err
	}

	return token, class.ID, nil
}
