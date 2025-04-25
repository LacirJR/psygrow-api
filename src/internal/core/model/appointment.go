package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// Appointment represents a scheduled appointment between a professional and a patient
type Appointment struct {
	ID                 uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID             uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	PatientID          uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	CustomRepasseType  *string    `gorm:"type:varchar(20)" validate:"omitempty,oneof=percent fixed"` // Use constants from model package
	CustomRepasseValue *int64     `gorm:"type:bigint"`                                               // Stored as cents or basis points (for percent)
	ProfessionalID     uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	CostCenterID       uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	CostCenter         CostCenter `gorm:"foreignKey:CostCenterID"`
	ServiceTitle       string     `gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	StartTime          time.Time  `gorm:"not null" validate:"required"`
	EndTime            time.Time  `gorm:"not null" validate:"required,gtfield=StartTime"`
	Status             string     `gorm:"type:varchar(20);default:scheduled;not null;index" validate:"required,oneof=scheduled done canceled no_show"` // Use constants from model package
	Notes              string     `gorm:"type:text" validate:"max=1000"`
	CreatedAt          time.Time  `gorm:"autoCreateTime"`
	UpdatedAt          time.Time  `gorm:"autoUpdateTime"`
}

// Validate performs validation on the Appointment struct
func (a *Appointment) Validate() error {
	validate := validator.New()

	// Custom validation for CustomRepasseValue based on CustomRepasseType
	validate.RegisterValidation("custom_repasse_value_valid", func(fl validator.FieldLevel) bool {
		// Get the parent struct
		appointment, ok := fl.Parent().Interface().(Appointment)
		if !ok {
			return false
		}

		// If CustomRepasseType is not set, then CustomRepasseValue should also not be set
		if appointment.CustomRepasseType == nil {
			return appointment.CustomRepasseValue == nil
		}

		// If CustomRepasseValue is not set, that's an error when CustomRepasseType is set
		if appointment.CustomRepasseValue == nil {
			return false
		}

		// If CustomRepasseType is percent, ensure value is between 0 and 10000 (0-100%)
		if *appointment.CustomRepasseType == RepasseTypePercent {
			value := *appointment.CustomRepasseValue
			return value >= 0 && value <= 10000 // 0-100% with 2 decimal places (e.g., 10.50% = 1050)
		}

		// For fixed type, just ensure it's not negative
		return *appointment.CustomRepasseValue >= 0
	})

	return validate.Struct(a)
}

// Session represents an actual session that occurred when an appointment is marked as done
type Session struct {
	ID             uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	AppointmentID  uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	Appointment    Appointment `gorm:"foreignKey:AppointmentID"`
	UserID         uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	PatientID      uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	ProfessionalID uuid.UUID   `gorm:"type:uuid;not null;index" validate:"required"`
	StartTime      time.Time   `gorm:"not null" validate:"required"`
	EndTime        time.Time   `gorm:"not null" validate:"required,gtfield=StartTime"`
	WasAttended    bool        `gorm:"default:true"`
	CreatedAt      time.Time   `gorm:"autoCreateTime"`
}

// Validate performs validation on the Session struct
func (s *Session) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

// Evolution represents clinical notes generated only if the session was conducted
type Evolution struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SessionID      uuid.UUID `gorm:"type:uuid;not null;index" validate:"required"`
	Session        Session   `gorm:"foreignKey:SessionID"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;index" validate:"required"`
	ProfessionalID uuid.UUID `gorm:"type:uuid;not null;index" validate:"required"`
	PatientID      uuid.UUID `gorm:"type:uuid;not null;index" validate:"required"`
	Content        string    `gorm:"type:text;not null" validate:"required,min=1"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

// Validate performs validation on the Evolution struct
func (e *Evolution) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}
