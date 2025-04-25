package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

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
