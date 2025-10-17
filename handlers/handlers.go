package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type TestAPIRequest map[string]interface{}

type TestAPIResponse struct {
	Data TestAPIRequest `json:"data"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *Handler) TestAPIHandler(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "unknown"
	}

	var payload TestAPIRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if len(payload) == 0 {
		http.Error(w, "Request body cannot be empty", http.StatusBadRequest)
		return
	}

	log.Printf("APP-API-%s: Processing request", port)

	response := TestAPIResponse{Data: payload}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "healthy"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
