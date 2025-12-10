/**
 * Global error handler plugin
 * Catches unhandled errors and promise rejections
 */
export default defineNuxtPlugin(nuxtApp => {
  const { handleError } = useErrorHandler()
  const logger = useLogger('GlobalErrorHandler')

  // Handle Vue errors
  nuxtApp.vueApp.config.errorHandler = (error, instance, info) => {
    logger.error('Vue Error', {
      error,
      info,
      component: instance?.$options?.name || 'Unknown',
    })
    handleError(error, `Vue Error: ${info}`)
  }

  // Handle unhandled promise rejections (client-side only)
  if (import.meta.client) {
    window.addEventListener('unhandledrejection', event => {
      logger.error('Unhandled Promise Rejection', { reason: event.reason })
      handleError(event.reason, 'Unhandled Promise Rejection')
      event.preventDefault() // Prevent default browser error handling
    })

    // Handle general errors (client-side only)
    window.addEventListener('error', event => {
      logger.error('Global Error', { error: event.error, message: event.message })
      handleError(event.error, 'Global Error')
    })
  }

  // Provide error handler to the app
  return {
    provide: {
      handleError,
    },
  }
})
