package dto

import "github.com/google/uuid"

type RegisterUserDto struct {
	ID           uuid.UUID `json:"id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required,email"`
	PasswordHash string    `json:"passwordHash" binding:"required"`
	Role         string    `json:"role" binding:"required; default: 'professional'"`
	Phone        *string   `json:"phone" binding:"phone"`
}
