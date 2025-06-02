package repository

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnamneseTemplateRepository implementation
type anamneseTemplateRepository struct {
	db *gorm.DB
}

func NewAnamneseTemplateRepository(db *gorm.DB) port.AnamneseTemplateRepository {
	return &anamneseTemplateRepository{db: db}
}

func (r *anamneseTemplateRepository) Save(template *model.AnamneseTemplate) error {
	return r.db.Create(template).Error
}

func (r *anamneseTemplateRepository) FindByID(id string) (*model.AnamneseTemplate, error) {
	var template model.AnamneseTemplate
	templateID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", templateID).First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *anamneseTemplateRepository) FindByUserID(userID string) ([]*model.AnamneseTemplate, error) {
	var templates []*model.AnamneseTemplate
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&templates).Error
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func (r *anamneseTemplateRepository) Update(template *model.AnamneseTemplate) error {
	return r.db.Save(template).Error
}

func (r *anamneseTemplateRepository) Delete(id string) error {
	templateID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&model.AnamneseTemplate{}, templateID).Error
}

// AnamneseFieldRepository implementation
type anamneseFieldRepository struct {
	db *gorm.DB
}

func NewAnamneseFieldRepository(db *gorm.DB) port.AnamneseFieldRepository {
	return &anamneseFieldRepository{db: db}
}

func (r *anamneseFieldRepository) Save(field *model.AnamneseField) error {
	return r.db.Create(field).Error
}

func (r *anamneseFieldRepository) FindByID(id string) (*model.AnamneseField, error) {
	var field model.AnamneseField
	fieldID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", fieldID).Preload("Options").First(&field).Error
	if err != nil {
		return nil, err
	}
	return &field, nil
}

func (r *anamneseFieldRepository) FindByAnamneseID(anamneseID string) ([]*model.AnamneseField, error) {
	var fields []*model.AnamneseField
	parsedAnamneseID, err := uuid.Parse(anamneseID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("anamnese_id = ?", parsedAnamneseID).Preload("Options").Order("field_number").Find(&fields).Error
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func (r *anamneseFieldRepository) Update(field *model.AnamneseField) error {
	return r.db.Save(field).Error
}

func (r *anamneseFieldRepository) Delete(id string) error {
	fieldID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&model.AnamneseField{}, fieldID).Error
}

// PatientAnamneseRepository implementation
type patientAnamneseRepository struct {
	db *gorm.DB
}

func NewPatientAnamneseRepository(db *gorm.DB) port.PatientAnamneseRepository {
	return &patientAnamneseRepository{db: db}
}

func (r *patientAnamneseRepository) Save(patientAnamnese *model.PatientAnamnese) error {
	return r.db.Create(patientAnamnese).Error
}

func (r *patientAnamneseRepository) FindByID(id string) (*model.PatientAnamnese, error) {
	var patientAnamnese model.PatientAnamnese
	patientAnamneseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", patientAnamneseID).First(&patientAnamnese).Error
	if err != nil {
		return nil, err
	}
	return &patientAnamnese, nil
}

func (r *patientAnamneseRepository) FindByPatientID(patientID string) ([]*model.PatientAnamnese, error) {
	var patientAnamneses []*model.PatientAnamnese
	parsedPatientID, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("patient_id = ?", parsedPatientID).Find(&patientAnamneses).Error
	if err != nil {
		return nil, err
	}
	return patientAnamneses, nil
}

func (r *patientAnamneseRepository) FindByUserID(userID string) ([]*model.PatientAnamnese, error) {
	var patientAnamneses []*model.PatientAnamnese
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&patientAnamneses).Error
	if err != nil {
		return nil, err
	}
	return patientAnamneses, nil
}

// AnamneseFieldOptionRepository implementation
type anamneseFieldOptionRepository struct {
	db *gorm.DB
}

func NewAnamneseFieldOptionRepository(db *gorm.DB) port.AnamneseFieldOptionRepository {
	return &anamneseFieldOptionRepository{db: db}
}

func (r *anamneseFieldOptionRepository) Save(option *model.AnamneseFieldOption) error {
	return r.db.Create(option).Error
}

func (r *anamneseFieldOptionRepository) SaveBulk(options []*model.AnamneseFieldOption) error {
	return r.db.Create(options).Error
}

func (r *anamneseFieldOptionRepository) FindByID(id string) (*model.AnamneseFieldOption, error) {
	var option model.AnamneseFieldOption
	optionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", optionID).First(&option).Error
	if err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *anamneseFieldOptionRepository) FindByAnamneseFieldID(anamneseFieldID string) ([]*model.AnamneseFieldOption, error) {
	var options []*model.AnamneseFieldOption
	parsedAnamneseFieldID, err := uuid.Parse(anamneseFieldID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("anamnese_field_id = ?", parsedAnamneseFieldID).Order("option_order").Find(&options).Error
	if err != nil {
		return nil, err
	}
	return options, nil
}

func (r *anamneseFieldOptionRepository) Update(option *model.AnamneseFieldOption) error {
	return r.db.Save(option).Error
}

func (r *anamneseFieldOptionRepository) Delete(id string) error {
	optionID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&model.AnamneseFieldOption{}, optionID).Error
}

func (r *anamneseFieldOptionRepository) DeleteByAnamneseFieldID(anamneseFieldID string) error {
	parsedAnamneseFieldID, err := uuid.Parse(anamneseFieldID)
	if err != nil {
		return err
	}

	return r.db.Where("anamnese_field_id = ?", parsedAnamneseFieldID).Delete(&model.AnamneseFieldOption{}).Error
}

// PatientAnamneseFieldRepository implementation
type patientAnamneseFieldRepository struct {
	db *gorm.DB
}

func NewPatientAnamneseFieldRepository(db *gorm.DB) port.PatientAnamneseFieldRepository {
	return &patientAnamneseFieldRepository{db: db}
}

func (r *patientAnamneseFieldRepository) Save(field *model.PatientAnamneseField) error {
	return r.db.Create(field).Error
}

func (r *patientAnamneseFieldRepository) FindByPatientAnamneseID(patientAnamneseID string) ([]*model.PatientAnamneseField, error) {
	var fields []*model.PatientAnamneseField
	parsedPatientAnamneseID, err := uuid.Parse(patientAnamneseID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("patient_anamnese_id = ?", parsedPatientAnamneseID).Find(&fields).Error
	if err != nil {
		return nil, err
	}
	return fields, nil
}
