package repositories

import (
	"ainopay-server/internal/models"

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

func (r *PaymentRepository) FindAll(userID uuid.UUID, limit, offset int, status, search string) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	query := r.db.Model(&models.Payment{}).Where("user_id = ?", userID)

	// Filter by status
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Search in description
	if search != "" {
		query = query.Where("description ILIKE ?", "%"+search+"%")
	}

	// Get total count
	query.Count(&total)

	// Get paginated results with preloaded relations
	err := query.Preload("User").Preload("PaymentMethod").Preload("Category").
		Order("transaction_date DESC").
		Limit(limit).Offset(offset).
		Find(&payments).Error

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
