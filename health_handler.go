package main

import (
	"encoding/json"
	"net/http"
)

// @Summary Health Check
// @Description Memeriksa status kesehatan server
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}
