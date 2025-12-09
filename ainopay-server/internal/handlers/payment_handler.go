package handlers

import (
	"ainopay-server/internal/middleware"
	"ainopay-server/internal/services"
	"ainopay-server/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

// CreatePaymentRequest represents the request body for creating a payment
type CreatePaymentRequest struct {
	Amount          float64 `json:"amount" validate:"required,gt=0"`
	CategoryID      string  `json:"category_id" validate:"required,uuid4"`
	PaymentMethodID string  `json:"payment_method_id" validate:"required,uuid4"`
	Description     string  `json:"description" validate:"max=500"`
	TransactionDate string  `json:"transaction_date" validate:"required"`
}

// UpdatePaymentRequest represents the request body for updating a payment
type UpdatePaymentRequest struct {
	Amount          float64 `json:"amount" validate:"required,gt=0"`
	Status          string  `json:"status" validate:"required,oneof=pending completed failed refunded"`
	CategoryID      string  `json:"category_id" validate:"required,uuid4"`
	PaymentMethodID string  `json:"payment_method_id" validate:"required,uuid4"`
	Description     string  `json:"description" validate:"max=500"`
	TransactionDate string  `json:"transaction_date" validate:"required"`
}

// Create godoc
// @Summary Create new payment
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body services.CreatePaymentRequest true "Create Payment Request"
// @Success 201 {object} utils.Response
// @Router /payments [post]
func (h *PaymentHandler) Create(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(uuid.UUID)

	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Validate request
	if validationErrors := middleware.ValidateStruct(&req); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validation failed",
			"details": validationErrors,
		})
		return
	}

	// Parse transaction date
	transactionDate, err := time.Parse(time.RFC3339, req.TransactionDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid transaction date format")
		return
	}

	// Parse UUIDs
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	paymentMethodID, err := uuid.Parse(req.PaymentMethodID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID")
		return
	}

	// Convert to service request
	serviceReq := &services.CreatePaymentRequest{
		Amount:          req.Amount,
		CategoryID:      categoryID,
		PaymentMethodID: paymentMethodID,
		Description:     req.Description,
		TransactionDate: transactionDate,
	}

	payment, err := h.paymentService.Create(id, serviceReq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Payment created successfully", payment)
}

// GetAll godoc
// @Summary Get all payments
// @Tags payments
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param status query string false "Filter by status"
// @Param search query string false "Search in description"
// @Param min_amount query number false "Minimum amount"
// @Param max_amount query number false "Maximum amount"
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} utils.Response
// @Router /payments [get]
func (h *PaymentHandler) GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(uuid.UUID)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.Query("status")
	search := c.Query("search")

	// Parse advanced filters
	var minAmount, maxAmount *float64
	if val := c.Query("min_amount"); val != "" {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			minAmount = &v
		}
	}
	if val := c.Query("max_amount"); val != "" {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			maxAmount = &v
		}
	}

	var startDate, endDate *time.Time
	if val := c.Query("start_date"); val != "" {
		if t, err := time.Parse("2006-01-02", val); err == nil {
			startDate = &t
		}
	}
	if val := c.Query("end_date"); val != "" {
		if t, err := time.Parse("2006-01-02", val); err == nil {
			// Set to end of day
			t = t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			endDate = &t
		}
	}

	result, err := h.paymentService.GetAll(id, page, limit, status, search, minAmount, maxAmount, startDate, endDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payments retrieved successfully", result)
}

// Export godoc
// @Summary Export payments to CSV
// @Tags payments
// @Produce text/csv
// @Security BearerAuth
// @Param status query string false "Filter by status"
// @Param search query string false "Search in description"
// @Param min_amount query number false "Minimum amount"
// @Param max_amount query number false "Maximum amount"
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Success 200 {file} file "payments.csv"
// @Router /payments/export [get]
func (h *PaymentHandler) Export(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(uuid.UUID)

	status := c.Query("status")
	search := c.Query("search")

	// Parse advanced filters
	var minAmount, maxAmount *float64
	if val := c.Query("min_amount"); val != "" {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			minAmount = &v
		}
	}
	if val := c.Query("max_amount"); val != "" {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			maxAmount = &v
		}
	}

	var startDate, endDate *time.Time
	if val := c.Query("start_date"); val != "" {
		if t, err := time.Parse("2006-01-02", val); err == nil {
			startDate = &t
		}
	}
	if val := c.Query("end_date"); val != "" {
		if t, err := time.Parse("2006-01-02", val); err == nil {
			// Set to end of day
			t = t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			endDate = &t
		}
	}

	csvData, err := h.paymentService.Export(id, status, search, minAmount, maxAmount, startDate, endDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=payments.csv")
	c.Data(http.StatusOK, "text/csv", csvData)
}

// GetByID godoc
// @Summary Get payment by ID
// @Tags payments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} utils.Response
// @Router /payments/{id} [get]
func (h *PaymentHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment ID")
		return
	}

	payment, err := h.paymentService.GetByID(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Payment not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment retrieved successfully", payment)
}

// Update godoc
// @Summary Update payment
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Param request body services.UpdatePaymentRequest true "Update Payment Request"
// @Success 200 {object} utils.Response
// @Router /payments/{id} [put]
func (h *PaymentHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment ID")
		return
	}

	var req UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Validate request
	if validationErrors := middleware.ValidateStruct(&req); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validation failed",
			"details": validationErrors,
		})
		return
	}

	// Parse transaction date
	transactionDate, err := time.Parse(time.RFC3339, req.TransactionDate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid transaction date format")
		return
	}

	// Parse UUIDs
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid category ID")
		return
	}

	paymentMethodID, err := uuid.Parse(req.PaymentMethodID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment method ID")
		return
	}

	// Convert to service request
	serviceReq := &services.UpdatePaymentRequest{
		Amount:          req.Amount,
		Status:          req.Status,
		CategoryID:      categoryID,
		PaymentMethodID: paymentMethodID,
		Description:     req.Description,
		TransactionDate: transactionDate,
	}

	payment, err := h.paymentService.Update(id, serviceReq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment updated successfully", payment)
}

// Delete godoc
// @Summary Delete payment
// @Tags payments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} utils.Response
// @Router /payments/{id} [delete]
func (h *PaymentHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid payment ID")
		return
	}

	if err := h.paymentService.Delete(id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment deleted successfully", nil)
}
