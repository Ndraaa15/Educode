package dto

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Class    string `json:"class"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignInRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Class    string `json:"class"`
}
