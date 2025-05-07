package port

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/google/uuid"
)

type PatientFamilyRepository interface {
	// Create creates a new patient family member
	Create(patientFamily *model.PatientFamily) error

	// FindByID finds a patient family member by ID
	FindByID(id uuid.UUID) (*model.PatientFamily, error)

	// FindByPatient finds all family members for a patient
	FindByPatient(patientID uuid.UUID) ([]model.PatientFamily, error)

	// Update updates a patient family member
	Update(patientFamily *model.PatientFamily) error

	// Delete deletes a patient family member
	Delete(id uuid.UUID) error

	// FindByRelationship finds family members by relationship type
	FindByRelationship(patientID uuid.UUID, relationship string) ([]model.PatientFamily, error)
}
