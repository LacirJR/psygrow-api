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

// CreateAnamneseTemplate handles the creation of a new anamnese template
func CreateAnamneseTemplate(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var req dto.AnamneseTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Parse user ID
	userIDParsed, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Create template ID
	templateID := uuid.New()

	// Create template model
	template := &model.AnamneseTemplate{
		ID:        templateID,
		Title:     req.Title,
		UserID:    userIDParsed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create repository and save template
	repo := repository.NewAnamneseTemplateRepository(config.DB)
	if err := repo.Save(template); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create template", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.NewAnamneseTemplateResponse(templateID, req.Title, userIDParsed))
}

// GetAnamneseTemplates returns all anamnese templates for the authenticated user
func GetAnamneseTemplates(c *gin.Context) {
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

	// Create repository and fetch templates
	repo := repository.NewAnamneseTemplateRepository(config.DB)
	templates, err := repo.FindByUserID(clientID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch templates", "details": err.Error()})
		return
	}

	// Convert to response DTOs
	response := make([]dto.AnamneseTemplateResponse, len(templates))
	for i, t := range templates {
		response[i] = dto.NewAnamneseTemplateResponse(t.ID, t.Title, t.UserID)
	}

	c.JSON(http.StatusOK, response)
}

// GetAnamneseTemplate returns a specific anamnese template
func GetAnamneseTemplate(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse template ID from URL
	templateID := c.Param("template_id")

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Create repository and fetch template
	repo := repository.NewAnamneseTemplateRepository(config.DB)
	template, err := repo.FindByID(templateID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}

	// Check if template belongs to the authenticated user
	if template.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this template"})
		return
	}

	c.JSON(http.StatusOK, dto.NewAnamneseTemplateResponse(template.ID, template.Title, template.UserID))
}

// UpdateAnamneseTemplate updates a specific anamnese template
func UpdateAnamneseTemplate(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse template ID from URL
	templateID := c.Param("template_id")

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req dto.AnamneseTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create repository and fetch template
	repo := repository.NewAnamneseTemplateRepository(config.DB)
	template, err := repo.FindByID(templateID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}

	// Check if template belongs to the authenticated user
	if template.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this template"})
		return
	}

	// Update template
	template.Title = req.Title
	template.UpdatedAt = time.Now()

	// Save updated template
	if err := repo.Update(template); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update template", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewAnamneseTemplateResponse(template.ID, template.Title, template.UserID))
}

// DeleteAnamneseTemplate deletes a specific anamnese template
func DeleteAnamneseTemplate(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse template ID from URL
	templateID := c.Param("template_id")

	// Parse user ID
	clientID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Create repository and fetch template
	repo := repository.NewAnamneseTemplateRepository(config.DB)
	template, err := repo.FindByID(templateID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}

	// Check if template belongs to the authenticated user
	if template.UserID != clientID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this template"})
		return
	}

	// Delete template
	if err := repo.Delete(templateID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete template", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}
