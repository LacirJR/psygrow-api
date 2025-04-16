package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string    `gorm:"type:varchar(100);not null"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Role         string    `gorm:"type:varchar(50);default:'professional';not null"`
	Phone        *string   `gorm:"type:varchar(20)"`
	IsActive     bool      `gorm:"default:true"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	LastLoginAt  *time.Time
}
