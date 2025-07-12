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

func TestCreatePatient(t *testing.T) {
	// Setup router
	r := gin.Default()
	r.POST("/patients", CreatePatient)

	// Create a request body
	patient := map[string]interface{}{
		"full_name":  "John Doe",
		"birth_date": "1990-01-01",
		"email":      "johndoe@example.com",
	}
	jsonValue, _ := json.Marshal(patient)

	// Create a request
	req, _ := http.NewRequest("POST", "/patients", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}
