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

// GetSessions returns all sessions for the authenticated user
func GetSessions(c *gin.Context) {
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
	repo := repository.NewSessionRepository(config.DB)

	// Get sessions
	sessions, err := repo.FindByUserID(clientID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sessions", "details": err.Error()})
		return
	}

	// Convert to response format
	response := make([]gin.H, len(sessions))
	for i, session := range sessions {
		response[i] = gin.H{
			"id":              session.ID.String(),
			"appointment_id":  session.AppointmentID.String(),
			"client_id":       session.UserID.String(),
			"patient_id":      session.PatientID.String(),
			"professional_id": session.ProfessionalID.String(),
			"start_time":      session.StartTime,
			"end_time":        session.EndTime,
			"was_attended":    session.WasAttended,
			"created_at":      session.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetSession returns a specific session
func GetSession(c *gin.Context) {
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

	// Parse session ID from URL
	sessionID := c.Param("id")

	// Create repository
	repo := repository.NewSessionRepository(config.DB)

	// Get session
	session, err := repo.FindByID(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Check if session belongs to the authenticated user
	if session.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              session.ID.String(),
		"appointment_id":  session.AppointmentID.String(),
		"client_id":       session.UserID.String(),
		"patient_id":      session.PatientID.String(),
		"professional_id": session.ProfessionalID.String(),
		"start_time":      session.StartTime,
		"end_time":        session.EndTime,
		"was_attended":    session.WasAttended,
		"created_at":      session.CreatedAt,
	})
}

// CreateEvolution handles the creation of a new evolution for a session
func CreateEvolution(c *gin.Context) {
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

	// Parse session ID from URL
	sessionID := c.Param("session_id")
	parsedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID format"})
		return
	}

	var req dto.EvolutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create repositories
	sessionRepo := repository.NewSessionRepository(config.DB)
	evolutionRepo := repository.NewEvolutionRepository(config.DB)

	// Get session
	session, err := sessionRepo.FindByID(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Check if session belongs to the authenticated user
	if session.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to add evolution to this session"})
		return
	}

	// Create evolution ID
	evolutionID := uuid.New()

	// Create evolution model
	evolution := &model.Evolution{
		ID:             evolutionID,
		SessionID:      parsedSessionID,
		UserID:         clientID,
		ProfessionalID: session.ProfessionalID,
		PatientID:      session.PatientID,
		Content:        req.Content,
		CreatedAt:      time.Now(),
	}

	// Save evolution
	if err := evolutionRepo.Save(evolution); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create evolution", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":              evolution.ID.String(),
		"session_id":      evolution.SessionID.String(),
		"client_id":       evolution.UserID.String(),
		"patient_id":      evolution.PatientID.String(),
		"professional_id": evolution.ProfessionalID.String(),
		"content":         evolution.Content,
		"created_at":      evolution.CreatedAt,
	})
}

// GetEvolution returns a specific evolution
func GetEvolution(c *gin.Context) {
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

	// Parse evolution ID from URL
	evolutionID := c.Param("id")

	// Create repository
	repo := repository.NewEvolutionRepository(config.DB)

	// Get evolution
	evolution, err := repo.FindByID(evolutionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evolution not found"})
		return
	}

	// Check if evolution belongs to the authenticated user
	if evolution.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this evolution"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              evolution.ID.String(),
		"session_id":      evolution.SessionID.String(),
		"client_id":       evolution.UserID.String(),
		"patient_id":      evolution.PatientID.String(),
		"professional_id": evolution.ProfessionalID.String(),
		"content":         evolution.Content,
		"created_at":      evolution.CreatedAt,
	})
}

// GetEvolutionsByPatient returns all evolutions for a specific patient
func GetEvolutionsByPatient(c *gin.Context) {
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
	repo := repository.NewEvolutionRepository(config.DB)

	// Get evolutions
	evolutions, err := repo.FindByPatientID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch evolutions", "details": err.Error()})
		return
	}

	// Filter evolutions that belong to the authenticated user
	var filteredEvolutions []*model.Evolution
	for _, evolution := range evolutions {
		if evolution.UserID == clientID {
			filteredEvolutions = append(filteredEvolutions, evolution)
		}
	}

	// Convert to response format
	response := make([]gin.H, len(filteredEvolutions))
	for i, evolution := range filteredEvolutions {
		response[i] = gin.H{
			"id":              evolution.ID.String(),
			"session_id":      evolution.SessionID.String(),
			"client_id":       evolution.UserID.String(),
			"patient_id":      evolution.PatientID.String(),
			"professional_id": evolution.ProfessionalID.String(),
			"content":         evolution.Content,
			"created_at":      evolution.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, response)
}
