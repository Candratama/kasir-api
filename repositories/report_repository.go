package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetDailySalesReport() (*models.DailySalesReport, error) {
	today := time.Now().Format("2006-01-02")
	return repo.GetSalesReportByDateRange(today, today)
}

func (repo *ReportRepository) GetSalesReportByDateRange(startDate, endDate string) (*models.DailySalesReport, error) {
	report := &models.DailySalesReport{}

	// Get total revenue and total transactions
	summaryQuery := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) >= $1 AND DATE(created_at) <= $2
	`
	err := repo.db.QueryRow(summaryQuery, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get top selling product
	topProductQuery := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) >= $1 AND DATE(t.created_at) <= $2
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`
	var topProduct models.TopProduct
	err = repo.db.QueryRow(topProductQuery, startDate, endDate).Scan(&topProduct.Nama, &topProduct.QtyTerjual)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = nil
	} else if err != nil {
		return nil, err
	} else {
		report.ProdukTerlaris = &topProduct
	}

	return report, nil
}
