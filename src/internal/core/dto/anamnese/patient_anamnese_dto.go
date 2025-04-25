package dto

import (
	"github.com/google/uuid"
	"time"
)

// PatientAnamneseRequest represents the request to create a new patient anamnese
type PatientAnamneseRequest struct {
	PatientID  string                        `json:"patient_id" binding:"required,uuid"`
	AnamneseID string                        `json:"anamnese_id" binding:"required,uuid"`
	Fields     []PatientAnamneseFieldRequest `json:"fields" binding:"required,dive"`
}

// PatientAnamneseResponse represents the response for a patient anamnese
type PatientAnamneseResponse struct {
	ID         string                         `json:"id"`
	PatientID  string                         `json:"patient_id"`
	AnamneseID string                         `json:"anamnese_id"`
	UserID     string                         `json:"user_id"`
	AnsweredAt time.Time                      `json:"answered_at"`
	Fields     []PatientAnamneseFieldResponse `json:"fields,omitempty"`
}

// NewPatientAnamneseResponse creates a new PatientAnamneseResponse from the given parameters
func NewPatientAnamneseResponse(
	id uuid.UUID,
	patientID uuid.UUID,
	anamneseID uuid.UUID,
	userID uuid.UUID,
	answeredAt time.Time,
) PatientAnamneseResponse {
	return PatientAnamneseResponse{
		ID:         id.String(),
		PatientID:  patientID.String(),
		AnamneseID: anamneseID.String(),
		UserID:     userID.String(),
		AnsweredAt: answeredAt,
		Fields:     []PatientAnamneseFieldResponse{},
	}
}

// PatientAnamneseFieldRequest represents the request to create a new patient anamnese field
type PatientAnamneseFieldRequest struct {
	FieldID string `json:"field_id" binding:"required,uuid"`
	Value   string `json:"value"`
}

// PatientAnamneseFieldResponse represents the response for a patient anamnese field
type PatientAnamneseFieldResponse struct {
	ID                string `json:"id"`
	PatientAnamneseID string `json:"patient_anamnese_id"`
	FieldID           string `json:"field_id"`
	Value             string `json:"value"`
	// Include field details for convenience
	FieldType  string `json:"field_type,omitempty"`
	FieldTitle string `json:"field_title,omitempty"`
}

// NewPatientAnamneseFieldResponse creates a new PatientAnamneseFieldResponse from the given parameters
func NewPatientAnamneseFieldResponse(
	id uuid.UUID,
	patientAnamneseID uuid.UUID,
	fieldID uuid.UUID,
	value string,
	fieldType string,
	fieldTitle string,
) PatientAnamneseFieldResponse {
	return PatientAnamneseFieldResponse{
		ID:                id.String(),
		PatientAnamneseID: patientAnamneseID.String(),
		FieldID:           fieldID.String(),
		Value:             value,
		FieldType:         fieldType,
		FieldTitle:        fieldTitle,
	}
}
