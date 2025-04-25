package dto

import (
	"github.com/google/uuid"
	"time"
)

// PaymentRequest represents the request to create a new payment
type PaymentRequest struct {
	PatientID      *string   `json:"patient_id,omitempty" binding:"omitempty,uuid"`
	CostCenterID   string    `json:"cost_center_id" binding:"required,uuid"`
	PaymentDate    time.Time `json:"payment_date" binding:"required"`
	Amount         int64     `json:"amount" binding:"required,gt=0"`
	Method         string    `json:"method" binding:"required"`
	Notes          string    `json:"notes"`
	AppointmentIDs []string  `json:"appointment_ids,omitempty" binding:"omitempty,dive,uuid"`
}

// PaymentResponse represents the response for a payment
type PaymentResponse struct {
	ID           string                       `json:"id"`
	UserID       string                       `json:"user_id"`
	PatientID    *string                      `json:"patient_id,omitempty"`
	CostCenterID string                       `json:"cost_center_id"`
	PaymentDate  time.Time                    `json:"payment_date"`
	Amount       int64                        `json:"amount"`
	Method       string                       `json:"method"`
	Notes        string                       `json:"notes,omitempty"`
	CreatedAt    time.Time                    `json:"created_at"`
	Appointments []PaymentAppointmentResponse `json:"appointments,omitempty"`
}

// PaymentAppointmentResponse represents the response for a payment appointment
type PaymentAppointmentResponse struct {
	ID            string `json:"id"`
	PaymentID     string `json:"payment_id"`
	AppointmentID string `json:"appointment_id"`
}

// NewPaymentResponse creates a new PaymentResponse from the given parameters
func NewPaymentResponse(
	id uuid.UUID,
	userID uuid.UUID,
	patientID *uuid.UUID,
	costCenterID uuid.UUID,
	paymentDate time.Time,
	amount int64,
	method string,
	notes string,
	createdAt time.Time,
) PaymentResponse {
	var patientIDStr *string
	if patientID != nil {
		idStr := patientID.String()
		patientIDStr = &idStr
	}

	return PaymentResponse{
		ID:           id.String(),
		UserID:       userID.String(),
		PatientID:    patientIDStr,
		CostCenterID: costCenterID.String(),
		PaymentDate:  paymentDate,
		Amount:       amount,
		Method:       method,
		Notes:        notes,
		CreatedAt:    createdAt,
		Appointments: []PaymentAppointmentResponse{},
	}
}

// NewPaymentAppointmentResponse creates a new PaymentAppointmentResponse from the given parameters
func NewPaymentAppointmentResponse(
	id uuid.UUID,
	paymentID uuid.UUID,
	appointmentID uuid.UUID,
) PaymentAppointmentResponse {
	return PaymentAppointmentResponse{
		ID:            id.String(),
		PaymentID:     paymentID.String(),
		AppointmentID: appointmentID.String(),
	}
}
