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

// CreateAnamneseField handles the creation of a new anamnese field
func CreateAnamneseField(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse anamnese template ID from URL
	anamneseID := c.Param("template_id")

	// Parse user ID
	userIDParsed, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req dto.AnamneseFieldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create repositories
	templateRepo := repository.NewAnamneseTemplateRepository(config.DB)
	fieldRepo := repository.NewAnamneseFieldRepository(config.DB)

	// Check if template exists and belongs to the user
	template, err := templateRepo.FindByID(anamneseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anamnese template not found"})
		return
	}

	if template.UserID != userIDParsed {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to add fields to this template"})
		return
	}

	// Create field ID
	fieldID := uuid.New()
	parsedAnamneseID, _ := uuid.Parse(anamneseID)

	// Create field model
	field := &model.AnamneseField{
		ID:            fieldID,
		FieldNumber:   req.FieldNumber,
		FieldType:     req.FieldType,
		FieldTitle:    req.FieldTitle,
		FieldRequired: req.FieldRequired,
		FieldActive:   true,
		UserID:        userIDParsed,
		AnamneseID:    parsedAnamneseID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Save field
	if err := fieldRepo.Save(field); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create field", "details": err.Error()})
		return
	}

	// Create options if provided
	var options []gin.H
	if req.Options != nil && len(req.Options) > 0 && (field.FieldType == "select" || field.FieldType == "multiselect") {
		optionRepo := repository.NewAnamneseFieldOptionRepository(config.DB)
		var fieldOptions []*model.AnamneseFieldOption

		for _, item := range req.Options {
			option := &model.AnamneseFieldOption{
				AnamneseFieldID: fieldID,
				OptionValue:     item.OptionValue,
				OptionOrder:     item.OptionOrder,
			}
			fieldOptions = append(fieldOptions, option)
		}

		if err := optionRepo.SaveBulk(fieldOptions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create field options", "details": err.Error()})
			return
		}

		// Convert options to response format
		options = make([]gin.H, len(fieldOptions))
		for i, option := range fieldOptions {
			options[i] = gin.H{
				"id":                option.ID.String(),
				"option_value":      option.OptionValue,
				"option_order":      option.OptionOrder,
				"anamnese_field_id": field.ID.String(),
			}
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":             field.ID.String(),
		"field_number":   field.FieldNumber,
		"field_type":     field.FieldType,
		"field_title":    field.FieldTitle,
		"field_required": field.FieldRequired,
		"field_active":   field.FieldActive,
		"anamnese_id":    field.AnamneseID.String(),
		"user_id":        field.UserID.String(),
		"options":        options,
	})
}

// GetAnamneseFields returns all fields for a specific anamnese template
func GetAnamneseFields(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse anamnese template ID from URL
	anamneseID := c.Param("template_id")

	// Parse user ID
	userIDParsed, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Create repositories
	templateRepo := repository.NewAnamneseTemplateRepository(config.DB)
	fieldRepo := repository.NewAnamneseFieldRepository(config.DB)
	patientAnamneseRepo := repository.NewPatientAnamneseRepository(config.DB)
	patientAnamneseFieldRepo := repository.NewPatientAnamneseFieldRepository(config.DB)

	// Check if template exists and belongs to the user
	template, err := templateRepo.FindByID(anamneseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anamnese template not found"})
		return
	}

	if template.UserID != userIDParsed {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view fields for this template"})
		return
	}

	// Get fields
	fields, err := fieldRepo.FindByAnamneseID(anamneseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch fields", "details": err.Error()})
		return
	}

	// Get all patient anamneses for the user
	patientAnamneses, err := patientAnamneseRepo.FindByUserID(userIDParsed.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patient anamneses", "details": err.Error()})
		return
	}

	// Filter patient anamneses to only include those that use the template
	var filteredAnamneses []*model.PatientAnamnese
	for _, anamnese := range patientAnamneses {
		if anamnese.AnamneseID.String() == anamneseID {
			filteredAnamneses = append(filteredAnamneses, anamnese)
		}
	}

	// Create a map to store filled fields by field ID
	filledFieldsMap := make(map[string][]gin.H)

	// For each patient anamnese, retrieve the filled fields
	for _, anamnese := range filteredAnamneses {
		filledFields, err := patientAnamneseFieldRepo.FindByPatientAnamneseID(anamnese.ID.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch filled fields", "details": err.Error()})
			return
		}

		// Group filled fields by field ID
		for _, filledField := range filledFields {
			fieldID := filledField.FieldID.String()
			filledFieldsMap[fieldID] = append(filledFieldsMap[fieldID], gin.H{
				"id":                  filledField.ID.String(),
				"patient_anamnese_id": filledField.PatientAnamneseID.String(),
				"value":               filledField.Value,
			})
		}
	}

	// Convert to response format
	response := make([]gin.H, len(fields))
	for i, field := range fields {
		// Convert options to response format
		options := make([]gin.H, len(field.Options))
		for j, option := range field.Options {
			options[j] = gin.H{
				"id":                option.ID.String(),
				"option_value":      option.OptionValue,
				"option_order":      option.OptionOrder,
				"anamnese_field_id": field.ID.String(),
			}
		}

		// Get filled fields for this field
		fieldID := field.ID.String()
		filledFields := filledFieldsMap[fieldID]

		response[i] = gin.H{
			"id":             field.ID.String(),
			"field_number":   field.FieldNumber,
			"field_type":     field.FieldType,
			"field_title":    field.FieldTitle,
			"field_required": field.FieldRequired,
			"field_active":   field.FieldActive,
			"anamnese_id":    field.AnamneseID.String(),
			"user_id":        field.UserID.String(),
			"options":        options,
			"filled_fields":  filledFields,
		}
	}

	c.JSON(http.StatusOK, response)
}

// UpdateAnamneseField updates a specific anamnese field
func UpdateAnamneseField(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse field ID from URL
	fieldID := c.Param("field_id")

	// Parse user ID
	userIDParsed, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req dto.AnamneseFieldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create repository
	fieldRepo := repository.NewAnamneseFieldRepository(config.DB)

	// Check if field exists and belongs to the user
	field, err := fieldRepo.FindByID(fieldID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}

	if field.UserID != userIDParsed {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this field"})
		return
	}

	// Update field
	field.FieldNumber = req.FieldNumber
	field.FieldType = req.FieldType
	field.FieldTitle = req.FieldTitle
	field.FieldRequired = req.FieldRequired
	field.UpdatedAt = time.Now()

	// Save updated field
	if err := fieldRepo.Update(field); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update field", "details": err.Error()})
		return
	}

	// Update options if provided
	if req.Options != nil && len(req.Options) > 0 && (field.FieldType == "select" || field.FieldType == "multiselect") {
		optionRepo := repository.NewAnamneseFieldOptionRepository(config.DB)

		// Delete existing options
		if err := optionRepo.DeleteByAnamneseFieldID(fieldID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete existing options", "details": err.Error()})
			return
		}

		// Create new options
		var fieldOptions []*model.AnamneseFieldOption
		fieldIDParsed, _ := uuid.Parse(fieldID)

		for _, item := range req.Options {
			option := &model.AnamneseFieldOption{
				AnamneseFieldID: fieldIDParsed,
				OptionValue:     item.OptionValue,
				OptionOrder:     item.OptionOrder,
			}
			fieldOptions = append(fieldOptions, option)
		}

		if err := optionRepo.SaveBulk(fieldOptions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create field options", "details": err.Error()})
			return
		}

		// Reload field to get options
		field, err = fieldRepo.FindByID(fieldID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reload field", "details": err.Error()})
			return
		}
	}

	// Convert options to response format
	options := make([]gin.H, len(field.Options))
	for i, option := range field.Options {
		options[i] = gin.H{
			"id":                option.ID.String(),
			"option_value":      option.OptionValue,
			"option_order":      option.OptionOrder,
			"anamnese_field_id": field.ID.String(),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             field.ID.String(),
		"field_number":   field.FieldNumber,
		"field_type":     field.FieldType,
		"field_title":    field.FieldTitle,
		"field_required": field.FieldRequired,
		"field_active":   field.FieldActive,
		"anamnese_id":    field.AnamneseID.String(),
		"user_id":        field.UserID.String(),
		"options":        options,
	})
}

// DeleteAnamneseField deletes a specific anamnese field
func DeleteAnamneseField(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse field ID from URL
	fieldID := c.Param("field_id")

	// Parse user ID
	userIDParsed, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Create repository
	fieldRepo := repository.NewAnamneseFieldRepository(config.DB)

	// Check if field exists and belongs to the user
	field, err := fieldRepo.FindByID(fieldID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}

	if field.UserID != userIDParsed {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this field"})
		return
	}

	// Delete field options first
	optionRepo := repository.NewAnamneseFieldOptionRepository(config.DB)
	if err := optionRepo.DeleteByAnamneseFieldID(fieldID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete field options", "details": err.Error()})
		return
	}

	// Delete field
	if err := fieldRepo.Delete(fieldID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete field", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Field deleted successfully"})
}
