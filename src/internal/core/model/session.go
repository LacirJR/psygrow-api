package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

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
