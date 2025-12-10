/**
 * Error handling composable for centralized error management
 */
export function useErrorHandler() {
  const uiStore = useUIStore()

  /**
   * Get user-friendly error message from error object
   */
  const getErrorMessage = (error: unknown): string => {
    // Handle API response errors
    if (typeof error === 'object' && error !== null && 'error' in error) {
      return (error as { error: string }).error
    }

    // Handle Error instances
    if (error instanceof Error) {
      // Network errors
      if (error.message.includes('fetch') || error.message.includes('network')) {
        return 'Unable to connect to server. Please check your internet connection.'
      }
      return error.message
    }

    // Handle string errors
    if (typeof error === 'string') {
      return error
    }

    return 'An unexpected error occurred'
  }

  /**
   * Get error message based on HTTP status code
   */
  const getStatusErrorMessage = (status: number): string => {
    const messages: Record<number, string> = {
      400: 'Invalid request. Please check your input.',
      401: 'Your session has expired. Please login again.',
      403: "You don't have permission to access this resource.",
      404: 'The requested resource was not found.',
      408: 'Request timeout. Please try again.',
      429: 'Too many requests. Please try again later.',
      500: 'Something went wrong on our end. Please try again later.',
      502: 'Server is temporarily unavailable. Please try again later.',
      503: 'Service is temporarily unavailable. Please try again later.',
    }

    return messages[status] || 'An error occurred. Please try again.'
  }

  /**
   * Log error to console (and future error tracking service)
   */
  const logError = (error: unknown, context?: string) => {
    const logger = useLogger(context || 'Error')
    const errorMessage = getErrorMessage(error)

    logger.error(errorMessage, {
      error,
      url: typeof window !== 'undefined' ? window.location.href : undefined,
    })

    // TODO: Send to error tracking service (e.g., Sentry)
  }

  /**
   * Handle API response errors
   */
  const handleApiError = (response: { success: boolean; error?: string }, context?: string) => {
    if (!response.success && response.error) {
      logError(response.error, context)
      showErrorToast(response.error)
    }
  }

  /**
   * Show error toast notification
   */
  const showErrorToast = (message: string) => {
    uiStore.showToast(message, 'error')
  }

  /**
   * Handle general errors with optional context
   */
  const handleError = (error: unknown, context?: string) => {
    logError(error, context)

    const message = getErrorMessage(error)

    // Handle authentication errors
    if (
      message.includes('401') ||
      message.includes('unauthorized') ||
      message.includes('session')
    ) {
      showErrorToast('Your session has expired. Please login again.')
      const authStore = useAuthStore()
      authStore.logout()
      return
    }

    // Show error toast
    showErrorToast(message)
  }

  /**
   * Handle navigation errors
   */
  const handleNavigationError = (error: unknown) => {
    logError(error, 'Navigation')

    // Don't show toast for navigation cancellations
    if (error instanceof Error && error.message.includes('navigation')) {
      return
    }

    handleError(error)
  }

  /**
   * Create error for Nuxt error page
   */
  const createError = (statusCode: number, message?: string) => {
    const errorMessage = message || getStatusErrorMessage(statusCode)

    return {
      statusCode,
      message: errorMessage,
      fatal: statusCode >= 500,
    }
  }

  return {
    getErrorMessage,
    getStatusErrorMessage,
    logError,
    handleApiError,
    handleError,
    handleNavigationError,
    showErrorToast,
    createError,
  }
}
