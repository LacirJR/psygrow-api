package dto

import (
	"github.com/google/uuid"
)

// AnamneseFieldOptionRequest represents the request to create a new anamnese field option
type AnamneseFieldOptionRequest struct {
	OptionValue     string `json:"option_value" binding:"required"`
	OptionOrder     int    `json:"option_order" binding:"required"`
	AnamneseFieldID string `json:"anamnese_field_id" binding:"required,uuid"`
}

// AnamneseFieldOptionResponse represents the response for an anamnese field option
type AnamneseFieldOptionResponse struct {
	ID              string `json:"id"`
	OptionValue     string `json:"option_value"`
	OptionOrder     int    `json:"option_order"`
	AnamneseFieldID string `json:"anamnese_field_id"`
}

// NewAnamneseFieldOptionResponse creates a new AnamneseFieldOptionResponse from the given parameters
func NewAnamneseFieldOptionResponse(
	id uuid.UUID,
	optionValue string,
	optionOrder int,
	anamneseFieldID uuid.UUID,
) AnamneseFieldOptionResponse {
	return AnamneseFieldOptionResponse{
		ID:              id.String(),
		OptionValue:     optionValue,
		OptionOrder:     optionOrder,
		AnamneseFieldID: anamneseFieldID.String(),
	}
}

// AnamneseFieldOptionBulkRequest represents a request to create multiple anamnese field options at once
type AnamneseFieldOptionBulkRequest struct {
	AnamneseFieldID string                    `json:"anamnese_field_id" binding:"required,uuid"`
	Options         []AnamneseFieldOptionItem `json:"options" binding:"required,dive"`
}

// AnamneseFieldOptionItem represents a single option item in a bulk request
type AnamneseFieldOptionItem struct {
	OptionValue string `json:"option_value" binding:"required"`
	OptionOrder int    `json:"option_order" binding:"required"`
}
