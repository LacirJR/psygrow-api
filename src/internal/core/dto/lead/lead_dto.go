package dto

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/google/uuid"
	"time"
)

// LeadRequest represents the request body for creating or updating a lead
type LeadRequest struct {
	FullName         string     `json:"full_name" binding:"required,min=2,max=100"`
	Phone            *string    `json:"phone"`
	Email            *string    `json:"email"`
	BirthDate        *time.Time `json:"birth_date"`
	ContactDate      time.Time  `json:"contact_date" binding:"required"`
	Status           string     `json:"status" binding:"required,oneof=new in_analysis converted lost"` // TODO: Use constants from model.LeadStatus*
	WasAttended      bool       `json:"was_attended"`
	Notes            *string    `json:"notes"`
	Origin           *string    `json:"origin"`
	GdprBlockContact bool       `json:"gdpr_block_contact"`
}

// LeadResponse represents the response body for a lead
type LeadResponse struct {
	ID               uuid.UUID  `json:"id"`
	FullName         string     `json:"full_name"`
	Phone            *string    `json:"phone"`
	Email            *string    `json:"email"`
	BirthDate        *time.Time `json:"birth_date"`
	ContactDate      time.Time  `json:"contact_date"`
	Status           string     `json:"status"`
	WasAttended      bool       `json:"was_attended"`
	ConvertedAt      *time.Time `json:"converted_at"`
	Notes            *string    `json:"notes"`
	Origin           *string    `json:"origin"`
	GdprBlockContact bool       `json:"gdpr_block_contact"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// NewLeadResponse creates a new LeadResponse from a Lead model
func NewLeadResponse(lead model.Lead) LeadResponse {
	return LeadResponse{
		ID:               lead.ID,
		FullName:         lead.FullName,
		Phone:            lead.Phone,
		Email:            lead.Email,
		BirthDate:        lead.BirthDate,
		ContactDate:      lead.ContactDate,
		Status:           lead.Status,
		WasAttended:      lead.WasAttended,
		ConvertedAt:      lead.ConvertedAt,
		Notes:            lead.Notes,
		Origin:           lead.Origin,
		GdprBlockContact: lead.GdprBlockContact,
		CreatedAt:        lead.CreatedAt,
		UpdatedAt:        lead.UpdatedAt,
	}
}

// ConvertLeadRequest represents the request body for converting a lead to a patient
type ConvertLeadRequest struct {
	CostCenterID uuid.UUID `json:"cost_center_id" binding:"required"`
}
