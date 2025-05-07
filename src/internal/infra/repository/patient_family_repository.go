package repository

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type patientFamilyRepository struct {
	db *gorm.DB
}

func NewPatientFamilyRepository(db *gorm.DB) port.PatientFamilyRepository {
	return &patientFamilyRepository{db: db}
}

func (r *patientFamilyRepository) Create(patientFamily *model.PatientFamily) error {
	return r.db.Create(patientFamily).Error
}

func (r *patientFamilyRepository) FindByID(id uuid.UUID) (*model.PatientFamily, error) {
	var patientFamily model.PatientFamily
	err := r.db.Preload("Patient").Where("id = ?", id).First(&patientFamily).Error
	if err != nil {
		return nil, err
	}
	return &patientFamily, nil
}

func (r *patientFamilyRepository) FindByPatient(patientID uuid.UUID) ([]model.PatientFamily, error) {
	var patientFamilies []model.PatientFamily
	err := r.db.Where("patient_id = ?", patientID).
		Order("name ASC").
		Find(&patientFamilies).Error
	return patientFamilies, err
}

func (r *patientFamilyRepository) Update(patientFamily *model.PatientFamily) error {
	return r.db.Save(patientFamily).Error
}

func (r *patientFamilyRepository) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.PatientFamily{}).Error
}

func (r *patientFamilyRepository) FindByRelationship(patientID uuid.UUID, relationship string) ([]model.PatientFamily, error) {
	var patientFamilies []model.PatientFamily
	err := r.db.Where("patient_id = ? AND relationship = ?", patientID, relationship).
		Order("name ASC").
		Find(&patientFamilies).Error
	return patientFamilies, err
}