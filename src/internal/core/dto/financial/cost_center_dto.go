package dto

import (
	"github.com/google/uuid"
)

// CostCenterRequest represents the request to create a new cost center
type CostCenterRequest struct {
	Name         string `json:"name" binding:"required"`
	RepasseModel string `json:"repasse_model" binding:"required,oneof=clinic_pays professional_pays"`
	RepasseType  string `json:"repasse_type" binding:"required,oneof=percent fixed"`
	RepasseValue int64  `json:"repasse_value" binding:"required"`
	Active       bool   `json:"active"`
}

// CostCenterUpdateRequest represents the request to update a cost center
type CostCenterUpdateRequest struct {
	Name         *string `json:"name,omitempty"`
	RepasseModel *string `json:"repasse_model,omitempty" binding:"omitempty,oneof=clinic_pays professional_pays"`
	RepasseType  *string `json:"repasse_type,omitempty" binding:"omitempty,oneof=percent fixed"`
	RepasseValue *int64  `json:"repasse_value,omitempty"`
	Active       *bool   `json:"active,omitempty"`
}

// CostCenterResponse represents the response for a cost center
type CostCenterResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	RepasseModel string `json:"repasse_model"`
	RepasseType  string `json:"repasse_type"`
	RepasseValue int64  `json:"repasse_value"`
	Active       bool   `json:"active"`
}

// NewCostCenterResponse creates a new CostCenterResponse from the given parameters
func NewCostCenterResponse(
	id uuid.UUID,
	userID uuid.UUID,
	name string,
	repasseModel string,
	repasseType string,
	repasseValue int64,
	active bool,
) CostCenterResponse {
	return CostCenterResponse{
		ID:           id.String(),
		UserID:       userID.String(),
		Name:         name,
		RepasseModel: repasseModel,
		RepasseType:  repasseType,
		RepasseValue: repasseValue,
		Active:       active,
	}
}
