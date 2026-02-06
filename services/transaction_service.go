package services

import (
	"simple-crud-3/models"
	"simple-crud-3/repositories"
)


type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.Create(items)
}