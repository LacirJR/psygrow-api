package repository

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CostCenterRepository implementation
type costCenterRepository struct {
	db *gorm.DB
}

func NewCostCenterRepository(db *gorm.DB) port.CostCenterRepository {
	return &costCenterRepository{db: db}
}

func (r *costCenterRepository) Save(costCenter *model.CostCenter) error {
	return r.db.Create(costCenter).Error
}

func (r *costCenterRepository) FindByID(id string) (*model.CostCenter, error) {
	var costCenter model.CostCenter
	costCenterID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", costCenterID).First(&costCenter).Error
	if err != nil {
		return nil, err
	}
	return &costCenter, nil
}

func (r *costCenterRepository) FindByUserID(userID string) ([]*model.CostCenter, error) {
	var costCenters []*model.CostCenter
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&costCenters).Error
	if err != nil {
		return nil, err
	}
	return costCenters, nil
}

func (r *costCenterRepository) Update(costCenter *model.CostCenter) error {
	return r.db.Save(costCenter).Error
}

func (r *costCenterRepository) Delete(id string) error {
	costCenterID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&model.CostCenter{}, costCenterID).Error
}

// PaymentRepository implementation
type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) port.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Save(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) FindByID(id string) (*model.Payment, error) {
	var payment model.Payment
	paymentID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", paymentID).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) FindByUserID(userID string) ([]*model.Payment, error) {
	var payments []*model.Payment
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) FindByPatientID(patientID string) ([]*model.Payment, error) {
	var payments []*model.Payment
	parsedPatientID, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("patient_id = ?", parsedPatientID).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) FindByCostCenterID(costCenterID string) ([]*model.Payment, error) {
	var payments []*model.Payment
	parsedCostCenterID, err := uuid.Parse(costCenterID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("cost_center_id = ?", parsedCostCenterID).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// PaymentAppointmentRepository implementation
type paymentAppointmentRepository struct {
	db *gorm.DB
}

func NewPaymentAppointmentRepository(db *gorm.DB) port.PaymentAppointmentRepository {
	return &paymentAppointmentRepository{db: db}
}

func (r *paymentAppointmentRepository) Save(paymentAppointment *model.PaymentAppointment) error {
	return r.db.Create(paymentAppointment).Error
}

func (r *paymentAppointmentRepository) FindByPaymentID(paymentID string) ([]*model.PaymentAppointment, error) {
	var paymentAppointments []*model.PaymentAppointment
	parsedPaymentID, err := uuid.Parse(paymentID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("payment_id = ?", parsedPaymentID).Find(&paymentAppointments).Error
	if err != nil {
		return nil, err
	}
	return paymentAppointments, nil
}

func (r *paymentAppointmentRepository) FindByAppointmentID(appointmentID string) ([]*model.PaymentAppointment, error) {
	var paymentAppointments []*model.PaymentAppointment
	parsedAppointmentID, err := uuid.Parse(appointmentID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("appointment_id = ?", parsedAppointmentID).Find(&paymentAppointments).Error
	if err != nil {
		return nil, err
	}
	return paymentAppointments, nil
}

func (r *paymentAppointmentRepository) Delete(id string) error {
	paymentAppointmentID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&model.PaymentAppointment{}, paymentAppointmentID).Error
}

// RepasseRepository implementation
type repasseRepository struct {
	db *gorm.DB
}

func NewRepasseRepository(db *gorm.DB) port.RepasseRepository {
	return &repasseRepository{db: db}
}

func (r *repasseRepository) Save(repasse *model.Repasse) error {
	return r.db.Create(repasse).Error
}

func (r *repasseRepository) FindByID(id string) (*model.Repasse, error) {
	var repasse model.Repasse
	repasseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", repasseID).First(&repasse).Error
	if err != nil {
		return nil, err
	}
	return &repasse, nil
}

func (r *repasseRepository) FindByUserID(userID string) ([]*model.Repasse, error) {
	var repasses []*model.Repasse
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("user_id = ?", parsedUserID).Find(&repasses).Error
	if err != nil {
		return nil, err
	}
	return repasses, nil
}

func (r *repasseRepository) FindByAppointmentID(appointmentID string) (*model.Repasse, error) {
	var repasse model.Repasse
	parsedAppointmentID, err := uuid.Parse(appointmentID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("appointment_id = ?", parsedAppointmentID).First(&repasse).Error
	if err != nil {
		return nil, err
	}
	return &repasse, nil
}

func (r *repasseRepository) FindByCostCenterID(costCenterID string) ([]*model.Repasse, error) {
	var repasses []*model.Repasse
	parsedCostCenterID, err := uuid.Parse(costCenterID)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("cost_center_id = ?", parsedCostCenterID).Find(&repasses).Error
	if err != nil {
		return nil, err
	}
	return repasses, nil
}

func (r *repasseRepository) FindByStatus(status string) ([]*model.Repasse, error) {
	var repasses []*model.Repasse

	err := r.db.Where("status = ?", status).Find(&repasses).Error
	if err != nil {
		return nil, err
	}
	return repasses, nil
}

func (r *repasseRepository) Update(repasse *model.Repasse) error {
	return r.db.Save(repasse).Error
}
