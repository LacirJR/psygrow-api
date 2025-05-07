package model

import (
	"github.com/google/uuid"
	"time"
)

// PatientAnamnese represents a filled response for a patient based on a template
type PatientAnamnese struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PatientID  uuid.UUID `gorm:"type:uuid;not null"`
	AnamneseID uuid.UUID `gorm:"type:uuid;not null"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	AnsweredAt time.Time `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
