package model

import (
	"github.com/google/uuid"
	"time"
)

// PatientAnamneseField represents individual responses for each field of the filled anamnesis
type PatientAnamneseField struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PatientAnamneseID uuid.UUID `gorm:"type:uuid;not null"`
	FieldID           uuid.UUID `gorm:"type:uuid;not null"`
	Value             string    `gorm:"type:text"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}
