package port

import "github.com/LacirJR/psygrow-api/src/internal/core/model"

type AnamneseTemplateRepository interface {
	Save(template *model.AnamneseTemplate) error
	FindByID(id string) (*model.AnamneseTemplate, error)
	FindByUserID(userID string) ([]*model.AnamneseTemplate, error)
	Update(template *model.AnamneseTemplate) error
	Delete(id string) error
}

type AnamneseFieldRepository interface {
	Save(field *model.AnamneseField) error
	FindByID(id string) (*model.AnamneseField, error)
	FindByAnamneseID(anamneseID string) ([]*model.AnamneseField, error)
	Update(field *model.AnamneseField) error
	Delete(id string) error
}

type AnamneseFieldOptionRepository interface {
	Save(option *model.AnamneseFieldOption) error
	SaveBulk(options []*model.AnamneseFieldOption) error
	FindByID(id string) (*model.AnamneseFieldOption, error)
	FindByAnamneseFieldID(anamneseFieldID string) ([]*model.AnamneseFieldOption, error)
	Update(option *model.AnamneseFieldOption) error
	Delete(id string) error
	DeleteByAnamneseFieldID(anamneseFieldID string) error
}

type PatientAnamneseRepository interface {
	Save(patientAnamnese *model.PatientAnamnese) error
	FindByID(id string) (*model.PatientAnamnese, error)
	FindByPatientID(patientID string) ([]*model.PatientAnamnese, error)
	FindByUserID(userID string) ([]*model.PatientAnamnese, error)
}

type PatientAnamneseFieldRepository interface {
	Save(field *model.PatientAnamneseField) error
	FindByPatientAnamneseID(patientAnamneseID string) ([]*model.PatientAnamneseField, error)
}
