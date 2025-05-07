package port

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/google/uuid"
	"time"
)

type LeadRepository interface {
	// Create creates a new lead
	Create(lead *model.Lead) error

	// FindByID finds a lead by ID
	FindByID(id uuid.UUID, userID uuid.UUID) (*model.Lead, error)

	// FindAll finds all leads for a user
	FindAll(userID uuid.UUID, limit, offset int) ([]model.Lead, error)

	// Update updates a lead
	Update(lead *model.Lead) error

	// Delete deletes a lead
	Delete(id uuid.UUID, userID uuid.UUID) error

	// ConvertToPatient converts a lead to a patient
	ConvertToPatient(leadID uuid.UUID, userID uuid.UUID, costCenterID uuid.UUID) (*model.Patient, error)

	// FindByStatus finds leads by status
	FindByStatus(userID uuid.UUID, status string, limit, offset int) ([]model.Lead, error)

	// FindByOrigin finds leads by origin
	FindByOrigin(userID uuid.UUID, origin string, limit, offset int) ([]model.Lead, error)

	// FindByContactDate finds leads by contact date range
	FindByContactDate(userID uuid.UUID, startDate, endDate time.Time, limit, offset int) ([]model.Lead, error)

	// FindByWasAttended finds leads by was_attended flag
	FindByWasAttended(userID uuid.UUID, wasAttended bool, limit, offset int) ([]model.Lead, error)

	// Count counts all leads for a user
	Count(userID uuid.UUID) (int64, error)
}
