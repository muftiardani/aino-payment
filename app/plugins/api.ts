import type { ApiResponse } from '~/types/api'

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()
  const authStore = useAuthStore()
  const logger = useLogger('API')

  const api = async <T = unknown>(
    endpoint: string,
    options: RequestInit & {
      params?: Record<string, string | number | boolean>
    } = {}
  ): Promise<ApiResponse<T>> => {
    try {
      // Build URL with query params
      let url = `${config.public.apiBase}${endpoint}`
      if (options.params) {
        const queryString = new URLSearchParams(
          Object.entries(options.params).map(([key, value]) => [key, String(value)])
        ).toString()
        url += `?${queryString}`
      }

      // Prepare fetch options
      const fetchOptions: RequestInit = {
        ...options,
        headers: {
          'Content-Type': 'application/json',
          ...options.headers,
        },
      }

      // Add authorization header if token exists
      if (authStore.token) {
        fetchOptions.headers = {
          ...fetchOptions.headers,
          Authorization: `Bearer ${authStore.token}`,
        }
      }

      // Log request (DEBUG level)
      logger.debug(`Request: ${options.method || 'GET'} ${endpoint}`, {
        method: options.method || 'GET',
        params: options.params,
        hasBody: !!options.body,
      })

      const response = await fetch(url, fetchOptions)

      // Log response (DEBUG level)
      logger.debug(`Response: ${response.status} ${endpoint}`, {
        status: response.status,
        ok: response.ok,
      })

      // Handle 401 Unauthorized - auto logout
      if (response.status === 401) {
        logger.warn('Unauthorized request, logging out', { endpoint })
        authStore.logout()
        navigateTo('/login')
        return {
          success: false,
          error: 'Session expired. Please login again.',
        }
      }

      // Handle blob responses (file downloads)
      const headers = options.headers as Record<string, string> | undefined
      if (headers?.['Accept'] === 'application/octet-stream') {
        const blob = await response.blob()
        return {
          success: true,
          data: blob as T,
        }
      }

      const data = await response.json()

      if (!response.ok) {
        logger.warn(`API Error: ${response.status} ${endpoint}`, {
          status: response.status,
          error: data.error,
        })
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      return data as ApiResponse<T>
    } catch (error) {
      // Log error
      logger.error(`API Error: ${endpoint}`, {
        error,
        method: options.method || 'GET',
      })

      // Return user-friendly error message
      let errorMessage = 'Network error'

      if (error instanceof Error) {
        if (error.message.includes('fetch') || error.message.includes('Failed to fetch')) {
          errorMessage = 'Unable to connect to server. Please check your internet connection.'
        } else {
          errorMessage = error.message
        }
      }

      return {
        success: false,
        error: errorMessage,
      }
    }
  }

  return {
    provide: {
      api,
    },
  }
})
