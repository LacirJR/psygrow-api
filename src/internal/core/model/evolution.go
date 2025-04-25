package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

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
