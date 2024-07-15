package service

import (
	"context"
	"fmt"
	"os"

	"github.com/Ndraaa15/Educode/internal/app/api/quiz/repository"
	"github.com/Ndraaa15/Educode/internal/pkg/dto"
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"github.com/Ndraaa15/Educode/pkg/cloudinary"
	"github.com/Ndraaa15/Educode/pkg/pdf"
	"github.com/google/uuid"
)

type IQuizService interface {
	GetQuizByClassID(ctx context.Context, classID int64) ([]entity.Quiz, error)
	GetQuizByID(ctx context.Context, quizID int64) (entity.Quiz, error)
	AttemptQuiz(ctx context.Context, req dto.AttemptQuizRequest, userID uuid.UUID) (dto.AttemptQuizResponse, error)
	GetQuizResult(ctx context.Context, quizID int64) (string, error)
}

type QuizService struct {
	course repository.IQuizRepository
}

func NewQuizService(course repository.IQuizRepository) IQuizService {
	return &QuizService{course: course}
}

func (s *QuizService) GetQuizByClassID(ctx context.Context, classID int64) ([]entity.Quiz, error) {
	db := s.course.NewClient(false)

	quizzes, err := db.GetQuizByClassID(ctx, classID)
	if err != nil {
		return quizzes, err
	}

	return quizzes, nil
}

func (s *QuizService) GetQuizByID(ctx context.Context, quizID int64) (entity.Quiz, error) {
	db := s.course.NewClient(false)

	quiz, err := db.GetQuizByID(ctx, quizID)
	if err != nil {
		return quiz, err
	}

	return quiz, nil
}

func (s *QuizService) AttemptQuiz(ctx context.Context, req dto.AttemptQuizRequest, userID uuid.UUID) (dto.AttemptQuizResponse, error) {
	db := s.course.NewClient(true)

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

	quiz, err := db.GetQuizByID(ctx, req.QuizID)
	if err != nil {
		errx = err
		return dto.AttemptQuizResponse{}, err
	}

	questionsWithAnswer := make(map[int64]string, 0)
	for _, question := range quiz.Questions {
		questionsWithAnswer[question.ID] = question.Answer
	}

	var totalRight int
	for _, attempt := range req.Answers {
		if questionsWithAnswer[attempt.QuizQuestionID] == attempt.Answer {
			totalRight++
		}
	}

	score := float64(totalRight) / float64(len(quiz.Questions)) * 100

	attempt := entity.ResultQuiz{
		UserID:      userID,
		QuizID:      req.QuizID,
		Result:      score,
		RightAnswer: totalRight,
		WrongAnswer: len(quiz.Questions) - totalRight,
	}

	if err := db.CreateResultQuiz(ctx, attempt); err != nil {
		errx = err
		return dto.AttemptQuizResponse{}, err
	}

	if err := db.Commit(); err != nil {
		errx = err
		return dto.AttemptQuizResponse{}, err
	}

	return dto.AttemptQuizResponse{
		Result:      score,
		RightAnswer: totalRight,
		WrongAnswer: len(quiz.Questions) - totalRight,
	}, nil
}

func (s *QuizService) GetQuizResult(ctx context.Context, quizID int64) (string, error) {
	db := s.course.NewClient(false)

	result, err := db.GetQuizResultByQuizID(ctx, quizID)
	if err != nil {
		return "", err
	}

	data := PopulatePDFData(result)

	pdfPath, err := pdf.GenerateResultQuiz(data)
	if err != nil {
		return "", err
	}

	cld, err := cloudinary.NewCloudinary()
	if err != nil {
		return "", err
	}

	file, err := os.Open(pdfPath)
	if err != nil {
		return "", err
	}

	invoiceURI, err := cld.UploadFile(ctx, file)
	if err != nil {
		return "", err
	}

	file.Close()
	err = os.Remove(pdfPath)
	if err != nil {
		return "", err
	}

	return invoiceURI, nil
}

func PopulatePDFData(results []entity.ResultQuiz) map[string]interface{} {
	var resultQuiz [][]string

	for _, result := range results {
		resultQuiz = append(resultQuiz, []string{
			result.User.Name,
			fmt.Sprintf("%.2f", result.Result),
		})
	}

	data := map[string]interface{}{
		"resultQuiz": resultQuiz,
	}

	return data
}
