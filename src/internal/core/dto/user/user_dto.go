package dto

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	//Role     string    `json:"role" binding:"default: 'professional'"`
	Phone *string `json:"phone"`
}
