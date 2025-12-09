package handlers

import (
	"ainopay-server/internal/repositories"
	"ainopay-server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentMethodHandler struct {
	paymentMethodRepo *repositories.PaymentMethodRepository
}

func NewPaymentMethodHandler(paymentMethodRepo *repositories.PaymentMethodRepository) *PaymentMethodHandler {
	return &PaymentMethodHandler{paymentMethodRepo: paymentMethodRepo}
}

// GetAll godoc
// @Summary Get all payment methods
// @Tags payment-methods
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response
// @Router /payment-methods [get]
func (h *PaymentMethodHandler) GetAll(c *gin.Context) {
	methods, err := h.paymentMethodRepo.FindAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Payment methods retrieved successfully", methods)
}
