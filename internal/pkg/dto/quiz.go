package dto

type AttemptQuizRequest struct {
	QuizID  int64         `json:"quizID"`
	Answers []AttemptQuiz `json:"answers"`
}

type AttemptQuiz struct {
	QuizQuestionID int64  `json:"quizQuestionID"`
	Answer         string `json:"answer"`
}

type AttemptQuizResponse struct {
	Result      float64 `json:"result"`
	RightAnswer int     `json:"rightAnswer"`
	WrongAnswer int     `json:"wrongAnswer"`
}
