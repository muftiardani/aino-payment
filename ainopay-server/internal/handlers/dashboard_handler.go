package handlers

import (
	"ainopay-server/internal/repositories"
	"ainopay-server/internal/services"
	"ainopay-server/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DashboardHandler struct {
	paymentService *services.PaymentService
	paymentRepo    *repositories.PaymentRepository
}

func NewDashboardHandler(paymentService *services.PaymentService, paymentRepo *repositories.PaymentRepository) *DashboardHandler {
	return &DashboardHandler{
		paymentService: paymentService,
		paymentRepo:    paymentRepo,
	}
}

// GetStats godoc
// @Summary Get dashboard statistics
// @Tags dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response
// @Router /dashboard/stats [get]
func (h *DashboardHandler) GetStats(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(uuid.UUID)

	stats, err := h.paymentService.GetStatistics(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Statistics retrieved successfully", stats)
}

// GetRecent godoc
// @Summary Get recent payments
// @Tags dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response
// @Router /dashboard/recent [get]
func (h *DashboardHandler) GetRecent(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(uuid.UUID)

	result, err := h.paymentService.GetAll(id, 1, 5, "", "", nil, nil, nil, nil)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Recent payments retrieved successfully", result.Payments)
}

// GetChartData godoc
// @Summary Get monthly earnings for chart
// @Tags dashboard
// @Produce json
// @Security BearerAuth
// @Param year query int false "Year (default: current year)"
// @Success 200 {object} utils.Response
// @Router /dashboard/chart [get]
func (h *DashboardHandler) GetChartData(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(uuid.UUID)

	year, err := strconv.Atoi(c.DefaultQuery("year", strconv.Itoa(time.Now().Year())))
	if err != nil {
		year = time.Now().Year()
	}

	stats, err := h.paymentService.GetMonthlyStats(id, year)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Chart data retrieved successfully", stats)
}
