package repositories

import (
	"ainopay-server/internal/models"

	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

func (r *PaymentMethodRepository) FindAll() ([]models.PaymentMethod, error) {
	var methods []models.PaymentMethod
	err := r.db.Where("is_active = ?", true).Order("name ASC").Find(&methods).Error
	return methods, err
}
