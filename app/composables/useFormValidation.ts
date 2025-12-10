import type { ZodSchema } from 'zod'

/**
 * Composable for form validation using Zod schemas
 * @param schema - Zod schema to validate against
 * @returns Validation utilities and reactive errors
 */
export function useFormValidation<T extends ZodSchema>(schema: T) {
  // Reactive errors object
  const errors = reactive<Record<string, string>>({})

  /**
   * Validate entire form data against schema
   * @param data - Form data to validate
   * @returns true if valid, false if invalid
   */
  const validate = (data: unknown): boolean => {
    // Clear previous errors
    clearErrors()

    // Validate with Zod
    const result = schema.safeParse(data)

    if (!result.success) {
      // Map Zod errors to our errors object
      result.error.issues.forEach(issue => {
        const field = issue.path.join('.')
        errors[field] = issue.message
      })
      return false
    }

    return true
  }

  /**
   * Validate a single field
   * @param field - Field name to validate
   * @param value - Field value
   * @returns true if valid, false if invalid
   */
  const validateField = (field: string, value: unknown): boolean => {
    // Clear field error
    clearFieldError(field)

    // Get field schema from the main schema
    if ('shape' in schema && schema.shape) {
      const fieldSchema = (schema.shape as Record<string, ZodSchema>)[field]

      if (fieldSchema) {
        const result = fieldSchema.safeParse(value)

        if (!result.success) {
          errors[field] = result.error.issues[0]?.message || 'Invalid value'
          return false
        }
      }
    }

    return true
  }

  /**
   * Clear all validation errors
   */
  const clearErrors = () => {
    Object.keys(errors).forEach(key => {
      errors[key] = ''
    })
  }

  /**
   * Clear error for a specific field
   * @param field - Field name to clear error for
   */
  const clearFieldError = (field: string) => {
    errors[field] = ''
  }

  /**
   * Check if form has any errors
   */
  const hasErrors = computed(() => Object.keys(errors).length > 0)

  /**
   * Get error for a specific field
   * @param field - Field name
   */
  const getError = (field: string): string => errors[field] || ''

  return {
    errors: readonly(errors),
    validate,
    validateField,
    clearErrors,
    clearFieldError,
    hasErrors,
    getError,
  }
}
