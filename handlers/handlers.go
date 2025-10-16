package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type TestApiRequest map[string]interface{}

type TestApiResponse struct {
	Data TestApiRequest `json:"data"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *Handler) TestApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Test API request received - %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

	port := r.Host
	if port == "" {
		port = "unknown"
	}

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

	response := TestApiResponse{Data: payload}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "healthy"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
