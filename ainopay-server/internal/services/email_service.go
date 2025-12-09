package services

import (
	"fmt"
	"log"
)

// EmailService handles email sending
type EmailService struct {
	// In production, you would add SMTP config or API client here
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

// SendPasswordResetEmail sends a password reset link to the user
func (s *EmailService) SendPasswordResetEmail(email, token string) error {
	// For MVP/Demo, we'll just log the link to the console
	// In production, use SMTP or SendGrid/Mailgun
	
	resetLink := fmt.Sprintf("http://localhost:3000/auth/reset-password?token=%s", token)
	
	log.Printf("----------------------------------------------------------------")
	log.Printf("📧 EMAIL SIMULATION - Password Reset")
	log.Printf("To: %s", email)
	log.Printf("Subject: Reset Your Password")
	log.Printf("Body: Click here to reset your password: %s", resetLink)
	log.Printf("----------------------------------------------------------------")
	
	return nil
}
