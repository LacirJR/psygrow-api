package handler

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	dto "github.com/LacirJR/psygrow-api/src/internal/core/dto/financial"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/infra/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// CreatePayment handles the creation of a new payment
func CreatePayment(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req dto.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validate payment method
	validMethod := false
	switch req.Method {
	case model.PaymentMethodPix, model.PaymentMethodCash, model.PaymentMethodCard, model.PaymentMethodOther:
		validMethod = true
	}

	if !validMethod {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Método de pagamento inválido"})
		return
	}

	// Parse cost center ID
	costCenterID, err := uuid.Parse(req.CostCenterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cost center ID format"})
		return
	}

	// Create payment ID
	paymentID := uuid.New()

	// Create payment model
	payment := &model.Payment{
		ID:           paymentID,
		UserID:       parsedUserID,
		CostCenterID: costCenterID,
		PaymentDate:  req.PaymentDate,
		Amount:       req.Amount,
		Method:       req.Method,
		Notes:        req.Notes,
		CreatedAt:    time.Now(),
	}

	// Parse patient ID if provided
	if req.PatientID != nil {
		patientID, err := uuid.Parse(*req.PatientID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID format"})
			return
		}
		payment.PatientID = &patientID
	}

	// Create repositories
	paymentRepo := repository.NewPaymentRepository(config.DB)
	paymentAppointmentRepo := repository.NewPaymentAppointmentRepository(config.DB)

	// Save payment
	if err := paymentRepo.Save(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment", "details": err.Error()})
		return
	}

	// Link payment to appointments if provided
	if req.AppointmentIDs != nil && len(req.AppointmentIDs) > 0 {
		for _, appointmentIDStr := range req.AppointmentIDs {
			appointmentID, err := uuid.Parse(appointmentIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID format", "appointment_id": appointmentIDStr})
				return
			}

			paymentAppointment := &model.PaymentAppointment{
				ID:            uuid.New(),
				PaymentID:     paymentID,
				AppointmentID: appointmentID,
				CreatedAt:     time.Now(),
			}

			if err := paymentAppointmentRepo.Save(paymentAppointment); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link payment to appointment", "details": err.Error()})
				return
			}
		}
	}

	// Prepare response
	response := gin.H{
		"id":             payment.ID.String(),
		"user_id":        payment.UserID.String(),
		"cost_center_id": payment.CostCenterID.String(),
		"payment_date":   payment.PaymentDate,
		"amount":         payment.Amount,
		"method":         payment.Method,
		"notes":          payment.Notes,
		"created_at":     payment.CreatedAt,
	}

	if payment.PatientID != nil {
		response["patient_id"] = payment.PatientID.String()
	}

	c.JSON(http.StatusCreated, response)
}

// GetPayments returns all payments for the authenticated user
func GetPayments(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Create repository
	repo := repository.NewPaymentRepository(config.DB)

	// Get payments
	payments, err := repo.FindByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments", "details": err.Error()})
		return
	}

	// Convert to response format
	response := make([]gin.H, len(payments))
	for i, payment := range payments {
		paymentResponse := gin.H{
			"id":             payment.ID.String(),
			"user_id":        payment.UserID.String(),
			"cost_center_id": payment.CostCenterID.String(),
			"payment_date":   payment.PaymentDate,
			"amount":         payment.Amount,
			"method":         payment.Method,
			"notes":          payment.Notes,
			"created_at":     payment.CreatedAt,
		}

		if payment.PatientID != nil {
			paymentResponse["patient_id"] = payment.PatientID.String()
		}

		response[i] = paymentResponse
	}

	c.JSON(http.StatusOK, response)
}

// GetPayment returns a specific payment
func GetPayment(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Parse payment ID from URL
	paymentID := c.Param("id")

	// Create repositories
	paymentRepo := repository.NewPaymentRepository(config.DB)
	paymentAppointmentRepo := repository.NewPaymentAppointmentRepository(config.DB)

	// Get payment
	payment, err := paymentRepo.FindByID(paymentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Check if payment belongs to the authenticated user
	if payment.UserID != parsedUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this payment"})
		return
	}

	// Get linked appointments
	paymentAppointments, err := paymentAppointmentRepo.FindByPaymentID(paymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payment appointments", "details": err.Error()})
		return
	}

	// Extract appointment IDs
	appointmentIDs := make([]string, len(paymentAppointments))
	for i, pa := range paymentAppointments {
		appointmentIDs[i] = pa.AppointmentID.String()
	}

	// Prepare response
	response := gin.H{
		"id":             payment.ID.String(),
		"user_id":        payment.UserID.String(),
		"cost_center_id": payment.CostCenterID.String(),
		"payment_date":   payment.PaymentDate,
		"amount":         payment.Amount,
		"method":         payment.Method,
		"notes":          payment.Notes,
		"created_at":     payment.CreatedAt,
		"appointments":   appointmentIDs,
	}

	if payment.PatientID != nil {
		response["patient_id"] = payment.PatientID.String()
	}

	c.JSON(http.StatusOK, response)
}

// CreateRepasse handles the creation of a new repasse
func CreateRepasse(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req dto.RepasseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validate repasse status
	validStatus := false
	switch req.Status {
	case model.RepasseStatusPending, model.RepasseStatusPaid, model.RepasseStatusInformational:
		validStatus = true
	}

	if !validStatus {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status de repasse inválido"})
		return
	}

	// Parse IDs
	appointmentID, err := uuid.Parse(req.AppointmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID format"})
		return
	}

	costCenterID, err := uuid.Parse(req.CostCenterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cost center ID format"})
		return
	}

	// Create repasse ID
	repasseID := uuid.New()

	// Create repasse model
	repasse := &model.Repasse{
		ID:                repasseID,
		UserID:            parsedUserID,
		AppointmentID:     appointmentID,
		CostCenterID:      costCenterID,
		Value:             req.Value,
		DoesClinicReceive: req.ClinicReceives,
		Status:            req.Status,
		Notes:             req.Notes,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	// Set paid date if provided
	if req.PaidAt != nil {
		repasse.PaidAt = req.PaidAt
	}

	// Create repository and save repasse
	repo := repository.NewRepasseRepository(config.DB)
	if err := repo.Save(repasse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create repasse", "details": err.Error()})
		return
	}

	// Prepare response
	response := gin.H{
		"id":              repasse.ID.String(),
		"user_id":         repasse.UserID.String(),
		"appointment_id":  repasse.AppointmentID.String(),
		"cost_center_id":  repasse.CostCenterID.String(),
		"value":           repasse.Value,
		"clinic_receives": repasse.DoesClinicReceive,
		"status":          repasse.Status,
		"notes":           repasse.Notes,
		"created_at":      repasse.CreatedAt,
		"updated_at":      repasse.UpdatedAt,
	}

	if repasse.PaidAt != nil {
		response["paid_at"] = repasse.PaidAt
	}

	c.JSON(http.StatusCreated, response)
}

// GetRepasses returns all repasses for the authenticated user
func GetRepasses(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Create repository
	repo := repository.NewRepasseRepository(config.DB)

	// Get repasses
	repasses, err := repo.FindByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch repasses", "details": err.Error()})
		return
	}

	// Convert to response format
	response := make([]gin.H, len(repasses))
	for i, repasse := range repasses {
		repasseResponse := gin.H{
			"id":              repasse.ID.String(),
			"user_id":         repasse.UserID.String(),
			"appointment_id":  repasse.AppointmentID.String(),
			"cost_center_id":  repasse.CostCenterID.String(),
			"value":           repasse.Value,
			"clinic_receives": repasse.DoesClinicReceive,
			"status":          repasse.Status,
			"notes":           repasse.Notes,
			"created_at":      repasse.CreatedAt,
			"updated_at":      repasse.UpdatedAt,
		}

		if repasse.PaidAt != nil {
			repasseResponse["paid_at"] = repasse.PaidAt
		}

		response[i] = repasseResponse
	}

	c.JSON(http.StatusOK, response)
}

// UpdateRepasseStatus updates the status of a specific repasse
func UpdateRepasseStatus(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Parse repasse ID from URL
	repasseID := c.Param("id")

	var req dto.RepasseUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create repository
	repo := repository.NewRepasseRepository(config.DB)

	// Get repasse
	repasse, err := repo.FindByID(repasseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repasse not found"})
		return
	}

	// Check if repasse belongs to the authenticated user
	if repasse.UserID != parsedUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this repasse"})
		return
	}

	// Update repasse status
	if req.Status != nil {
		// Validate repasse status
		validStatus := false
		switch *req.Status {
		case model.RepasseStatusPending, model.RepasseStatusPaid, model.RepasseStatusInformational:
			validStatus = true
		}

		if !validStatus {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Status de repasse inválido"})
			return
		}

		repasse.Status = *req.Status

		// Set paid date if status is "paid"
		if *req.Status == model.RepasseStatusPaid {
			now := time.Now()
			repasse.PaidAt = &now
		}
	}

	// Save updated repasse
	if err := repo.Update(repasse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update repasse", "details": err.Error()})
		return
	}

	// Prepare response
	response := gin.H{
		"id":              repasse.ID.String(),
		"user_id":         repasse.UserID.String(),
		"appointment_id":  repasse.AppointmentID.String(),
		"cost_center_id":  repasse.CostCenterID.String(),
		"value":           repasse.Value,
		"clinic_receives": repasse.DoesClinicReceive,
		"status":          repasse.Status,
		"notes":           repasse.Notes,
		"created_at":      repasse.CreatedAt,
		"updated_at":      repasse.UpdatedAt,
	}

	if repasse.PaidAt != nil {
		response["paid_at"] = repasse.PaidAt
	}

	c.JSON(http.StatusOK, response)
}
