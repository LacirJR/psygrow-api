package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// CostCenter defines the origin of the service and the associated repasse rule.
// It represents where the appointment takes place (clinic, private practice, institution, etc.)
// and defines how financial compensation is handled.
type CostCenter struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index" validate:"required"`
	Name         string    `gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	RepasseModel string    `gorm:"type:varchar(20);not null" validate:"required,oneof=clinic_pays professional_pays"` // Use constants from model package
	RepasseType  string    `gorm:"type:varchar(20);not null" validate:"required,oneof=percent fixed"`                 // Use constants from model package
	RepasseValue int64     `gorm:"type:bigint;not null" validate:"required,min=0"`                                    // Stored as cents (e.g., $10.50 = 1050)
	IsActive     bool      `gorm:"column:active;default:true"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

// Validate performs validation on the CostCenter struct
func (c *CostCenter) Validate() error {
	validate := validator.New()

	// Custom validation for RepasseValue based on RepasseType
	validate.RegisterValidation("repasse_value_valid", func(fl validator.FieldLevel) bool {
		// Get the parent struct
		costCenter, ok := fl.Parent().Interface().(CostCenter)
		if !ok {
			return false
		}

		// If RepasseType is percent, ensure value is between 0 and 10000 (0-100%)
		if costCenter.RepasseType == RepasseTypePercent {
			value := costCenter.RepasseValue
			return value >= 0 && value <= 10000 // 0-100% with 2 decimal places (e.g., 10.50% = 1050)
		}

		// For fixed type, just ensure it's not negative
		return costCenter.RepasseValue >= 0
	})

	return validate.Struct(c)
}

// Payment represents a financial entry record for services provided
type Payment struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	PatientID    *uuid.UUID `gorm:"type:uuid;index"`
	CostCenterID uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	CostCenter   CostCenter `gorm:"foreignKey:CostCenterID"`
	PaymentDate  time.Time  `gorm:"not null" validate:"required"`
	Amount       int64      `gorm:"type:bigint;not null" validate:"required,min=1"`                          // Stored as cents (e.g., $10.50 = 1050)
	Method       string     `gorm:"type:varchar(50);not null" validate:"required,oneof=pix cash card other"` // Use constants from model package
	Notes        string     `gorm:"type:text" validate:"max=1000"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
}

// Validate performs validation on the Payment struct
func (p *Payment) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

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

// Repasse defines the amount that the professional needs to pass on to the clinic or institution
type Repasse struct {
	ID                uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID            uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	AppointmentID     uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	Appointment       Appointment `gorm:"foreignKey:AppointmentID"`
	CostCenterID      uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	CostCenter        CostCenter  `gorm:"foreignKey:CostCenterID"`
	Value             int64       `gorm:"type:bigint;not null" validate:"required,min=0"`                                       // Stored as cents (e.g., $10.50 = 1050)
	DoesClinicReceive bool        `gorm:"column:clinic_receives;default:true"`                                                  // If true, professional pays clinic; if false, clinic already retained amount
	Status            string      `gorm:"type:varchar(20);not null;index" validate:"required,oneof=pending paid informational"` // Use constants from model package
	PaidAt            *time.Time
	Notes             string    `gorm:"type:text" validate:"max=1000"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}

// Validate performs validation on the Repasse struct
func (r *Repasse) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
