package handler

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	dto "github.com/LacirJR/psygrow-api/src/internal/core/dto/anamnese"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/infra/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateAnamneseFieldOption creates a new anamnese field option
func CreateAnamneseFieldOption(c *gin.Context) {
	var request dto.AnamneseFieldOptionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Parse anamnese field ID
	anamneseFieldID, err := uuid.Parse(request.AnamneseFieldID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anamnese field ID"})
		return
	}

	// Check if anamnese field exists and belongs to the user
	anamneseFieldRepo := repository.NewAnamneseFieldRepository(config.DB)
	anamneseField, err := anamneseFieldRepo.FindByID(request.AnamneseFieldID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anamnese field not found"})
		return
	}

	if anamneseField.UserID.String() != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to add options to this anamnese field"})
		return
	}

	// Create anamnese field option
	option := &model.AnamneseFieldOption{
		AnamneseFieldID: anamneseFieldID,
		OptionValue:     request.OptionValue,
		OptionOrder:     request.OptionOrder,
	}

	// Save anamnese field option
	optionRepo := repository.NewAnamneseFieldOptionRepository(config.DB)
	if err := optionRepo.Save(option); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create anamnese field option"})
		return
	}

	// Return response
	response := dto.NewAnamneseFieldOptionResponse(
		option.ID,
		option.OptionValue,
		option.OptionOrder,
		option.AnamneseFieldID,
	)

	c.JSON(http.StatusCreated, response)
}

// CreateAnamneseFieldOptionsBulk creates multiple anamnese field options at once
func CreateAnamneseFieldOptionsBulk(c *gin.Context) {
	var request dto.AnamneseFieldOptionBulkRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Parse anamnese field ID
	anamneseFieldID, err := uuid.Parse(request.AnamneseFieldID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anamnese field ID"})
		return
	}

	// Check if anamnese field exists and belongs to the user
	anamneseFieldRepo := repository.NewAnamneseFieldRepository(config.DB)
	anamneseField, err := anamneseFieldRepo.FindByID(request.AnamneseFieldID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anamnese field not found"})
		return
	}

	if anamneseField.UserID.String() != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to add options to this anamnese field"})
		return
	}

	// Delete existing options for this field
	optionRepo := repository.NewAnamneseFieldOptionRepository(config.DB)
	if err := optionRepo.DeleteByAnamneseFieldID(request.AnamneseFieldID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete existing options"})
		return
	}

	// Create anamnese field options
	var options []*model.AnamneseFieldOption
	for _, item := range request.Options {
		option := &model.AnamneseFieldOption{
			AnamneseFieldID: anamneseFieldID,
			OptionValue:     item.OptionValue,
			OptionOrder:     item.OptionOrder,
		}
		options = append(options, option)
	}

	// Save anamnese field options
	if err := optionRepo.SaveBulk(options); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create anamnese field options"})
		return
	}

	// Return response
	var responses []dto.AnamneseFieldOptionResponse
	for _, option := range options {
		response := dto.NewAnamneseFieldOptionResponse(
			option.ID,
			option.OptionValue,
			option.OptionOrder,
			option.AnamneseFieldID,
		)
		responses = append(responses, response)
	}

	c.JSON(http.StatusCreated, gin.H{"options": responses})
}

// GetAnamneseFieldOptions gets all options for an anamnese field
func GetAnamneseFieldOptions(c *gin.Context) {
	fieldID := c.Param("field_id")

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if anamnese field exists and belongs to the user
	anamneseFieldRepo := repository.NewAnamneseFieldRepository(config.DB)
	anamneseField, err := anamneseFieldRepo.FindByID(fieldID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anamnese field not found"})
		return
	}

	if anamneseField.UserID.String() != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view options for this anamnese field"})
		return
	}

	// Get options for anamnese field
	optionRepo := repository.NewAnamneseFieldOptionRepository(config.DB)
	options, err := optionRepo.FindByAnamneseFieldID(fieldID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get anamnese field options"})
		return
	}

	// Return response
	var responses []dto.AnamneseFieldOptionResponse
	for _, option := range options {
		response := dto.NewAnamneseFieldOptionResponse(
			option.ID,
			option.OptionValue,
			option.OptionOrder,
			option.AnamneseFieldID,
		)
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, gin.H{"options": responses})
}
