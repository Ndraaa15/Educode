package postgresql

import (
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Class{},
		&entity.Quiz{},
		&entity.QuestionQuiz{},
		&entity.Video{},
		&entity.Course{},
		&entity.ResultQuiz{},
	)
	if err != nil {
		return err
	}

	return nil
}

func Drop(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&entity.User{},
		&entity.Class{},
		&entity.Quiz{},
		&entity.QuestionQuiz{},
		&entity.Video{},
		&entity.Course{},
		&entity.ResultQuiz{},
	)
	if err != nil {
		return err
	}

	return nil
}
