package database

import (
	"ainopay-server/internal/config"
	"ainopay-server/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establishes database connection
func Connect(cfg *config.DatabaseConfig) error {
	dsn := cfg.DSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return nil
}

// Migrate runs database migrations
func Migrate() error {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.PaymentMethod{},
		&models.Payment{},
	)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

// Seed inserts initial data
func Seed() error {
	log.Println("Seeding database...")

	// Seed payment methods
	paymentMethods := []models.PaymentMethod{
		{Name: "Bank Transfer", Code: "bank_transfer", IsActive: true},
		{Name: "Credit Card", Code: "credit_card", IsActive: true},
		{Name: "E-Wallet", Code: "e_wallet", IsActive: true},
		{Name: "Cash", Code: "cash", IsActive: true},
	}

	for _, pm := range paymentMethods {
		var existing models.PaymentMethod
		if err := DB.Where("code = ?", pm.Code).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := DB.Create(&pm).Error; err != nil {
				return fmt.Errorf("failed to seed payment method: %w", err)
			}
		}
	}

	// Seed categories
	categories := []models.Category{
		{Name: "Subscription", Description: "Monthly or yearly subscriptions"},
		{Name: "Purchase", Description: "One-time purchases"},
		{Name: "Service", Description: "Service payments"},
		{Name: "Donation", Description: "Charitable donations"},
		{Name: "Other", Description: "Other payments"},
	}

	for _, cat := range categories {
		var existing models.Category
		if err := DB.Where("name = ?", cat.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := DB.Create(&cat).Error; err != nil {
				return fmt.Errorf("failed to seed category: %w", err)
			}
		}
	}

	log.Println("Database seeding completed successfully")
	return nil
}
