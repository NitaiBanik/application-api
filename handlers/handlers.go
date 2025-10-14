package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type TestApiRequest map[string]interface{}

type TestApiResponse struct {
	Data      TestApiRequest `json:"data"`
	Timestamp time.Time      `json:"timestamp"`
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

type InfoResponse struct {
	Message   string   `json:"message"`
	Endpoints []string `json:"endpoints"`
	Version   string   `json:"version"`
}

func (h *Handler) TestApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Test API request received - %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

	if r.Method != http.MethodPost {
		log.Printf("Invalid method %s for testapi endpoint", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var payload TestApiRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Failed to decode JSON: %v", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if len(payload) == 0 {
		log.Printf("Empty payload received")
		http.Error(w, "Request body cannot be empty", http.StatusBadRequest)
		return
	}

	response := TestApiResponse{
		Data:      payload,
		Timestamp: time.Now().UTC(),
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status:    "healthy",
		Service:   "application-api",
		Timestamp: time.Now().UTC(),
		Version:   "1.0.0",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
