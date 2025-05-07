package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

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
