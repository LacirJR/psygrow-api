package dto

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/google/uuid"
	"time"
)

// PatientFamilyRequest represents the request body for creating or updating a patient family member
type PatientFamilyRequest struct {
	PatientID    uuid.UUID  `json:"patient_id" binding:"required"`
	Relationship string     `json:"relationship" binding:"required"` // Use constants from model.Relationship*
	Name         string     `json:"name" binding:"required,min=2,max=100"`
	BirthDate    *time.Time `json:"birth_date"`
	Schooling    *string    `json:"schooling"`
	Occupation   *string    `json:"occupation"`
}

// PatientFamilyResponse represents the response body for a patient family member
type PatientFamilyResponse struct {
	ID           uuid.UUID  `json:"id"`
	PatientID    uuid.UUID  `json:"patient_id"`
	Relationship string     `json:"relationship"`
	Name         string     `json:"name"`
	BirthDate    *time.Time `json:"birth_date"`
	Schooling    *string    `json:"schooling"`
	Occupation   *string    `json:"occupation"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// NewPatientFamilyResponse creates a new PatientFamilyResponse from a PatientFamily model
func NewPatientFamilyResponse(patientFamily model.PatientFamily) PatientFamilyResponse {
	return PatientFamilyResponse{
		ID:           patientFamily.ID,
		PatientID:    patientFamily.PatientID,
		Relationship: patientFamily.Relationship,
		Name:         patientFamily.Name,
		BirthDate:    patientFamily.BirthDate,
		Schooling:    patientFamily.Schooling,
		Occupation:   patientFamily.Occupation,
		CreatedAt:    patientFamily.CreatedAt,
		UpdatedAt:    patientFamily.UpdatedAt,
	}
}
