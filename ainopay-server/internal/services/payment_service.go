package services

import (
	"ainopay-server/internal/models"
	"ainopay-server/internal/repositories"
	"bytes"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepo *repositories.PaymentRepository
}

func NewPaymentService(paymentRepo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{paymentRepo: paymentRepo}
}

type CreatePaymentRequest struct {
	Amount          float64   `json:"amount" binding:"required,gt=0"`
	PaymentMethodID uuid.UUID `json:"payment_method_id" binding:"required"`
	CategoryID      uuid.UUID `json:"category_id" binding:"required"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date" binding:"required"`
}

type UpdatePaymentRequest struct {
	Amount          float64   `json:"amount" binding:"required,gt=0"`
	Status          string    `json:"status" binding:"required,oneof=pending completed failed refunded"`
	PaymentMethodID uuid.UUID `json:"payment_method_id" binding:"required"`
	CategoryID      uuid.UUID `json:"category_id" binding:"required"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date" binding:"required"`
}

type PaymentListResponse struct {
	Payments []models.Payment `json:"payments"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	Limit    int              `json:"limit"`
}

func (s *PaymentService) Create(userID uuid.UUID, req *CreatePaymentRequest) (*models.Payment, error) {
	payment := &models.Payment{
		UserID:          userID,
		Amount:          req.Amount,
		Status:          "pending",
		PaymentMethodID: req.PaymentMethodID,
		CategoryID:      req.CategoryID,
		Description:     req.Description,
		TransactionDate: req.TransactionDate,
	}

	if err := s.paymentRepo.Create(payment); err != nil {
		return nil, err
	}

	// Reload with relations
	return s.paymentRepo.FindByID(payment.ID)
}

func (s *PaymentService) GetByID(id uuid.UUID) (*models.Payment, error) {
	return s.paymentRepo.FindByID(id)
}

func (s *PaymentService) GetAll(userID uuid.UUID, page, limit int, status, search string, minAmount, maxAmount *float64, startDate, endDate *time.Time) (*PaymentListResponse, error) {
	offset := (page - 1) * limit

	filter := repositories.PaymentFilter{
		Limit:     limit,
		Offset:    offset,
		Status:    status,
		Search:    search,
		MinAmount: minAmount,
		MaxAmount: maxAmount,
		StartDate: startDate,
		EndDate:   endDate,
	}

	payments, total, err := s.paymentRepo.FindAll(userID, filter)
	if err != nil {
		return nil, err
	}

	return &PaymentListResponse{
		Payments: payments,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (s *PaymentService) Update(id uuid.UUID, req *UpdatePaymentRequest) (*models.Payment, error) {
	payment, err := s.paymentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	payment.Amount = req.Amount
	payment.Status = req.Status
	payment.PaymentMethodID = req.PaymentMethodID
	payment.CategoryID = req.CategoryID
	payment.Description = req.Description
	payment.TransactionDate = req.TransactionDate

	if err := s.paymentRepo.Update(payment); err != nil {
		return nil, err
	}

	return s.paymentRepo.FindByID(id)
}

func (s *PaymentService) Delete(id uuid.UUID) error {
	return s.paymentRepo.Delete(id)
}

func (s *PaymentService) GetStatistics(userID uuid.UUID) (map[string]interface{}, error) {
	return s.paymentRepo.GetStatistics(userID)
}

func (s *PaymentService) GetMonthlyStats(userID uuid.UUID, year int) ([]models.MonthlyStats, error) {
	return s.paymentRepo.GetMonthlyEarnings(userID, year)
}

func (s *PaymentService) Export(userID uuid.UUID, status, search string, minAmount, maxAmount *float64, startDate, endDate *time.Time) ([]byte, error) {
	// Limit 0 means fetch all
	filter := repositories.PaymentFilter{
		Limit:     0,
		Status:    status,
		Search:    search,
		MinAmount: minAmount,
		MaxAmount: maxAmount,
		StartDate: startDate,
		EndDate:   endDate,
	}

	payments, _, err := s.paymentRepo.FindAll(userID, filter)
	if err != nil {
		return nil, err
	}

	// Generate CSV
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	// Write header
	header := []string{"Transaction Date", "Description", "Amount", "Category", "Payment Method", "Status"}
	if err := w.Write(header); err != nil {
		return nil, err
	}

	// Write rows
	for _, p := range payments {
		row := []string{
			p.TransactionDate.Format("2006-01-02 15:04"),
			p.Description,
			fmt.Sprintf("%.2f", p.Amount),
			p.Category.Name,
			p.PaymentMethod.Name,
			p.Status,
		}
		if err := w.Write(row); err != nil {
			return nil, err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
