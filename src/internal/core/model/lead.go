package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// Lead represents a pre-registration of a person interested in starting treatment
type Lead struct {
	ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID           uuid.UUID  `gorm:"type:uuid;not null;index" validate:"required"`
	FullName         string     `gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	Phone            *string    `gorm:"type:varchar(20)"`
	Email            *string    `gorm:"type:varchar(100)"`
	BirthDate        *time.Time `gorm:"type:date"`
	ContactDate      time.Time  `gorm:"not null" validate:"required"`
	Status           string     `gorm:"type:varchar(20);default:new;not null;index" validate:"required,oneof=new in_analysis converted lost"`
	WasAttended      bool       `gorm:"default:false"`
	ConvertedAt      *time.Time
	Notes            *string   `gorm:"type:text"`
	Origin           *string   `gorm:"type:varchar(50)"`
	GdprBlockContact bool      `gorm:"default:false"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

// Validate performs validation on the Lead struct
func (l *Lead) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}
