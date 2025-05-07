package port

import "github.com/LacirJR/psygrow-api/src/internal/core/model"

type CostCenterRepository interface {
	Save(costCenter *model.CostCenter) error
	FindByID(id string) (*model.CostCenter, error)
	FindByUserID(userID string) ([]*model.CostCenter, error)
	Update(costCenter *model.CostCenter) error
	Delete(id string) error
}

type PaymentRepository interface {
	Save(payment *model.Payment) error
	FindByID(id string) (*model.Payment, error)
	FindByUserID(userID string) ([]*model.Payment, error)
	FindByPatientID(patientID string) ([]*model.Payment, error)
	FindByCostCenterID(costCenterID string) ([]*model.Payment, error)
}

type PaymentAppointmentRepository interface {
	Save(paymentAppointment *model.PaymentAppointment) error
	FindByPaymentID(paymentID string) ([]*model.PaymentAppointment, error)
	FindByAppointmentID(appointmentID string) ([]*model.PaymentAppointment, error)
	Delete(id string) error
}

type RepasseRepository interface {
	Save(repasse *model.Repasse) error
	FindByID(id string) (*model.Repasse, error)
	FindByUserID(userID string) ([]*model.Repasse, error)
	FindByAppointmentID(appointmentID string) (*model.Repasse, error)
	FindByCostCenterID(costCenterID string) ([]*model.Repasse, error)
	FindByStatus(status string) ([]*model.Repasse, error)
	Update(repasse *model.Repasse) error
}
