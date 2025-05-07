package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// PatientFamily represents a family member of a patient
type PatientFamily struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PatientID    uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	Patient      Patient    `gorm:"foreignKey:PatientID"`
	Relationship string     `gorm:"type:varchar(50);not null" validate:"required"`
	Name         string     `gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	BirthDate    *time.Time `gorm:"type:date"`
	Schooling    *string    `gorm:"type:varchar(50)"`
	Occupation   *string    `gorm:"type:varchar(100)"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}

// Validate performs validation on the PatientFamily struct
func (pf *PatientFamily) Validate() error {
	validate := validator.New()
	return validate.Struct(pf)
}
