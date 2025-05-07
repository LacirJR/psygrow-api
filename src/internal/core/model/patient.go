package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// Patient represents a patient with active or previous treatment
type Patient struct {
	ID                    uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID                uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	CostCenterID          uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	CostCenter            CostCenter `gorm:"foreignKey:CostCenterID"`
	FullName              string     `gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	SocialName            *string    `gorm:"type:varchar(100)"`
	BirthDate             time.Time  `gorm:"type:date;not null" validate:"required"`
	Document              *string    `gorm:"type:varchar(20)"`
	Phone                 *string    `gorm:"type:varchar(20)"`
	Email                 *string    `gorm:"type:varchar(100)"`
	Gender                *string    `gorm:"type:varchar(20)"`
	Address               *string    `gorm:"type:varchar(255)"`
	ResidesWith           *string    `gorm:"type:varchar(100)"`
	EmergencyContactName  *string    `gorm:"type:varchar(100)"`
	EmergencyContactPhone *string    `gorm:"type:varchar(20)"`
	Observation           *string    `gorm:"type:text"`
	DefaultRepasseType    *string    `gorm:"type:varchar(20)" validate:"omitempty,oneof=percent fixed"`
	DefaultRepasseValue   *int64     `gorm:"type:bigint"` // Stored as cents or basis points (for percent)
	IsActive              bool       `gorm:"column:active;default:true"`
	CreatedAt             time.Time  `gorm:"autoCreateTime"`
	UpdatedAt             time.Time  `gorm:"autoUpdateTime"`
}

// Validate performs validation on the Patient struct
func (p *Patient) Validate() error {
	validate := validator.New()

	// Custom validation for DefaultRepasseValue based on DefaultRepasseType
	err := validate.RegisterValidation("default_repasse_value_valid", func(fl validator.FieldLevel) bool {
		// Get the parent struct
		patient, ok := fl.Parent().Interface().(Patient)
		if !ok {
			return false
		}

		// If DefaultRepasseType is not set, then DefaultRepasseValue should also not be set
		if patient.DefaultRepasseType == nil {
			return patient.DefaultRepasseValue == nil
		}

		// If DefaultRepasseValue is not set, that's an error when DefaultRepasseType is set
		if patient.DefaultRepasseValue == nil {
			return false
		}

		// If DefaultRepasseType is percent, ensure value is between 0 and 10000 (0-100%)
		if *patient.DefaultRepasseType == RepasseTypePercent {
			value := *patient.DefaultRepasseValue
			return value >= 0 && value <= 10000 // 0-100% with 2 decimal places (e.g., 10.50% = 1050)
		}

		// For fixed type, just ensure it's not negative
		return *patient.DefaultRepasseValue >= 0
	})

	if err != nil {
		return err
	}

	return validate.Struct(p)
}
