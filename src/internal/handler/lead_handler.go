package handler

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	dto "github.com/LacirJR/psygrow-api/src/internal/core/dto/lead"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/infra/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

// CreateLead creates a new lead
func CreateLead(c *gin.Context) {
	var req dto.LeadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validate lead status
	validStatus := false
	switch req.Status {
	case model.LeadStatusNew, model.LeadStatusInAnalysis, model.LeadStatusConverted, model.LeadStatusLost:
		validStatus = true
	}

	if !validStatus {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status de lead inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	lead := model.Lead{
		ID:               uuid.New(),
		UserID:           userID,
		FullName:         req.FullName,
		Phone:            req.Phone,
		Email:            req.Email,
		BirthDate:        req.BirthDate,
		ContactDate:      req.ContactDate,
		Status:           req.Status,
		WasAttended:      req.WasAttended,
		Notes:            req.Notes,
		Origin:           req.Origin,
		GdprBlockContact: req.GdprBlockContact,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	leadRepo := repository.NewLeadRepository(config.DB)
	if err := leadRepo.Create(&lead); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar lead", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.NewLeadResponse(lead))
}

// GetLead gets a lead by ID
func GetLead(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	leadRepo := repository.NewLeadRepository(config.DB)
	lead, err := leadRepo.FindByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead não encontrado"})
		return
	}

	c.JSON(http.StatusOK, dto.NewLeadResponse(*lead))
}

// GetLeads gets all leads for a user
func GetLeads(c *gin.Context) {
	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	// Parse pagination parameters
	limit, offset := getPaginationParams(c)

	leadRepo := repository.NewLeadRepository(config.DB)
	leads, err := leadRepo.FindAll(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar leads", "details": err.Error()})
		return
	}

	// Convert leads to response DTOs
	var responses []dto.LeadResponse
	for _, lead := range leads {
		responses = append(responses, dto.NewLeadResponse(lead))
	}

	// Get total count
	count, err := leadRepo.Count(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar leads", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   responses,
		"total":  count,
		"limit":  limit,
		"offset": offset,
	})
}

// UpdateLead updates a lead
func UpdateLead(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.LeadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validate lead status
	validStatus := false
	switch req.Status {
	case model.LeadStatusNew, model.LeadStatusInAnalysis, model.LeadStatusConverted, model.LeadStatusLost:
		validStatus = true
	}

	if !validStatus {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status de lead inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	leadRepo := repository.NewLeadRepository(config.DB)
	lead, err := leadRepo.FindByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead não encontrado"})
		return
	}

	// Update lead fields
	lead.FullName = req.FullName
	lead.Phone = req.Phone
	lead.Email = req.Email
	lead.BirthDate = req.BirthDate
	lead.ContactDate = req.ContactDate
	lead.Status = req.Status
	lead.WasAttended = req.WasAttended
	lead.Notes = req.Notes
	lead.Origin = req.Origin
	lead.GdprBlockContact = req.GdprBlockContact
	lead.UpdatedAt = time.Now()

	if err := leadRepo.Update(lead); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar lead", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewLeadResponse(*lead))
}

// DeleteLead deletes a lead
func DeleteLead(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	leadRepo := repository.NewLeadRepository(config.DB)
	if err := leadRepo.Delete(id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir lead", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lead excluído com sucesso"})
}

// ConvertLeadToPatient converts a lead to a patient
func ConvertLeadToPatient(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.ConvertLeadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	leadRepo := repository.NewLeadRepository(config.DB)
	patient, err := leadRepo.ConvertToPatient(id, userID, req.CostCenterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter lead para paciente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Lead convertido para paciente com sucesso",
		"patient_id": patient.ID,
	})
}

// Helper function to get pagination parameters
func getPaginationParams(c *gin.Context) (int, int) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	return limit, offset
}

// Helper function to get user ID from token
func getUserIDFromToken(c *gin.Context) (uuid.UUID, error) {
	userIDStr, exists := c.Get("userID")
	if !exists {
		return uuid.Nil, ErrUserIDNotFound
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

// ErrUserIDNotFound is returned when the user ID is not found in the token
var ErrUserIDNotFound = &customError{"User ID not found in token"}

type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}
