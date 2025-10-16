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

func startServer() *exec.Cmd {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Start()
	time.Sleep(2 * time.Second)
	return cmd
}

func TestAPI(t *testing.T) {
	cmd := startServer()
	defer cmd.Process.Kill()

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

	client := &http.Client{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "http://localhost:8080/testapi", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedStatus == 200 {
				var response map[string]interface{}
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)

				originalData, _ := json.Marshal(tt.body)
				echoedData, _ := json.Marshal(response["data"])
				assert.Equal(t, string(originalData), string(echoedData))
			}
		})
	}

}

func TestHealthEndpoint(t *testing.T) {
	cmd := startServer()
	defer cmd.Process.Kill()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/health", nil)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Health check failed: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", response["status"])
}
