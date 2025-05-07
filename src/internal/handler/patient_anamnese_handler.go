package handler

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	dto "github.com/LacirJR/psygrow-api/src/internal/core/dto/anamnese"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/infra/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// CreatePatientAnamnese handles the creation of a new patient anamnese
func CreatePatientAnamnese(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse user ID
	userIDParsed, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req dto.PatientAnamneseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Parse patient ID and anamnese ID
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID format"})
		return
	}

	anamneseID, err := uuid.Parse(req.AnamneseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anamnese ID format"})
		return
	}

	// Create repositories
	templateRepo := repository.NewAnamneseTemplateRepository(config.DB)
	patientAnamneseRepo := repository.NewPatientAnamneseRepository(config.DB)
	patientAnamneseFieldRepo := repository.NewPatientAnamneseFieldRepository(config.DB)

	// Check if template exists and belongs to the user
	template, err := templateRepo.FindByID(req.AnamneseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anamnese template not found"})
		return
	}

	if template.UserID != userIDParsed {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to use this template"})
		return
	}

	// Create patient anamnese ID
	patientAnamneseID := uuid.New()

	// Create patient anamnese model
	patientAnamnese := &model.PatientAnamnese{
		ID:         patientAnamneseID,
		PatientID:  patientID,
		AnamneseID: anamneseID,
		UserID:     userIDParsed,
		AnsweredAt: time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save patient anamnese
	if err := patientAnamneseRepo.Save(patientAnamnese); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient anamnese", "details": err.Error()})
		return
	}

	// Save patient anamnese fields
	for _, field := range req.Fields {
		fieldID, err := uuid.Parse(field.FieldID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field ID format", "field_id": field.FieldID})
			return
		}

		patientAnamneseField := &model.PatientAnamneseField{
			ID:                uuid.New(),
			PatientAnamneseID: patientAnamneseID,
			FieldID:           fieldID,
			Value:             field.Value,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		if err := patientAnamneseFieldRepo.Save(patientAnamneseField); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient anamnese field", "details": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          patientAnamneseID.String(),
		"patient_id":  patientID.String(),
		"anamnese_id": anamneseID.String(),
		"user_id":     userIDParsed.String(),
		"answered_at": patientAnamnese.AnsweredAt,
	})
}

// GetPatientAnamneses returns all anamneses for a specific patient
func GetPatientAnamneses(c *gin.Context) {
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

	// Parse patient ID from URL
	patientID := c.Param("patient_id")

	// Create repository
	patientAnamneseRepo := repository.NewPatientAnamneseRepository(config.DB)

	// Get patient anamneses
	patientAnamneses, err := patientAnamneseRepo.FindByPatientID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patient anamneses", "details": err.Error()})
		return
	}

	// Filter anamneses that belong to the authenticated user
	var filteredAnamneses []*model.PatientAnamnese
	for _, anamnese := range patientAnamneses {
		if anamnese.UserID == clientID {
			filteredAnamneses = append(filteredAnamneses, anamnese)
		}
	}

	// Convert to response format
	response := make([]gin.H, len(filteredAnamneses))
	for i, anamnese := range filteredAnamneses {
		response[i] = gin.H{
			"id":          anamnese.ID.String(),
			"patient_id":  anamnese.PatientID.String(),
			"anamnese_id": anamnese.AnamneseID.String(),
			"client_id":   anamnese.UserID.String(),
			"answered_at": anamnese.AnsweredAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetPatientAnamneseDetails returns details of a specific patient anamnese including fields
func GetPatientAnamneseDetails(c *gin.Context) {
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

	// Parse patient anamnese ID from URL
	patientAnamneseID := c.Param("id")

	// Create repositories
	patientAnamneseRepo := repository.NewPatientAnamneseRepository(config.DB)
	patientAnamneseFieldRepo := repository.NewPatientAnamneseFieldRepository(config.DB)

	// Get patient anamnese
	patientAnamnese, err := patientAnamneseRepo.FindByID(patientAnamneseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient anamnese not found"})
		return
	}

	// Check if anamnese belongs to the authenticated user
	if patientAnamnese.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this anamnese"})
		return
	}

	// Get patient anamnese fields
	fields, err := patientAnamneseFieldRepo.FindByPatientAnamneseID(patientAnamneseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patient anamnese fields", "details": err.Error()})
		return
	}

	// Convert fields to response format
	fieldResponses := make([]gin.H, len(fields))
	for i, field := range fields {
		fieldResponses[i] = gin.H{
			"id":                  field.ID.String(),
			"patient_anamnese_id": field.PatientAnamneseID.String(),
			"field_id":            field.FieldID.String(),
			"value":               field.Value,
		}
	}

	// Create response
	response := gin.H{
		"id":          patientAnamnese.ID.String(),
		"patient_id":  patientAnamnese.PatientID.String(),
		"anamnese_id": patientAnamnese.AnamneseID.String(),
		"client_id":   patientAnamnese.UserID.String(),
		"answered_at": patientAnamnese.AnsweredAt,
		"fields":      fieldResponses,
	}

	c.JSON(http.StatusOK, response)
}
