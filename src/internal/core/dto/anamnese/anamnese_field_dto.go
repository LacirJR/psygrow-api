package dto

import (
	"github.com/google/uuid"
)

// AnamneseFieldRequest represents the request to create a new anamnese field
type AnamneseFieldRequest struct {
	FieldNumber   int                       `json:"field_number" binding:"required"`
	FieldType     string                    `json:"field_type" binding:"required,oneof=date datetime text number checkbox select multiselect"`
	FieldTitle    string                    `json:"field_title" binding:"required"`
	FieldRequired bool                      `json:"field_required"`
	FieldActive   bool                      `json:"field_active"`
	AnamneseID    string                    `json:"anamnese_id" binding:"required,uuid"`
	Options       []AnamneseFieldOptionItem `json:"options,omitempty"`
}

// AnamneseFieldResponse represents the response for an anamnese field
type AnamneseFieldResponse struct {
	ID            string                        `json:"id"`
	FieldNumber   int                           `json:"field_number"`
	FieldType     string                        `json:"field_type"`
	FieldTitle    string                        `json:"field_title"`
	FieldRequired bool                          `json:"field_required"`
	FieldActive   bool                          `json:"field_active"`
	UserID        string                        `json:"user_id"`
	AnamneseID    string                        `json:"anamnese_id"`
	Options       []AnamneseFieldOptionResponse `json:"options,omitempty"`
}

// NewAnamneseFieldResponse creates a new AnamneseFieldResponse from the given parameters
func NewAnamneseFieldResponse(
	id uuid.UUID,
	fieldNumber int,
	fieldType string,
	fieldTitle string,
	fieldRequired bool,
	fieldActive bool,
	userID uuid.UUID,
	anamneseID uuid.UUID,
	options []AnamneseFieldOptionResponse,
) AnamneseFieldResponse {
	return AnamneseFieldResponse{
		ID:            id.String(),
		FieldNumber:   fieldNumber,
		FieldType:     fieldType,
		FieldTitle:    fieldTitle,
		FieldRequired: fieldRequired,
		FieldActive:   fieldActive,
		UserID:        userID.String(),
		AnamneseID:    anamneseID.String(),
		Options:       options,
	}
}
