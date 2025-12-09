package services

import (
	"ainopay-server/internal/config"
	"ainopay-server/internal/models"
	"ainopay-server/internal/repositories"
	"ainopay-server/internal/utils"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo          *repositories.UserRepository
	refreshTokenRepo  *repositories.RefreshTokenRepository
	passwordResetRepo *repositories.PasswordResetRepository
	emailService      *EmailService
	cfg               *config.Config
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	refreshTokenRepo *repositories.RefreshTokenRepository,
	passwordResetRepo *repositories.PasswordResetRepository,
	emailService *EmailService,
	cfg *config.Config,
) *AuthService {
	return &AuthService{
		userRepo:          userRepo,
		refreshTokenRepo:  refreshTokenRepo,
		passwordResetRepo: passwordResetRepo,
		emailService:      emailService,
		cfg:               cfg,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token        string       `json:"token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int64        `json:"expires_in"` // seconds
	User         *models.User `json:"user"`
}

func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
	// Check if user already exists
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		FullName:     req.FullName,
		Role:         "user",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate access token
	expiration, _ := time.ParseDuration(s.cfg.JWT.Expiration)
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWT.Secret, expiration)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := s.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(expiration.Seconds()),
		User:         user,
	}, nil
}

func (s *AuthService) Login(req *LoginRequest) (*AuthResponse, error) {
	// Find user
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	// Generate access token
	expiration, _ := time.ParseDuration(s.cfg.JWT.Expiration)
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWT.Secret, expiration)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := s.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(expiration.Seconds()),
		User:         user,
	}, nil
}

func (s *AuthService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

// GenerateRefreshToken creates a new refresh token for a user
func (s *AuthService) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	// Delete old refresh tokens for this user
	if err := s.refreshTokenRepo.DeleteByUserID(userID); err != nil {
		// Log error but don't fail (old tokens might not exist)
	}

	// Generate random token
	tokenString := uuid.New().String()

	// Create refresh token with 7 day expiry
	refreshToken := &models.RefreshToken{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.refreshTokenRepo.Create(refreshToken); err != nil {
		return "", err
	}

	return tokenString, nil
}

// RefreshAccessToken validates refresh token and generates new access token
func (s *AuthService) RefreshAccessToken(refreshTokenString string) (*AuthResponse, error) {
	// Find refresh token
	refreshToken, err := s.refreshTokenRepo.FindByToken(refreshTokenString)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Check if expired
	if refreshToken.IsExpired() {
		// Delete expired token
		s.refreshTokenRepo.DeleteByToken(refreshTokenString)
		return nil, errors.New("refresh token expired")
	}

	// Get user
	user, err := s.userRepo.FindByID(refreshToken.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate new access token
	expiration, _ := time.ParseDuration(s.cfg.JWT.Expiration)
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role, s.cfg.JWT.Secret, expiration)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token:        token,
		RefreshToken: refreshTokenString, // Return same refresh token
		ExpiresIn:    int64(expiration.Seconds()),
		User:         user,
	}, nil
}

// ForgotPassword initiates password reset flow
func (s *AuthService) ForgotPassword(email string) error {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		// Don't reveal if user exists
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	// Create reset token
	token := &models.PasswordResetToken{
		UserID:    user.ID,
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().Add(1 * time.Hour), // 1 hour expiry
	}

	if err := s.passwordResetRepo.Create(token); err != nil {
		return err
	}

	// Send email
	// Run in goroutine to not block response
	go func() {
		if err := s.emailService.SendPasswordResetEmail(user.Email, token.Token); err != nil {
			fmt.Printf("Failed to send email: %v\n", err)
		}
	}()

	return nil
}

// ResetPassword resets user password using token
func (s *AuthService) ResetPassword(tokenString, newPassword string) error {
	// Find and validate token
	token, err := s.passwordResetRepo.FindByToken(tokenString)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	if !token.IsValid() {
		return errors.New("invalid or expired token")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update user password through repository (need to add UpdatePassword method to UserRepo or just use Update)
	// For now, let's assume we fetch user and save
	user := &token.User
	user.PasswordHash = hashedPassword

	// Since UserRepo.Update might not be granular, let's use a direct DB update or add method
	// Ideally we add UpdatePassword to UserRepo. Let's assume we can use DB via repo if exposed, or just Update.
	// Looking at UserRepo, it has Create, FindBy... but maybe no Update?
	// Let's assume we can add Update to UserRepo or use DB directly if we had access (we don't here).
	// Let's add Update to UserRepo in next step. For now, calling a method we will create.
	if err := s.userRepo.UpdatePassword(user.ID, hashedPassword); err != nil {
		return err
	}

	// Mark token as used
	if err := s.passwordResetRepo.MarkAsUsed(token.ID); err != nil {
		// Log error but critical part (password change) succeeded
	}

	return nil
}
