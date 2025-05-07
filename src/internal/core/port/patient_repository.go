package port

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/google/uuid"
)

type PatientRepository interface {
	// Create creates a new patient
	Create(patient *model.Patient) error

	// FindByID finds a patient by ID
	FindByID(id uuid.UUID, userID uuid.UUID) (*model.Patient, error)

	// FindAll finds all patients for a user
	FindAll(userID uuid.UUID, limit, offset int) ([]model.Patient, error)

	// Update updates a patient
	Update(patient *model.Patient) error

	// Delete deletes a patient
	Delete(id uuid.UUID, userID uuid.UUID) error

	// FindByName finds patients by name (partial match)
	FindByName(userID uuid.UUID, name string, limit, offset int) ([]model.Patient, error)

	// FindByEmail finds a patient by email
	FindByEmail(userID uuid.UUID, email string) (*model.Patient, error)

	// FindByPhone finds a patient by phone
	FindByPhone(userID uuid.UUID, phone string) (*model.Patient, error)

	// FindByDocument finds a patient by document
	FindByDocument(userID uuid.UUID, document string) (*model.Patient, error)

	// FindByCostCenter finds patients by cost center
	FindByCostCenter(userID uuid.UUID, costCenterID uuid.UUID, limit, offset int) ([]model.Patient, error)

	// FindActive finds active patients
	FindActive(userID uuid.UUID, limit, offset int) ([]model.Patient, error)

	// FindInactive finds inactive patients
	FindInactive(userID uuid.UUID, limit, offset int) ([]model.Patient, error)

	// Count counts all patients for a user
	Count(userID uuid.UUID) (int64, error)
}
