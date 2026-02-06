package repositories

import (
	"database/sql"
	"simple-crud-3/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetSalesReport(startDate, endDate string) (*models.SalesReport, error) {
	report := &models.SalesReport{}

	var summaryQuery string
	var topProductQuery string
	hasDateFilter := startDate != "" && endDate != ""

	// Get total revenue and total transaksi
	if hasDateFilter {
		summaryQuery = `
			SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
			FROM transactions
			WHERE created_at::date BETWEEN $1 AND $2`
	} else {
		summaryQuery = `
			SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
			FROM transactions`
	}

	var err error
	if hasDateFilter {
		err = r.db.QueryRow(summaryQuery, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	} else {
		err = r.db.QueryRow(summaryQuery).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	}
	if err != nil {
		return nil, err
	}

	// Get produk terlaris
	if hasDateFilter {
		topProductQuery = `
			SELECT p.name, SUM(td.quantity) AS total_qty
			FROM transaction_details td
			JOIN transactions t ON t.id = td.transaction_id
			JOIN products p ON p.id = td.product_id
			WHERE t.created_at::date BETWEEN $1 AND $2
			GROUP BY p.name
			ORDER BY total_qty DESC
			LIMIT 1`
	} else {
		topProductQuery = `
			SELECT p.name, SUM(td.quantity) AS total_qty
			FROM transaction_details td
			JOIN transactions t ON t.id = td.transaction_id
			JOIN products p ON p.id = td.product_id
			GROUP BY p.name
			ORDER BY total_qty DESC
			LIMIT 1`
	}

	var produk models.ProdukTerlaris
	if hasDateFilter {
		err = r.db.QueryRow(topProductQuery, startDate, endDate).Scan(&produk.Nama, &produk.QtyTerjual)
	} else {
		err = r.db.QueryRow(topProductQuery).Scan(&produk.Nama, &produk.QtyTerjual)
	}
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = nil
		return report, nil
	}
	if err != nil {
		return nil, err
	}

	report.ProdukTerlaris = &produk
	return report, nil
}
