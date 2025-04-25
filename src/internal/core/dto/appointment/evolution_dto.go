package dto

import (
	"github.com/google/uuid"
	"time"
)

// EvolutionRequest represents the request to create a new evolution
type EvolutionRequest struct {
	SessionID string `json:"session_id" binding:"required,uuid"`
	Content   string `json:"content" binding:"required"`
}

// EvolutionResponse represents the response for an evolution
type EvolutionResponse struct {
	ID             string    `json:"id"`
	SessionID      string    `json:"session_id"`
	UserID         string    `json:"user_id"`
	ProfessionalID string    `json:"professional_id"`
	PatientID      string    `json:"patient_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

// NewEvolutionResponse creates a new EvolutionResponse from the given parameters
func NewEvolutionResponse(
	id uuid.UUID,
	sessionID uuid.UUID,
	userID uuid.UUID,
	professionalID uuid.UUID,
	patientID uuid.UUID,
	content string,
	createdAt time.Time,
) EvolutionResponse {
	return EvolutionResponse{
		ID:             id.String(),
		SessionID:      sessionID.String(),
		UserID:         userID.String(),
		ProfessionalID: professionalID.String(),
		PatientID:      patientID.String(),
		Content:        content,
		CreatedAt:      createdAt,
	}
}
