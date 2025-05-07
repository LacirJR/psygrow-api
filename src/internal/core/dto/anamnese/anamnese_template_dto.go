package dto

import (
	"github.com/google/uuid"
)

// AnamneseTemplateRequest represents the request to create a new anamnese template
type AnamneseTemplateRequest struct {
	Title string `json:"title" binding:"required"`
}

// AnamneseTemplateResponse represents the response for an anamnese template
type AnamneseTemplateResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	UserID string `json:"user_id"`
}

// NewAnamneseTemplateResponse creates a new AnamneseTemplateResponse from the given parameters
func NewAnamneseTemplateResponse(id uuid.UUID, title string, userID uuid.UUID) AnamneseTemplateResponse {
	return AnamneseTemplateResponse{
		ID:     id.String(),
		Title:  title,
		UserID: userID.String(),
	}
}
