package repository

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type leadRepository struct {
	db *gorm.DB
}

func NewLeadRepository(db *gorm.DB) port.LeadRepository {
	return &leadRepository{db: db}
}

func (r *leadRepository) Create(lead *model.Lead) error {
	return r.db.Create(lead).Error
}

func (r *leadRepository) FindByID(id uuid.UUID, userID uuid.UUID) (*model.Lead, error) {
	var lead model.Lead
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&lead).Error
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

func (r *leadRepository) FindAll(userID uuid.UUID, limit, offset int) ([]model.Lead, error) {
	var leads []model.Lead
	err := r.db.Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&leads).Error
	return leads, err
}

func (r *leadRepository) Update(lead *model.Lead) error {
	return r.db.Save(lead).Error
}

func (r *leadRepository) Delete(id uuid.UUID, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Lead{}).Error
}

func (r *leadRepository) ConvertToPatient(leadID uuid.UUID, userID uuid.UUID, costCenterID uuid.UUID) (*model.Patient, error) {
	// Start a transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find the lead
	var lead model.Lead
	if err := tx.Where("id = ? AND user_id = ?", leadID, userID).First(&lead).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create a new patient from the lead
	patient := model.Patient{
		ID:           uuid.New(),
		UserID:       userID,
		CostCenterID: costCenterID,
		FullName:     lead.FullName,
		BirthDate:    time.Now(), // Default value, should be updated if lead.BirthDate is not nil
		Phone:        lead.Phone,
		Email:        lead.Email,
		IsActive:     true,
	}

	// Set birth date if available
	if lead.BirthDate != nil {
		patient.BirthDate = *lead.BirthDate
	}

	// Create the patient
	if err := tx.Create(&patient).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update the lead status to converted
	now := time.Now()
	lead.Status = model.LeadStatusConverted
	lead.ConvertedAt = &now
	if err := tx.Save(&lead).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (r *leadRepository) FindByStatus(userID uuid.UUID, status string, limit, offset int) ([]model.Lead, error) {
	var leads []model.Lead
	err := r.db.Where("user_id = ? AND status = ?", userID, status).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&leads).Error
	return leads, err
}

func (r *leadRepository) FindByOrigin(userID uuid.UUID, origin string, limit, offset int) ([]model.Lead, error) {
	var leads []model.Lead
	err := r.db.Where("user_id = ? AND origin = ?", userID, origin).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&leads).Error
	return leads, err
}

func (r *leadRepository) FindByContactDate(userID uuid.UUID, startDate, endDate time.Time, limit, offset int) ([]model.Lead, error) {
	var leads []model.Lead
	err := r.db.Where("user_id = ? AND contact_date BETWEEN ? AND ?", userID, startDate, endDate).
		Limit(limit).
		Offset(offset).
		Order("contact_date DESC").
		Find(&leads).Error
	return leads, err
}

func (r *leadRepository) FindByWasAttended(userID uuid.UUID, wasAttended bool, limit, offset int) ([]model.Lead, error) {
	var leads []model.Lead
	err := r.db.Where("user_id = ? AND was_attended = ?", userID, wasAttended).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&leads).Error
	return leads, err
}

func (r *leadRepository) Count(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&model.Lead{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}