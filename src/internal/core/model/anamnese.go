package model

import (
	"github.com/google/uuid"
	"time"
)

// AnamneseTemplate represents a template for anamnesis configured by each client
type AnamneseTemplate struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `gorm:"type:varchar(100);not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

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

// PatientAnamneseField represents individual responses for each field of the filled anamnesis
type PatientAnamneseField struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PatientAnamneseID uuid.UUID `gorm:"type:uuid;not null"`
	FieldID           uuid.UUID `gorm:"type:uuid;not null"`
	Value             string    `gorm:"type:text"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}
