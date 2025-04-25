package dto

import (
	"github.com/google/uuid"
	"time"
)

// SessionRequest represents the request to create a new session
type SessionRequest struct {
	AppointmentID string    `json:"appointment_id" binding:"required,uuid"`
	StartTime     time.Time `json:"start_time" binding:"required"`
	EndTime       time.Time `json:"end_time" binding:"required"`
	WasAttended   bool      `json:"was_attended"`
}

// SessionResponse represents the response for a session
type SessionResponse struct {
	ID             string    `json:"id"`
	AppointmentID  string    `json:"appointment_id"`
	UserID         string    `json:"user_id"`
	PatientID      string    `json:"patient_id"`
	ProfessionalID string    `json:"professional_id"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	WasAttended    bool      `json:"was_attended"`
	CreatedAt      time.Time `json:"created_at"`
}

// NewSessionResponse creates a new SessionResponse from the given parameters
func NewSessionResponse(
	id uuid.UUID,
	appointmentID uuid.UUID,
	userID uuid.UUID,
	patientID uuid.UUID,
	professionalID uuid.UUID,
	startTime time.Time,
	endTime time.Time,
	wasAttended bool,
	createdAt time.Time,
) SessionResponse {
	return SessionResponse{
		ID:             id.String(),
		AppointmentID:  appointmentID.String(),
		UserID:         userID.String(),
		PatientID:      patientID.String(),
		ProfessionalID: professionalID.String(),
		StartTime:      startTime,
		EndTime:        endTime,
		WasAttended:    wasAttended,
		CreatedAt:      createdAt,
	}
}
