package repositories

import (
	"ainopay-server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasswordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) *PasswordResetRepository {
	return &PasswordResetRepository{db: db}
}

// Create creates a new password reset token
func (r *PasswordResetRepository) Create(token *models.PasswordResetToken) error {
	return r.db.Create(token).Error
}

// FindByToken finds a token by its string value
func (r *PasswordResetRepository) FindByToken(token string) (*models.PasswordResetToken, error) {
	var prt models.PasswordResetToken
	// Preload User to ensure user exists
	if err := r.db.Preload("User").Where("token = ?", token).First(&prt).Error; err != nil {
		return nil, err
	}
	return &prt, nil
}

// MarkAsUsed marks a token as used
func (r *PasswordResetRepository) MarkAsUsed(id uuid.UUID) error {
	return r.db.Model(&models.PasswordResetToken{}).Where("id = ?", id).Update("used", true).Error
}

// DeleteByUserID deletes all tokens for a user (cleanup)
func (r *PasswordResetRepository) DeleteByUserID(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.PasswordResetToken{}).Error
}

// CleanupExpired deletes expired tokens
func (r *PasswordResetRepository) CleanupExpired() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&models.PasswordResetToken{}).Error
}
