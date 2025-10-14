package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestApiHandler(t *testing.T) {
	handler := NewHandler()

	tests := []struct {
		name           string
		method         string
		body           interface{}
		expectedStatus int
		expectedData   interface{}
	}{
		{
			name:           "should process valid JSON and return 200",
			method:         "POST",
			body:           map[string]interface{}{"message": "Hello", "test": true},
			expectedStatus: 200,
			expectedData:   map[string]interface{}{"message": "Hello", "test": true},
		},
		{
			name:           "should reject empty JSON with 400 error",
			method:         "POST",
			body:           map[string]interface{}{},
			expectedStatus: 400,
			expectedData:   nil,
		},
		{
			name:           "should handle complex nested JSON structures",
			method:         "POST",
			body:           map[string]interface{}{"user": map[string]interface{}{"id": 123, "name": "Test"}, "items": []int{1, 2, 3}},
			expectedStatus: 200,
			expectedData:   map[string]interface{}{"user": map[string]interface{}{"id": 123, "name": "Test"}, "items": []int{1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			if tt.body != nil {
				body, _ = json.Marshal(tt.body)
			}

			req, _ := http.NewRequest(tt.method, "/testapi", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			handler.TestApiHandler(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedData != nil && w.Code == 200 {
				var response TestApiResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				expectedJSON, _ := json.Marshal(tt.expectedData)
				actualJSON, _ := json.Marshal(response.Data)
				assert.Equal(t, string(expectedJSON), string(actualJSON))

				assert.NotNil(t, response.Timestamp)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			}
		})
	}
}
