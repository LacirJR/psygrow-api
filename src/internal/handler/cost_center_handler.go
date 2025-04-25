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

// CreateCostCenter handles the creation of a new cost center
func CreateCostCenter(c *gin.Context) {
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

	var req dto.CostCenterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create cost center ID
	costCenterID := uuid.New()

	// Create cost center model
	costCenter := &model.CostCenter{
		ID:           costCenterID,
		UserID:       parsedUserID,
		Name:         req.Name,
		RepasseModel: req.RepasseModel,
		RepasseType:  req.RepasseType,
		RepasseValue: req.RepasseValue,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Create repository and save cost center
	repo := repository.NewCostCenterRepository(config.DB)
	if err := repo.Save(costCenter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cost center", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":            costCenter.ID.String(),
		"user_id":       costCenter.UserID.String(),
		"name":          costCenter.Name,
		"repasse_model": costCenter.RepasseModel,
		"repasse_type":  costCenter.RepasseType,
		"repasse_value": costCenter.RepasseValue,
		"active":        costCenter.IsActive,
		"created_at":    costCenter.CreatedAt,
		"updated_at":    costCenter.UpdatedAt,
	})
}

// GetCostCenters returns all cost centers for the authenticated user
func GetCostCenters(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Create repository
	repo := repository.NewCostCenterRepository(config.DB)

	// Get cost centers
	costCenters, err := repo.FindByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cost centers", "details": err.Error()})
		return
	}

	// Convert to response format
	response := make([]gin.H, len(costCenters))
	for i, costCenter := range costCenters {
		response[i] = gin.H{
			"id":            costCenter.ID.String(),
			"user_id":       costCenter.UserID.String(),
			"name":          costCenter.Name,
			"repasse_model": costCenter.RepasseModel,
			"repasse_type":  costCenter.RepasseType,
			"repasse_value": costCenter.RepasseValue,
			"active":        costCenter.IsActive,
			"created_at":    costCenter.CreatedAt,
			"updated_at":    costCenter.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetCostCenter returns a specific cost center
func GetCostCenter(c *gin.Context) {
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

	// Parse cost center ID from URL
	costCenterID := c.Param("id")

	// Create repository
	repo := repository.NewCostCenterRepository(config.DB)

	// Get cost center
	costCenter, err := repo.FindByID(costCenterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cost center not found"})
		return
	}

	// Check if cost center belongs to the authenticated user
	if costCenter.UserID != parsedUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this cost center"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            costCenter.ID.String(),
		"user_id":       costCenter.UserID.String(),
		"name":          costCenter.Name,
		"repasse_model": costCenter.RepasseModel,
		"repasse_type":  costCenter.RepasseType,
		"repasse_value": costCenter.RepasseValue,
		"active":        costCenter.IsActive,
		"created_at":    costCenter.CreatedAt,
		"updated_at":    costCenter.UpdatedAt,
	})
}

// UpdateCostCenter updates a specific cost center
func UpdateCostCenter(c *gin.Context) {
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

	// Parse cost center ID from URL
	costCenterID := c.Param("id")

	var req dto.CostCenterUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Create repository
	repo := repository.NewCostCenterRepository(config.DB)

	// Get cost center
	costCenter, err := repo.FindByID(costCenterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cost center not found"})
		return
	}

	// Check if cost center belongs to the authenticated user
	if costCenter.UserID != parsedUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this cost center"})
		return
	}

	// Update cost center fields if provided
	if req.Name != nil {
		costCenter.Name = *req.Name
	}

	if req.RepasseModel != nil {
		costCenter.RepasseModel = *req.RepasseModel
	}

	if req.RepasseType != nil {
		costCenter.RepasseType = *req.RepasseType
	}

	if req.RepasseValue != nil {
		costCenter.RepasseValue = *req.RepasseValue
	}

	if req.Active != nil {
		costCenter.IsActive = *req.Active
	}

	costCenter.UpdatedAt = time.Now()

	// Save updated cost center
	if err := repo.Update(costCenter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cost center", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            costCenter.ID.String(),
		"user_id":       costCenter.UserID.String(),
		"name":          costCenter.Name,
		"repasse_model": costCenter.RepasseModel,
		"repasse_type":  costCenter.RepasseType,
		"repasse_value": costCenter.RepasseValue,
		"active":        costCenter.IsActive,
		"created_at":    costCenter.CreatedAt,
		"updated_at":    costCenter.UpdatedAt,
	})
}

// DeleteCostCenter deletes a specific cost center
func DeleteCostCenter(c *gin.Context) {
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

	// Parse cost center ID from URL
	costCenterID := c.Param("id")

	// Create repository
	repo := repository.NewCostCenterRepository(config.DB)

	// Get cost center
	costCenter, err := repo.FindByID(costCenterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cost center not found"})
		return
	}

	// Check if cost center belongs to the authenticated user
	if costCenter.UserID != parsedUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this cost center"})
		return
	}

	// Delete cost center
	if err := repo.Delete(costCenterID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cost center", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cost center deleted successfully"})
}
