package handler

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	dto "github.com/LacirJR/psygrow-api/src/internal/core/dto/patient"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/infra/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// CreatePatient creates a new patient
func CreatePatient(c *gin.Context) {
	var req dto.PatientRequest
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

	patient := model.Patient{
		ID:                    uuid.New(),
		UserID:                userID,
		CostCenterID:          req.CostCenterID,
		FullName:              req.FullName,
		SocialName:            req.SocialName,
		BirthDate:             req.BirthDate,
		Document:              req.Document,
		Phone:                 req.Phone,
		Email:                 req.Email,
		Gender:                req.Gender,
		Address:               req.Address,
		ResidesWith:           req.ResidesWith,
		EmergencyContactName:  req.EmergencyContactName,
		EmergencyContactPhone: req.EmergencyContactPhone,
		Observation:           req.Observation,
		DefaultRepasseType:    req.DefaultRepasseType,
		DefaultRepasseValue:   req.DefaultRepasseValue,
		IsActive:              req.IsActive,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	patientRepo := repository.NewPatientRepository(config.DB)
	if err := patientRepo.Create(&patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar paciente", "details": err.Error()})
		return
	}

	// Fetch the patient with the cost center to get the cost center name
	createdPatient, err := patientRepo.FindByID(patient.ID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar paciente criado", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.NewPatientResponse(*createdPatient))
}

// GetPatient gets a patient by ID
func GetPatient(c *gin.Context) {
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

	patientRepo := repository.NewPatientRepository(config.DB)
	patient, err := patientRepo.FindByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, dto.NewPatientResponse(*patient))
}

// GetPatients gets all patients for a user
func GetPatients(c *gin.Context) {
	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	// Parse pagination parameters
	limit, offset := getPaginationParams(c)

	patientRepo := repository.NewPatientRepository(config.DB)
	patients, err := patientRepo.FindAll(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pacientes", "details": err.Error()})
		return
	}

	// Convert patients to response DTOs
	var responses []dto.PatientResponse
	for _, patient := range patients {
		responses = append(responses, dto.NewPatientResponse(patient))
	}

	// Get total count
	count, err := patientRepo.Count(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao contar pacientes", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   responses,
		"total":  count,
		"limit":  limit,
		"offset": offset,
	})
}

// UpdatePatient updates a patient
func UpdatePatient(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.PatientRequest
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

	patientRepo := repository.NewPatientRepository(config.DB)
	patient, err := patientRepo.FindByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
		return
	}

	// Update patient fields
	patient.CostCenterID = req.CostCenterID
	patient.FullName = req.FullName
	patient.SocialName = req.SocialName
	patient.BirthDate = req.BirthDate
	patient.Document = req.Document
	patient.Phone = req.Phone
	patient.Email = req.Email
	patient.Gender = req.Gender
	patient.Address = req.Address
	patient.ResidesWith = req.ResidesWith
	patient.EmergencyContactName = req.EmergencyContactName
	patient.EmergencyContactPhone = req.EmergencyContactPhone
	patient.Observation = req.Observation
	patient.DefaultRepasseType = req.DefaultRepasseType
	patient.DefaultRepasseValue = req.DefaultRepasseValue
	patient.IsActive = req.IsActive
	patient.UpdatedAt = time.Now()

	if err := patientRepo.Update(patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar paciente", "details": err.Error()})
		return
	}

	// Fetch the updated patient with the cost center to get the cost center name
	updatedPatient, err := patientRepo.FindByID(id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar paciente atualizado", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewPatientResponse(*updatedPatient))
}

// DeletePatient deletes a patient
func DeletePatient(c *gin.Context) {
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

	patientRepo := repository.NewPatientRepository(config.DB)
	if err := patientRepo.Delete(id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir paciente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Paciente excluído com sucesso"})
}

// SearchPatientsByName searches patients by name
func SearchPatientsByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome não fornecido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	// Parse pagination parameters
	limit, offset := getPaginationParams(c)

	patientRepo := repository.NewPatientRepository(config.DB)
	patients, err := patientRepo.FindByName(userID, name, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pacientes", "details": err.Error()})
		return
	}

	// Convert patients to response DTOs
	var responses []dto.PatientResponse
	for _, patient := range patients {
		responses = append(responses, dto.NewPatientResponse(patient))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   responses,
		"limit":  limit,
		"offset": offset,
	})
}

// GetPatientsByCostCenter gets patients by cost center
func GetPatientsByCostCenter(c *gin.Context) {
	costCenterID, err := uuid.Parse(c.Param("cost_center_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do centro de custo inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	// Parse pagination parameters
	limit, offset := getPaginationParams(c)

	patientRepo := repository.NewPatientRepository(config.DB)
	patients, err := patientRepo.FindByCostCenter(userID, costCenterID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pacientes", "details": err.Error()})
		return
	}

	// Convert patients to response DTOs
	var responses []dto.PatientResponse
	for _, patient := range patients {
		responses = append(responses, dto.NewPatientResponse(patient))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   responses,
		"limit":  limit,
		"offset": offset,
	})
}
