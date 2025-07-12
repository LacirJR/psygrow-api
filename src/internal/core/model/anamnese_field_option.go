package model

import (
	"github.com/google/uuid"
	"time"
)

// AnamneseFieldOption represents options for anamnese fields
type AnamneseFieldOption struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	AnamneseFieldID uuid.UUID `gorm:"type:uuid;not null"` // Foreign key to AnamneseField
	OptionValue     string    `gorm:"type:varchar(255);not null"`
	OptionOrder     int       `gorm:"not null"` // To maintain the order of options
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
