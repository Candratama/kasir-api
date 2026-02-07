package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
	"strings"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// HandleDailyReport godoc
// @Summary Get daily sales report
// @Description Mengambil laporan penjualan hari ini
// @Tags Reports
// @Produce json
// @Success 200 {object} models.DailySalesReport
// @Failure 500 {object} map[string]string
// @Router /api/report/hari-ini [get]
func (h *ReportHandler) HandleDailyReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	report, err := h.service.GetDailySalesReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// HandleReport godoc
// @Summary Get sales report by date range
// @Description Mengambil laporan penjualan berdasarkan rentang tanggal
// @Tags Reports
// @Produce json
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} models.DailySalesReport
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/report [get]
func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if this is the /hari-ini endpoint
	if strings.HasSuffix(r.URL.Path, "/hari-ini") {
		h.HandleDailyReport(w, r)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	report, err := h.service.GetSalesReportByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
