package model

import (
	"github.com/google/uuid"
	"time"
)

// AnamneseField represents custom fields that make up the anamnesis template
type AnamneseField struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FieldNumber   int       `gorm:"not null"`
	FieldType     string    `gorm:"type:varchar(50);not null"` // date, datetime, text, number, checkbox
	FieldTitle    string    `gorm:"type:varchar(255);not null"`
	FieldRequired bool      `gorm:"default:false"`
	FieldActive   bool      `gorm:"default:true"`
	UserID        uuid.UUID `gorm:"type:uuid;not null"`
	AnamneseID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
