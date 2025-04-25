package dto

import (
	"github.com/google/uuid"
	"time"
)

// RepasseRequest represents the request to create a new repasse
type RepasseRequest struct {
	AppointmentID  string     `json:"appointment_id" binding:"required,uuid"`
	CostCenterID   string     `json:"cost_center_id" binding:"required,uuid"`
	Value          int64      `json:"value" binding:"required,gt=0"`
	ClinicReceives bool       `json:"clinic_receives"`
	Status         string     `json:"status" binding:"required,oneof=pending paid informational"` // TODO: Use constants from model.RepasseStatus*
	PaidAt         *time.Time `json:"paid_at,omitempty"`
	Notes          string     `json:"notes"`
}

// RepasseUpdateRequest represents the request to update a repasse
type RepasseUpdateRequest struct {
	Value          *float64   `json:"value,omitempty" binding:"omitempty,gt=0"`
	ClinicReceives *bool      `json:"clinic_receives,omitempty"`
	Status         *string    `json:"status,omitempty" binding:"omitempty,oneof=pending paid informational"`
	PaidAt         *time.Time `json:"paid_at,omitempty"`
	Notes          *string    `json:"notes,omitempty"`
}

// RepasseResponse represents the response for a repasse
type RepasseResponse struct {
	ID             string     `json:"id"`
	UserID         string     `json:"user_id"`
	AppointmentID  string     `json:"appointment_id"`
	CostCenterID   string     `json:"cost_center_id"`
	Value          float64    `json:"value"`
	ClinicReceives bool       `json:"clinic_receives"`
	Status         string     `json:"status"`
	PaidAt         *time.Time `json:"paid_at,omitempty"`
	Notes          string     `json:"notes,omitempty"`
}

// NewRepasseResponse creates a new RepasseResponse from the given parameters
func NewRepasseResponse(
	id uuid.UUID,
	userID uuid.UUID,
	appointmentID uuid.UUID,
	costCenterID uuid.UUID,
	value float64,
	clinicReceives bool,
	status string,
	paidAt *time.Time,
	notes string,
) RepasseResponse {
	return RepasseResponse{
		ID:             id.String(),
		UserID:         userID.String(),
		AppointmentID:  appointmentID.String(),
		CostCenterID:   costCenterID.String(),
		Value:          value,
		ClinicReceives: clinicReceives,
		Status:         status,
		PaidAt:         paidAt,
		Notes:          notes,
	}
}
