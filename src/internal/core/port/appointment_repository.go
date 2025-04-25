package port

import "github.com/LacirJR/psygrow-api/src/internal/core/model"

type AppointmentRepository interface {
	Save(appointment *model.Appointment) error
	FindByID(id string) (*model.Appointment, error)
	FindByUserID(userID string) ([]*model.Appointment, error)
	FindByPatientID(patientID string) ([]*model.Appointment, error)
	FindByProfessionalID(professionalID string) ([]*model.Appointment, error)
	Update(appointment *model.Appointment) error
	Delete(id string) error
}

type SessionRepository interface {
	Save(session *model.Session) error
	FindByID(id string) (*model.Session, error)
	FindByAppointmentID(appointmentID string) (*model.Session, error)
	FindByUserID(userID string) ([]*model.Session, error)
	FindByPatientID(patientID string) ([]*model.Session, error)
	FindByProfessionalID(professionalID string) ([]*model.Session, error)
}

type EvolutionRepository interface {
	Save(evolution *model.Evolution) error
	FindByID(id string) (*model.Evolution, error)
	FindBySessionID(sessionID string) (*model.Evolution, error)
	FindByUserID(userID string) ([]*model.Evolution, error)
	FindByPatientID(patientID string) ([]*model.Evolution, error)
	FindByProfessionalID(professionalID string) ([]*model.Evolution, error)
}
