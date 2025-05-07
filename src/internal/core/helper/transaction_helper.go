package helper

import (
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TransactionManager provides a way to execute operations within a transaction
type TransactionManager struct {
	DB *gorm.DB
}

// NewTransactionManager creates a new TransactionManager
func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{DB: db}
}

// WithTransaction executes the given function within a transaction
// If the function returns an error, the transaction is rolled back
// Otherwise, the transaction is committed
func (tm *TransactionManager) WithTransaction(fn func(tx *gorm.DB) error) error {
	tx := tm.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// SavePaymentWithAppointments saves a payment and its associated appointments within a transaction
func (tm *TransactionManager) SavePaymentWithAppointments(payment *model.Payment, appointmentIDs []string) error {
	return tm.WithTransaction(func(tx *gorm.DB) error {
		// Validate payment
		if err := payment.Validate(); err != nil {
			return err
		}

		// Save payment
		if err := tx.Create(payment).Error; err != nil {
			return err
		}

		// Link payment to appointments if provided
		if appointmentIDs != nil && len(appointmentIDs) > 0 {
			for _, appointmentIDStr := range appointmentIDs {
				appointmentID, err := uuid.Parse(appointmentIDStr)
				if err != nil {
					return err
				}

				paymentAppointment := &model.PaymentAppointment{
					ID:            uuid.New(),
					PaymentID:     payment.ID,
					AppointmentID: appointmentID,
				}

				// Validate payment appointment
				if err := paymentAppointment.Validate(); err != nil {
					return err
				}

				if err := tx.Create(paymentAppointment).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

/*
TODO: Implementation Notes

1. Repository Interface Updates:
   - Add transaction-aware methods to repository interfaces
   - Example: `SaveTx(tx *gorm.DB, entity *model.Entity) error`

2. Repository Implementation Updates:
   - Implement transaction-aware methods in repository implementations
   - Example: 
     ```go
     func (r *repository) SaveTx(tx *gorm.DB, entity *model.Entity) error {
         return tx.Create(entity).Error
     }
     ```

3. Handler Updates:
   - Use TransactionManager in handlers for operations that need transactions
   - Example:
     ```go
     txManager := helper.NewTransactionManager(config.DB)
     err := txManager.SavePaymentWithAppointments(payment, req.AppointmentIDs)
     if err != nil {
         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment", "details": err.Error()})
         return
     }
     ```

4. DTO Updates:
   - Update DTOs to match model changes (int64 for monetary values, etc.)
   - Add validation tags to DTOs

5. Migration Updates:
   - Update database migrations to add indices and change column types
*/
