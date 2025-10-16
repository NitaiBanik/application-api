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
}

func (h *Handler) TestApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Test API request received - %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	port := r.Host
	if port == "" {
		port = "unknown"
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

	log.Printf("APP-API-%s: Processing request", port)
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
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status:    "healthy",
		Service:   "application-api",
		Timestamp: time.Now().UTC(),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
