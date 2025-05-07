package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// PaymentAppointment links a payment to one or more specific sessions
type PaymentAppointment struct {
	ID            uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PaymentID     uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	Payment       Payment     `gorm:"foreignKey:PaymentID"`
	AppointmentID uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	Appointment   Appointment `gorm:"foreignKey:AppointmentID"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
}

// Validate performs validation on the PaymentAppointment struct
func (pa *PaymentAppointment) Validate() error {
	validate := validator.New()
	return validate.Struct(pa)
}
