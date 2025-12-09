package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	User            User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Amount          float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Status          string         `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, completed, failed, refunded
	PaymentMethodID uuid.UUID      `gorm:"type:uuid;not null" json:"payment_method_id"`
	PaymentMethod   PaymentMethod  `gorm:"foreignKey:PaymentMethodID" json:"payment_method,omitempty"`
	CategoryID      uuid.UUID      `gorm:"type:uuid;not null" json:"category_id"`
	Category        Category       `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Description     string         `gorm:"type:text" json:"description"`
	TransactionDate time.Time      `gorm:"not null" json:"transaction_date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

type MonthlyStats struct {
	Month       string  `json:"month"`
	TotalAmount float64 `json:"total_amount"`
	Count       int64   `json:"count"`
}
