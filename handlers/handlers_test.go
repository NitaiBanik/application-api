package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestAPIHandler(t *testing.T) {
	handler := NewHandler()

	tests := []struct {
		name           string
		body           string
		expectedStatus int
	}{
		{
			name:           "valid JSON",
			body:           `{"message": "Hello"}`,
			expectedStatus: 200,
		},
		{
			name:           "empty JSON",
			body:           `{}`,
			expectedStatus: 400,
		},
		{
			name:           "invalid JSON",
			body:           `{invalid json}`,
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/testapi", bytes.NewBuffer([]byte(tt.body)))
			w := httptest.NewRecorder()
			handler.TestAPIHandler(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestHealthHandler(t *testing.T) {
	handler := NewHandler()
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	handler.HealthHandler(w, req)

	assert.Equal(t, 200, w.Code)
}
