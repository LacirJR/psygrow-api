package handler

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	dto "github.com/LacirJR/psygrow-api/src/internal/core/dto/appointment"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/infra/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// CreateAppointment handles the creation of a new appointment
func CreateAppointment(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req dto.AppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Parse IDs
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID format"})
		return
	}

	professionalID, err := uuid.Parse(req.ProfessionalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid professional ID format"})
		return
	}

	costCenterID, err := uuid.Parse(req.CostCenterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cost center ID format"})
		return
	}

	// Use start and end times directly from request
	startTime := req.StartTime
	endTime := req.EndTime

	// Create appointment ID
	appointmentID := uuid.New()

	// Create appointment model
	appointment := &model.Appointment{
		ID:             appointmentID,
		UserID:         clientID,
		PatientID:      patientID,
		ProfessionalID: professionalID,
		CostCenterID:   costCenterID,
		ServiceTitle:   req.ServiceTitle,
		StartTime:      startTime,
		EndTime:        endTime,
		Status:         "scheduled",
		Notes:          req.Notes,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Set custom repasse values if provided
	if req.CustomRepasseType != nil {
		appointment.CustomRepasseType = req.CustomRepasseType
	}

	if req.CustomRepasseValue != nil {
		appointment.CustomRepasseValue = req.CustomRepasseValue
	}

	// Create repository and save appointment
	repo := repository.NewAppointmentRepository(config.DB)
	if err := repo.Save(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                   appointment.ID.String(),
		"client_id":            appointment.UserID.String(),
		"patient_id":           appointment.PatientID.String(),
		"professional_id":      appointment.ProfessionalID.String(),
		"cost_center_id":       appointment.CostCenterID.String(),
		"custom_repasse_type":  appointment.CustomRepasseType,
		"custom_repasse_value": appointment.CustomRepasseValue,
		"service_title":        appointment.ServiceTitle,
		"start_time":           appointment.StartTime,
		"end_time":             appointment.EndTime,
		"status":               appointment.Status,
		"notes":                appointment.Notes,
	})
}

// GetAppointments returns all appointments for the authenticated user
func GetAppointments(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Create repository
	repo := repository.NewAppointmentRepository(config.DB)

	// Get appointments
	appointments, err := repo.FindByUserID(clientID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments", "details": err.Error()})
		return
	}

	// Convert to response format
	response := make([]gin.H, len(appointments))
	for i, appointment := range appointments {
		response[i] = gin.H{
			"id":                   appointment.ID.String(),
			"client_id":            appointment.UserID.String(),
			"patient_id":           appointment.PatientID.String(),
			"professional_id":      appointment.ProfessionalID.String(),
			"cost_center_id":       appointment.CostCenterID.String(),
			"custom_repasse_type":  appointment.CustomRepasseType,
			"custom_repasse_value": appointment.CustomRepasseValue,
			"service_title":        appointment.ServiceTitle,
			"start_time":           appointment.StartTime,
			"end_time":             appointment.EndTime,
			"status":               appointment.Status,
			"notes":                appointment.Notes,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetAppointment returns a specific appointment
func GetAppointment(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Parse appointment ID from URL
	appointmentID := c.Param("id")

	// Create repository
	repo := repository.NewAppointmentRepository(config.DB)

	// Get appointment
	appointment, err := repo.FindByID(appointmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Check if appointment belongs to the authenticated user
	if appointment.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                   appointment.ID.String(),
		"client_id":            appointment.UserID.String(),
		"patient_id":           appointment.PatientID.String(),
		"professional_id":      appointment.ProfessionalID.String(),
		"cost_center_id":       appointment.CostCenterID.String(),
		"custom_repasse_type":  appointment.CustomRepasseType,
		"custom_repasse_value": appointment.CustomRepasseValue,
		"service_title":        appointment.ServiceTitle,
		"start_time":           appointment.StartTime,
		"end_time":             appointment.EndTime,
		"status":               appointment.Status,
		"notes":                appointment.Notes,
	})
}

// UpdateAppointment updates a specific appointment
func UpdateAppointment(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Parse appointment ID from URL
	appointmentID := c.Param("id")

	var req dto.AppointmentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create repository
	repo := repository.NewAppointmentRepository(config.DB)

	// Get appointment
	appointment, err := repo.FindByID(appointmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Check if appointment belongs to the authenticated user
	if appointment.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this appointment"})
		return
	}

	// Update appointment fields if provided
	if req.ServiceTitle != nil {
		appointment.ServiceTitle = *req.ServiceTitle
	}

	if req.StartTime != nil {
		appointment.StartTime = *req.StartTime
	}

	if req.EndTime != nil {
		appointment.EndTime = *req.EndTime
	}

	if req.Status != nil {
		appointment.Status = *req.Status
	}

	if req.Notes != nil {
		appointment.Notes = *req.Notes
	}

	if req.CustomRepasseType != nil {
		appointment.CustomRepasseType = req.CustomRepasseType
	}

	if req.CustomRepasseValue != nil {
		appointment.CustomRepasseValue = req.CustomRepasseValue
	}

	appointment.UpdatedAt = time.Now()

	// Save updated appointment
	if err := repo.Update(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment", "details": err.Error()})
		return
	}

	// If status is changed to "done", create a session
	if req.Status != nil && *req.Status == "done" {
		sessionRepo := repository.NewSessionRepository(config.DB)

		// Create session ID
		sessionID := uuid.New()

		// Create session model
		session := &model.Session{
			ID:             sessionID,
			AppointmentID:  appointment.ID,
			UserID:         appointment.UserID,
			PatientID:      appointment.PatientID,
			ProfessionalID: appointment.ProfessionalID,
			StartTime:      appointment.StartTime,
			EndTime:        appointment.EndTime,
			WasAttended:    true,
			CreatedAt:      time.Now(),
		}

		// Save session
		if err := sessionRepo.Save(session); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session", "details": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                   appointment.ID.String(),
		"client_id":            appointment.UserID.String(),
		"patient_id":           appointment.PatientID.String(),
		"professional_id":      appointment.ProfessionalID.String(),
		"cost_center_id":       appointment.CostCenterID.String(),
		"custom_repasse_type":  appointment.CustomRepasseType,
		"custom_repasse_value": appointment.CustomRepasseValue,
		"service_title":        appointment.ServiceTitle,
		"start_time":           appointment.StartTime,
		"end_time":             appointment.EndTime,
		"status":               appointment.Status,
		"notes":                appointment.Notes,
	})
}

// DeleteAppointment deletes a specific appointment
func DeleteAppointment(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Parse appointment ID from URL
	appointmentID := c.Param("id")

	// Create repository
	repo := repository.NewAppointmentRepository(config.DB)

	// Get appointment
	appointment, err := repo.FindByID(appointmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Check if appointment belongs to the authenticated user
	if appointment.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this appointment"})
		return
	}

	// Delete appointment
	if err := repo.Delete(appointmentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
