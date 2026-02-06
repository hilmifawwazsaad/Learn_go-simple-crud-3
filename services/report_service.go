package services

import (
	"simple-crud-3/models"
	"simple-crud-3/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetSalesReport(startDate, endDate string) (*models.SalesReport, error) {
	return s.repo.GetSalesReport(startDate, endDate)
}
