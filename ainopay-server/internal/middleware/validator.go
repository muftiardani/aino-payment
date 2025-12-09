package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateStruct validates a struct and returns formatted errors
func ValidateStruct(obj interface{}) []ValidationError {
	var errors []ValidationError

	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.Field = err.Field()
			element.Message = getErrorMessage(err)
			errors = append(errors, element)
		}
	}

	return errors
}

// getErrorMessage returns user-friendly error message based on validation tag
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short (minimum " + err.Param() + " characters)"
	case "max":
		return "Value is too long (maximum " + err.Param() + " characters)"
	case "gt":
		return "Value must be greater than " + err.Param()
	case "gte":
		return "Value must be greater than or equal to " + err.Param()
	case "lt":
		return "Value must be less than " + err.Param()
	case "lte":
		return "Value must be less than or equal to " + err.Param()
	case "uuid":
		return "Invalid UUID format"
	case "uuid4":
		return "Invalid UUID v4 format"
	default:
		return "Invalid value"
	}
}

// ValidateRequest is a middleware that validates request body
func ValidateRequest(validationFunc func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj, err := validationFunc(c)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "Invalid request format",
			})
			c.Abort()
			return
		}

		// Validate the struct
		validationErrors := ValidateStruct(obj)
		if len(validationErrors) > 0 {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "Validation failed",
				"details": validationErrors,
			})
			c.Abort()
			return
		}

		// Store validated data in context
		c.Set("validated_data", obj)
		c.Next()
	}
}
