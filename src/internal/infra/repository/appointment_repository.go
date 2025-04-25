package repository

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AppointmentRepository implementation
type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) port.AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) Save(appointment *model.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) FindByID(id string) (*model.Appointment, error) {
	var appointment model.Appointment
	appointmentID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", appointmentID).First(&appointment).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *appointmentRepository) FindByUserID(userID string) ([]*model.Appointment, error) {
	var appointments []*model.Appointment
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *appointmentRepository) FindByPatientID(patientID string) ([]*model.Appointment, error) {
	var appointments []*model.Appointment
	parsedPatientID, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("patient_id = ?", parsedPatientID).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *appointmentRepository) FindByProfessionalID(professionalID string) ([]*model.Appointment, error) {
	var appointments []*model.Appointment
	parsedProfessionalID, err := uuid.Parse(professionalID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("professional_id = ?", parsedProfessionalID).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *appointmentRepository) Update(appointment *model.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *appointmentRepository) Delete(id string) error {
	appointmentID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&model.Appointment{}, appointmentID).Error
}

// SessionRepository implementation
type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) port.SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Save(session *model.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepository) FindByID(id string) (*model.Session, error) {
	var session model.Session
	sessionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", sessionID).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepository) FindByAppointmentID(appointmentID string) (*model.Session, error) {
	var session model.Session
	parsedAppointmentID, err := uuid.Parse(appointmentID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("appointment_id = ?", parsedAppointmentID).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepository) FindByUserID(userID string) ([]*model.Session, error) {
	var sessions []*model.Session
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *sessionRepository) FindByPatientID(patientID string) ([]*model.Session, error) {
	var sessions []*model.Session
	parsedPatientID, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("patient_id = ?", parsedPatientID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *sessionRepository) FindByProfessionalID(professionalID string) ([]*model.Session, error) {
	var sessions []*model.Session
	parsedProfessionalID, err := uuid.Parse(professionalID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("professional_id = ?", parsedProfessionalID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

// EvolutionRepository implementation
type evolutionRepository struct {
	db *gorm.DB
}

func NewEvolutionRepository(db *gorm.DB) port.EvolutionRepository {
	return &evolutionRepository{db: db}
}

func (r *evolutionRepository) Save(evolution *model.Evolution) error {
	return r.db.Create(evolution).Error
}

func (r *evolutionRepository) FindByID(id string) (*model.Evolution, error) {
	var evolution model.Evolution
	evolutionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", evolutionID).First(&evolution).Error
	if err != nil {
		return nil, err
	}
	return &evolution, nil
}

func (r *evolutionRepository) FindBySessionID(sessionID string) (*model.Evolution, error) {
	var evolution model.Evolution
	parsedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("session_id = ?", parsedSessionID).First(&evolution).Error
	if err != nil {
		return nil, err
	}
	return &evolution, nil
}

func (r *evolutionRepository) FindByUserID(userID string) ([]*model.Evolution, error) {
	var evolutions []*model.Evolution
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&evolutions).Error
	if err != nil {
		return nil, err
	}
	return evolutions, nil
}

func (r *evolutionRepository) FindByPatientID(patientID string) ([]*model.Evolution, error) {
	var evolutions []*model.Evolution
	parsedPatientID, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("patient_id = ?", parsedPatientID).Find(&evolutions).Error
	if err != nil {
		return nil, err
	}
	return evolutions, nil
}

func (r *evolutionRepository) FindByProfessionalID(professionalID string) ([]*model.Evolution, error) {
	var evolutions []*model.Evolution
	parsedProfessionalID, err := uuid.Parse(professionalID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("professional_id = ?", parsedProfessionalID).Find(&evolutions).Error
	if err != nil {
		return nil, err
	}
	return evolutions, nil
}
