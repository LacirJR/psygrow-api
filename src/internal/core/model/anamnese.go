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
