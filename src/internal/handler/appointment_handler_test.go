package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAppointment(t *testing.T) {
	// Setup router
	r := gin.Default()
	r.POST("/appointments", CreateAppointment)

	// Create a request body
	appointment := map[string]interface{}{
		"patient_id":      "550e8400-e29b-41d4-a716-446655440000",
		"professional_id": "550e8400-e29b-41d4-a716-446655440000",
		"cost_center_id":  "550e8400-e29b-41d4-a716-446655440000",
		"service_title":   "Consulta",
		"start_time":      "2023-01-01T10:00:00Z",
		"end_time":        "2023-01-01T11:00:00Z",
	}
	jsonValue, _ := json.Marshal(appointment)

	// Create a request
	req, _ := http.NewRequest("POST", "/appointments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Consulta")
}
