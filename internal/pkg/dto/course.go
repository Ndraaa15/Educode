package dto

type CreateCourseRequest struct {
	ClassID     int64  `json:"classID" binding:"required"`
	Course      string `json:"course" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
