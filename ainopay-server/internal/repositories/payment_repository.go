package repositories

import (
	"ainopay-server/internal/models"
	//"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *PaymentRepository) FindByID(id uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("User").Preload("PaymentMethod").Preload("Category").
		First(&payment, "id = ?", id).Error
	return &payment, err
}

// PaymentFilter options
type PaymentFilter struct {
	Limit     int
	Offset    int
	Status    string
	Search    string
	MinAmount *float64
	MaxAmount *float64
	StartDate *time.Time
	EndDate   *time.Time
}

func (r *PaymentRepository) FindAll(userID uuid.UUID, filter PaymentFilter) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	query := r.db.Model(&models.Payment{}).Where("user_id = ?", userID)

	// Filter by status
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	// Search in description
	if filter.Search != "" {
		query = query.Where("description ILIKE ?", "%"+filter.Search+"%")
	}

	// Amount range
	if filter.MinAmount != nil {
		query = query.Where("amount >= ?", *filter.MinAmount)
	}
	if filter.MaxAmount != nil {
		query = query.Where("amount <= ?", *filter.MaxAmount)
	}

	// Date range
	if filter.StartDate != nil {
		query = query.Where("transaction_date >= ?", *filter.StartDate)
	}
	if filter.EndDate != nil {
		// Ensure end date includes the whole day
		query = query.Where("transaction_date <= ?", *filter.EndDate)
	}

	// Get total count
	query.Count(&total)

	// Get paginated results with preloaded relations
	query = query.Preload("User").Preload("PaymentMethod").Preload("Category").
		Order("transaction_date DESC")

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit).Offset(filter.Offset)
	}

	err := query.Find(&payments).Error

	return payments, total, err
}

func (r *PaymentRepository) Update(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *PaymentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Payment{}, "id = ?", id).Error
}

// GetStatistics returns payment statistics for a user
func (r *PaymentRepository) GetStatistics(userID uuid.UUID) (map[string]interface{}, error) {
	var totalPayments int64
	var totalAmount float64
	var completedCount int64
	var pendingCount int64

	r.db.Model(&models.Payment{}).Where("user_id = ?", userID).Count(&totalPayments)
	r.db.Model(&models.Payment{}).Where("user_id = ? AND status = ?", userID, "completed").Count(&completedCount)
	r.db.Model(&models.Payment{}).Where("user_id = ? AND status = ?", userID, "pending").Count(&pendingCount)
	r.db.Model(&models.Payment{}).Where("user_id = ? AND status = ?", userID, "completed").
		Select("COALESCE(SUM(amount), 0)").Scan(&totalAmount)

	return map[string]interface{}{
		"total_payments":    totalPayments,
		"completed_count":   completedCount,
		"pending_count":     pendingCount,
		"total_amount":      totalAmount,
	}, nil
}

// GetMonthlyEarnings returns earnings grouped by month for a specific year
func (r *PaymentRepository) GetMonthlyEarnings(userID uuid.UUID, year int) ([]models.MonthlyStats, error) {
	var stats []models.MonthlyStats

	// Postgres specific query
	err := r.db.Model(&models.Payment{}).
		Select("TO_CHAR(transaction_date, 'Mon') as month, SUM(amount) as total_amount, COUNT(*) as count").
		Where("user_id = ? AND status = ? AND EXTRACT(YEAR FROM transaction_date) = ?", userID, "completed", year).
		Group("TO_CHAR(transaction_date, 'Mon'), EXTRACT(MONTH FROM transaction_date)").
		Order("EXTRACT(MONTH FROM transaction_date)").
		Scan(&stats).Error

	return stats, err
}
