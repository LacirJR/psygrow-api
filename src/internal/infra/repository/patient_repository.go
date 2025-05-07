package repository

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) port.PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Create(patient *model.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepository) FindByID(id uuid.UUID, userID uuid.UUID) (*model.Patient, error) {
	var patient model.Patient
	err := r.db.Preload("CostCenter").Where("id = ? AND user_id = ?", id, userID).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) FindAll(userID uuid.UUID, limit, offset int) ([]model.Patient, error) {
	var patients []model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Order("full_name ASC").
		Find(&patients).Error
	return patients, err
}

func (r *patientRepository) Update(patient *model.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepository) Delete(id uuid.UUID, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Patient{}).Error
}

func (r *patientRepository) FindByName(userID uuid.UUID, name string, limit, offset int) ([]model.Patient, error) {
	var patients []model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ? AND full_name ILIKE ?", userID, "%"+name+"%").
		Limit(limit).
		Offset(offset).
		Order("full_name ASC").
		Find(&patients).Error
	return patients, err
}

func (r *patientRepository) FindByEmail(userID uuid.UUID, email string) (*model.Patient, error) {
	var patient model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ? AND email = ?", userID, email).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) FindByPhone(userID uuid.UUID, phone string) (*model.Patient, error) {
	var patient model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ? AND phone = ?", userID, phone).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) FindByDocument(userID uuid.UUID, document string) (*model.Patient, error) {
	var patient model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ? AND document = ?", userID, document).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) FindByCostCenter(userID uuid.UUID, costCenterID uuid.UUID, limit, offset int) ([]model.Patient, error) {
	var patients []model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ? AND cost_center_id = ?", userID, costCenterID).
		Limit(limit).
		Offset(offset).
		Order("full_name ASC").
		Find(&patients).Error
	return patients, err
}

func (r *patientRepository) FindActive(userID uuid.UUID, limit, offset int) ([]model.Patient, error) {
	var patients []model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ? AND active = ?", userID, true).
		Limit(limit).
		Offset(offset).
		Order("full_name ASC").
		Find(&patients).Error
	return patients, err
}

func (r *patientRepository) FindInactive(userID uuid.UUID, limit, offset int) ([]model.Patient, error) {
	var patients []model.Patient
	err := r.db.Preload("CostCenter").Where("user_id = ? AND active = ?", userID, false).
		Limit(limit).
		Offset(offset).
		Order("full_name ASC").
		Find(&patients).Error
	return patients, err
}

func (r *patientRepository) Count(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&model.Patient{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}