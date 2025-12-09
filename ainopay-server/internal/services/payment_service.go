package services

import (
	"ainopay-server/internal/models"
	"ainopay-server/internal/repositories"
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

func (s *PaymentService) GetAll(userID uuid.UUID, page, limit int, status, search string) (*PaymentListResponse, error) {
	offset := (page - 1) * limit

	payments, total, err := s.paymentRepo.FindAll(userID, limit, offset, status, search)
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
