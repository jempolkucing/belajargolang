package repositories

import (
	"belajargolang/models"
	"database/sql"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetReport(startDate, endDate time.Time) (*models.DailyReport, error) {
	report := &models.DailyReport{}

	queryStats := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*) 
		FROM transactions 
		WHERE created_at >= $1 AND created_at <= $2
	`
	err := r.db.QueryRow(queryStats, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	queryBestProduct := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0)
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE t.created_at >= $1 AND t.created_at <= $2
		GROUP BY p.name
		ORDER BY SUM(td.quantity) DESC
		LIMIT 1
	`
	err = r.db.QueryRow(queryBestProduct, startDate, endDate).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)
	if err != nil {
		if err == sql.ErrNoRows {
			report.ProdukTerlaris = models.ProductStats{Nama: "-", QtyTerjual: 0}
		} else {
			return nil, err
		}
	}

	return report, nil
}
