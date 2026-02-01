package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

var db *sql.DB

// SetDB sets the database connection for health checks
func SetDB(database *sql.DB) {
	db = database
}

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

// @Summary Database Health Check
// @Description Memeriksa koneksi ke database
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /health/db [get]
func DBHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if db == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ERROR",
			"message": "Database not initialized",
		})
		return
	}

	err := db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ERROR",
			"message": "Database connection failed: " + err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "Database connected",
	})
}
