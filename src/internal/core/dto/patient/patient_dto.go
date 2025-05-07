package dto

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/google/uuid"
	"time"
)

// PatientRequest represents the request body for creating or updating a patient
type PatientRequest struct {
	CostCenterID          uuid.UUID  `json:"cost_center_id" binding:"required"`
	FullName              string     `json:"full_name" binding:"required,min=2,max=100"`
	SocialName            *string    `json:"social_name"`
	BirthDate             time.Time  `json:"birth_date" binding:"required"`
	Document              *string    `json:"document"`
	Phone                 *string    `json:"phone"`
	Email                 *string    `json:"email"`
	Gender                *string    `json:"gender"`
	Address               *string    `json:"address"`
	ResidesWith           *string    `json:"resides_with"`
	EmergencyContactName  *string    `json:"emergency_contact_name"`
	EmergencyContactPhone *string    `json:"emergency_contact_phone"`
	Observation           *string    `json:"observation"`
	DefaultRepasseType    *string    `json:"default_repasse_type" binding:"omitempty,oneof=percent fixed"`
	DefaultRepasseValue   *int64     `json:"default_repasse_value"` // Stored as cents or basis points (for percent)
	IsActive              bool       `json:"is_active"`
}

// PatientResponse represents the response body for a patient
type PatientResponse struct {
	ID                    uuid.UUID  `json:"id"`
	CostCenterID          uuid.UUID  `json:"cost_center_id"`
	CostCenterName        string     `json:"cost_center_name"`
	FullName              string     `json:"full_name"`
	SocialName            *string    `json:"social_name"`
	BirthDate             time.Time  `json:"birth_date"`
	Document              *string    `json:"document"`
	Phone                 *string    `json:"phone"`
	Email                 *string    `json:"email"`
	Gender                *string    `json:"gender"`
	Address               *string    `json:"address"`
	ResidesWith           *string    `json:"resides_with"`
	EmergencyContactName  *string    `json:"emergency_contact_name"`
	EmergencyContactPhone *string    `json:"emergency_contact_phone"`
	Observation           *string    `json:"observation"`
	DefaultRepasseType    *string    `json:"default_repasse_type"`
	DefaultRepasseValue   *int64     `json:"default_repasse_value"`
	IsActive              bool       `json:"is_active"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

// NewPatientResponse creates a new PatientResponse from a Patient model
func NewPatientResponse(patient model.Patient) PatientResponse {
	return PatientResponse{
		ID:                    patient.ID,
		CostCenterID:          patient.CostCenterID,
		CostCenterName:        patient.CostCenter.Name,
		FullName:              patient.FullName,
		SocialName:            patient.SocialName,
		BirthDate:             patient.BirthDate,
		Document:              patient.Document,
		Phone:                 patient.Phone,
		Email:                 patient.Email,
		Gender:                patient.Gender,
		Address:               patient.Address,
		ResidesWith:           patient.ResidesWith,
		EmergencyContactName:  patient.EmergencyContactName,
		EmergencyContactPhone: patient.EmergencyContactPhone,
		Observation:           patient.Observation,
		DefaultRepasseType:    patient.DefaultRepasseType,
		DefaultRepasseValue:   patient.DefaultRepasseValue,
		IsActive:              patient.IsActive,
		CreatedAt:             patient.CreatedAt,
		UpdatedAt:             patient.UpdatedAt,
	}
}