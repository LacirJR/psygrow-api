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
	err := validate.RegisterValidation("repasse_value_valid", func(fl validator.FieldLevel) bool {
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

	if err != nil {
		return err
	}

	return validate.Struct(c)
}
