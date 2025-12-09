import type { AuthResponse } from '~/types/auth'

export const useTokenRefresh = () => {
  const authStore = useAuthStore()
  const { $api } = useNuxtApp()

  let refreshTimer: ReturnType<typeof setTimeout> | null = null

  /**
   * Schedule token refresh before expiry
   * @param expiresIn - Token expiry time in seconds
   */
  const scheduleRefresh = (expiresIn: number) => {
    // Clear existing timer
    if (refreshTimer) {
      clearTimeout(refreshTimer)
    }

    // Refresh 5 minutes (300 seconds) before expiry
    const refreshTime = Math.max((expiresIn - 300) * 1000, 60000) // At least 1 minute

    refreshTimer = setTimeout(async () => {
      await refreshToken()
    }, refreshTime)

    console.log(`Token refresh scheduled in ${refreshTime / 1000} seconds`)
  }

  /**
   * Refresh the access token using refresh token
   */
  const refreshToken = async () => {
    const refreshTokenValue = authStore.refreshToken

    if (!refreshTokenValue) {
      console.warn('No refresh token available')
      authStore.logout()
      return
    }

    try {
      const response = await $api<AuthResponse>('/auth/refresh', {
        method: 'POST',
        body: JSON.stringify({ refresh_token: refreshTokenValue }),
      })

      if (response.success && response.data) {
        // Update auth store with new token
        authStore.setAuth(response.data)

        // Schedule next refresh
        scheduleRefresh(response.data.expires_in)

        console.log('Token refreshed successfully')
      } else {
        console.error('Token refresh failed:', response.error)
        authStore.logout()
      }
    } catch (error) {
      console.error('Token refresh error:', error)
      authStore.logout()
    }
  }

  /**
   * Start auto-refresh based on current token expiry
   */
  const startAutoRefresh = () => {
    const expiresAt = authStore.tokenExpiresAt

    if (!expiresAt) {
      console.warn('No token expiry time available')
      return
    }

    const now = Date.now()
    const expiresIn = Math.floor((expiresAt - now) / 1000)

    if (expiresIn <= 0) {
      // Token already expired, refresh immediately
      refreshToken()
    } else {
      // Schedule refresh
      scheduleRefresh(expiresIn)
    }
  }

  /**
   * Stop auto-refresh (cleanup)
   */
  const stopAutoRefresh = () => {
    if (refreshTimer) {
      clearTimeout(refreshTimer)
      refreshTimer = null
    }
  }

  return {
    scheduleRefresh,
    refreshToken,
    startAutoRefresh,
    stopAutoRefresh,
  }
}
