package dto

import (
	"github.com/google/uuid"
	"time"
)

// AppointmentRequest represents the request to create a new appointment
type AppointmentRequest struct {
	PatientID          string    `json:"patient_id" binding:"required,uuid"`
	ProfessionalID     string    `json:"professional_id" binding:"required,uuid"`
	CostCenterID       string    `json:"cost_center_id" binding:"required,uuid"`
	ServiceTitle       string    `json:"service_title" binding:"required"`
	StartTime          time.Time `json:"start_time" binding:"required"`
	EndTime            time.Time `json:"end_time" binding:"required"`
	Notes              string    `json:"notes"`
	CustomRepasseType  *string   `json:"custom_repasse_type,omitempty" binding:"omitempty,oneof=percent fixed"`
	CustomRepasseValue *int64    `json:"custom_repasse_value,omitempty"`
}

// AppointmentUpdateRequest represents the request to update an appointment
type AppointmentUpdateRequest struct {
	PatientID          *string    `json:"patient_id,omitempty" binding:"omitempty,uuid"`
	ProfessionalID     *string    `json:"professional_id,omitempty" binding:"omitempty,uuid"`
	CostCenterID       *string    `json:"cost_center_id,omitempty" binding:"omitempty,uuid"`
	ServiceTitle       *string    `json:"service_title,omitempty"`
	StartTime          *time.Time `json:"start_time,omitempty"`
	EndTime            *time.Time `json:"end_time,omitempty"`
	Status             *string    `json:"status,omitempty" binding:"omitempty,oneof=scheduled done canceled no_show"`
	Notes              *string    `json:"notes,omitempty"`
	CustomRepasseType  *string    `json:"custom_repasse_type,omitempty" binding:"omitempty,oneof=percent fixed"`
	CustomRepasseValue *int64     `json:"custom_repasse_value,omitempty"`
}

// AppointmentResponse represents the response for an appointment
type AppointmentResponse struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	PatientID          string    `json:"patient_id"`
	ProfessionalID     string    `json:"professional_id"`
	CostCenterID       string    `json:"cost_center_id"`
	ServiceTitle       string    `json:"service_title"`
	StartTime          time.Time `json:"start_time"`
	EndTime            time.Time `json:"end_time"`
	Status             string    `json:"status"`
	Notes              string    `json:"notes,omitempty"`
	CustomRepasseType  *string   `json:"custom_repasse_type,omitempty"`
	CustomRepasseValue *float64  `json:"custom_repasse_value,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// NewAppointmentResponse creates a new AppointmentResponse from the given parameters
func NewAppointmentResponse(
	id uuid.UUID,
	userID uuid.UUID,
	patientID uuid.UUID,
	professionalID uuid.UUID,
	costCenterID uuid.UUID,
	serviceTitle string,
	startTime time.Time,
	endTime time.Time,
	status string,
	notes string,
	customRepasseType *string,
	customRepasseValue *float64,
	createdAt time.Time,
	updatedAt time.Time,
) AppointmentResponse {
	return AppointmentResponse{
		ID:                 id.String(),
		UserID:             userID.String(),
		PatientID:          patientID.String(),
		ProfessionalID:     professionalID.String(),
		CostCenterID:       costCenterID.String(),
		ServiceTitle:       serviceTitle,
		StartTime:          startTime,
		EndTime:            endTime,
		Status:             status,
		Notes:              notes,
		CustomRepasseType:  customRepasseType,
		CustomRepasseValue: customRepasseValue,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}
