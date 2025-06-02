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

func TestRegisterUser(t *testing.T) {
	// Setup router
	r := gin.Default()
	r.POST("/users", RegisterUser)

	// Create a request body
	user := map[string]interface{}{
		"name":     "Jane Doe",
		"email":    "janedoe@example.com",
		"password": "password123",
	}
	jsonValue, _ := json.Marshal(user)

	// Create a request
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Jane Doe")
}
