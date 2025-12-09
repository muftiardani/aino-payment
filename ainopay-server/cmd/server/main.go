package main

import (
	"ainopay-server/internal/config"
	"ainopay-server/internal/database"
	"ainopay-server/internal/handlers"
	"ainopay-server/internal/middleware"
	"ainopay-server/internal/repositories"
	"ainopay-server/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ainopay-server/docs" // Import generated docs
)

// @title AinoPay API
// @version 1.0
// @description Payment Management System API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@ainopay.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	gin.SetMode(cfg.Server.GinMode)

	// Connect to database
	if err := database.Connect(&cfg.Database); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Seed database
	if err := database.Seed(); err != nil {
		log.Fatal("Failed to seed database:", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)
	paymentRepo := repositories.NewPaymentRepository(database.DB)
	categoryRepo := repositories.NewCategoryRepository(database.DB)
	paymentMethodRepo := repositories.NewPaymentMethodRepository(database.DB)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg)
	paymentService := services.NewPaymentService(paymentRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)
	dashboardHandler := handlers.NewDashboardHandler(paymentService, paymentRepo)
	paymentMethodHandler := handlers.NewPaymentMethodHandler(paymentMethodRepo)

	// Setup router
	router := gin.Default()

	// Global middleware
	router.Use(middleware.CORSMiddleware(cfg.CORS.AllowedOrigins))
	router.Use(middleware.LoggerMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// API routes
	api := router.Group("/api")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/me", middleware.AuthMiddleware(cfg), authHandler.GetMe)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			// Payment routes
			payments := protected.Group("/payments")
			{
				payments.POST("", paymentHandler.Create)
				payments.GET("", paymentHandler.GetAll)
				payments.GET("/:id", paymentHandler.GetByID)
				payments.PUT("/:id", paymentHandler.Update)
				payments.DELETE("/:id", paymentHandler.Delete)
			}

			// Category routes
			categories := protected.Group("/categories")
			{
				categories.GET("", categoryHandler.GetAll)
			}

			// Payment method routes
			paymentMethods := protected.Group("/payment-methods")
			{
				paymentMethods.GET("", paymentMethodHandler.GetAll)
			}

			// Dashboard routes
			dashboard := protected.Group("/dashboard")
			{
				dashboard.GET("/stats", dashboardHandler.GetStats)
				dashboard.GET("/recent", dashboardHandler.GetRecent)
			}
		}
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
