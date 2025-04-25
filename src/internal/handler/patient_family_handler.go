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

// CreatePatientFamily creates a new patient family member
func CreatePatientFamily(c *gin.Context) {
	var req dto.PatientFamilyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validate relationship
	validRelationship := false
	switch req.Relationship {
	case model.RelationshipFather, model.RelationshipMother, model.RelationshipSpouse,
		model.RelationshipChild, model.RelationshipResponsible, model.RelationshipGrandparent,
		model.RelationshipSibling, model.RelationshipOther:
		validRelationship = true
	}

	if !validRelationship {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de relacionamento inválido"})
		return
	}

	// Validate that the patient exists and belongs to the user
	patientID := req.PatientID
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	patientRepo := repository.NewPatientRepository(config.DB)
	_, err = patientRepo.FindByID(patientID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
		return
	}

	patientFamily := model.PatientFamily{
		ID:           uuid.New(),
		PatientID:    req.PatientID,
		Relationship: req.Relationship,
		Name:         req.Name,
		BirthDate:    req.BirthDate,
		Schooling:    req.Schooling,
		Occupation:   req.Occupation,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	patientFamilyRepo := repository.NewPatientFamilyRepository(config.DB)
	if err := patientFamilyRepo.Create(&patientFamily); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar familiar do paciente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.NewPatientFamilyResponse(patientFamily))
}

// GetPatientFamily gets a patient family member by ID
func GetPatientFamily(c *gin.Context) {
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

	patientFamilyRepo := repository.NewPatientFamilyRepository(config.DB)
	patientFamily, err := patientFamilyRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Familiar do paciente não encontrado"})
		return
	}

	// Verify that the patient belongs to the user
	patientRepo := repository.NewPatientRepository(config.DB)
	_, err = patientRepo.FindByID(patientFamily.PatientID, userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado a acessar este familiar"})
		return
	}

	c.JSON(http.StatusOK, dto.NewPatientFamilyResponse(*patientFamily))
}

// GetPatientFamilies gets all family members for a patient
func GetPatientFamilies(c *gin.Context) {
	patientID, err := uuid.Parse(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do paciente inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	// Verify that the patient belongs to the user
	patientRepo := repository.NewPatientRepository(config.DB)
	_, err = patientRepo.FindByID(patientID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
		return
	}

	patientFamilyRepo := repository.NewPatientFamilyRepository(config.DB)
	patientFamilies, err := patientFamilyRepo.FindByPatient(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar familiares do paciente", "details": err.Error()})
		return
	}

	// Convert patient families to response DTOs
	var responses []dto.PatientFamilyResponse
	for _, patientFamily := range patientFamilies {
		responses = append(responses, dto.NewPatientFamilyResponse(patientFamily))
	}

	c.JSON(http.StatusOK, responses)
}

// UpdatePatientFamily updates a patient family member
func UpdatePatientFamily(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.PatientFamilyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Validate relationship
	validRelationship := false
	switch req.Relationship {
	case model.RelationshipFather, model.RelationshipMother, model.RelationshipSpouse,
		model.RelationshipChild, model.RelationshipResponsible, model.RelationshipGrandparent,
		model.RelationshipSibling, model.RelationshipOther:
		validRelationship = true
	}

	if !validRelationship {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de relacionamento inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	patientFamilyRepo := repository.NewPatientFamilyRepository(config.DB)
	patientFamily, err := patientFamilyRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Familiar do paciente não encontrado"})
		return
	}

	// Verify that the patient belongs to the user
	patientRepo := repository.NewPatientRepository(config.DB)
	_, err = patientRepo.FindByID(patientFamily.PatientID, userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado a acessar este familiar"})
		return
	}

	// Update patient family fields
	patientFamily.PatientID = req.PatientID
	patientFamily.Relationship = req.Relationship
	patientFamily.Name = req.Name
	patientFamily.BirthDate = req.BirthDate
	patientFamily.Schooling = req.Schooling
	patientFamily.Occupation = req.Occupation
	patientFamily.UpdatedAt = time.Now()

	if err := patientFamilyRepo.Update(patientFamily); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar familiar do paciente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewPatientFamilyResponse(*patientFamily))
}

// DeletePatientFamily deletes a patient family member
func DeletePatientFamily(c *gin.Context) {
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

	patientFamilyRepo := repository.NewPatientFamilyRepository(config.DB)
	patientFamily, err := patientFamilyRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Familiar do paciente não encontrado"})
		return
	}

	// Verify that the patient belongs to the user
	patientRepo := repository.NewPatientRepository(config.DB)
	_, err = patientRepo.FindByID(patientFamily.PatientID, userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado a acessar este familiar"})
		return
	}

	if err := patientFamilyRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir familiar do paciente", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Familiar do paciente excluído com sucesso"})
}

// GetPatientFamiliesByRelationship gets family members by relationship type
func GetPatientFamiliesByRelationship(c *gin.Context) {
	patientID, err := uuid.Parse(c.Param("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do paciente inválido"})
		return
	}

	relationship := c.Param("relationship")
	if relationship == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de relacionamento não fornecido"})
		return
	}

	// Validate relationship
	validRelationship := false
	switch relationship {
	case model.RelationshipFather, model.RelationshipMother, model.RelationshipSpouse,
		model.RelationshipChild, model.RelationshipResponsible, model.RelationshipGrandparent,
		model.RelationshipSibling, model.RelationshipOther:
		validRelationship = true
	}

	if !validRelationship {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de relacionamento inválido"})
		return
	}

	// Get user ID from token
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	// Verify that the patient belongs to the user
	patientRepo := repository.NewPatientRepository(config.DB)
	_, err = patientRepo.FindByID(patientID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
		return
	}

	patientFamilyRepo := repository.NewPatientFamilyRepository(config.DB)
	patientFamilies, err := patientFamilyRepo.FindByRelationship(patientID, relationship)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar familiares do paciente", "details": err.Error()})
		return
	}

	// Convert patient families to response DTOs
	var responses []dto.PatientFamilyResponse
	for _, patientFamily := range patientFamilies {
		responses = append(responses, dto.NewPatientFamilyResponse(patientFamily))
	}

	c.JSON(http.StatusOK, responses)
}
