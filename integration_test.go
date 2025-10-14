package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Start()
	defer cmd.Process.Kill()

	time.Sleep(2 * time.Second)

	tests := []struct {
		name           string
		body           interface{}
		expectedStatus int
	}{
		{
			name:           "should handle basic JSON payload",
			body:           map[string]interface{}{"test": true},
			expectedStatus: 200,
		},
		{
			name:           "should work with nested objects and arrays",
			body:           map[string]interface{}{"user": map[string]interface{}{"id": 123, "name": "Test"}, "items": []int{1, 2, 3}},
			expectedStatus: 200,
		},
		{
			name:           "should reject empty JSON object",
			body:           map[string]interface{}{},
			expectedStatus: 400,
		},
		{
			name:           "should echo back string messages",
			body:           map[string]interface{}{"message": "Hello World"},
			expectedStatus: 200,
		},
	}

	httpClient := &http.Client{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "http://localhost:8080/testapi", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := httpClient.Do(req)
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedStatus == 200 {
				var response map[string]interface{}
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.NotNil(t, response["data"])
				assert.NotNil(t, response["timestamp"])

				originalData, _ := json.Marshal(tt.body)
				echoedData, _ := json.Marshal(response["data"])
				assert.Equal(t, string(originalData), string(echoedData))
			}
		})
	}
}
