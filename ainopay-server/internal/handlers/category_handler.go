package handlers

import (
	"ainopay-server/internal/repositories"
	"ainopay-server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryHandler(categoryRepo *repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{categoryRepo: categoryRepo}
}

// GetAll godoc
// @Summary Get all categories
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response
// @Router /categories [get]
func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.categoryRepo.FindAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}
