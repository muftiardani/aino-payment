import { defineStore } from 'pinia'
import type {
  User,
  LoginRequest,
  RegisterRequest,
  AuthResponse,
  ForgotPasswordRequest,
  ResetPasswordRequest,
} from '~/types/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null,
    token: null as string | null,
    refreshToken: null as string | null,
    tokenExpiresAt: null as number | null, // timestamp
    isAuthenticated: false,
  }),

  actions: {
    async register(data: RegisterRequest) {
      const { $api } = useNuxtApp()
      const response = await $api<AuthResponse>('/auth/register', {
        method: 'POST',
        body: JSON.stringify(data),
      })

      if (response.data) {
        this.setAuth(response.data)
      }

      return response
    },

    async login(data: LoginRequest) {
      const { $api } = useNuxtApp()
      const response = await $api<AuthResponse>('/auth/login', {
        method: 'POST',
        body: JSON.stringify(data),
      })

      if (response.data) {
        this.setAuth(response.data)
      }

      return response
    },

    async fetchUser() {
      const { $api } = useNuxtApp()
      const response = await $api<User>('/auth/me')

      if (response.data) {
        this.user = response.data
      }

      return response
    },

    async forgotPassword(data: ForgotPasswordRequest) {
      const { $api } = useNuxtApp()
      return await $api('/auth/forgot-password', {
        method: 'POST',
        body: JSON.stringify(data),
      })
    },

    async resetPassword(data: ResetPasswordRequest) {
      const { $api } = useNuxtApp()
      return await $api('/auth/reset-password', {
        method: 'POST',
        body: JSON.stringify(data),
      })
    },

    setAuth(authData: AuthResponse) {
      this.user = authData.user
      this.token = authData.token
      this.refreshToken = authData.refresh_token
      this.isAuthenticated = true

      // Calculate expiry timestamp
      const expiresAt = Date.now() + authData.expires_in * 1000
      this.tokenExpiresAt = expiresAt

      // Save to localStorage
      if (import.meta.client) {
        localStorage.setItem('token', authData.token)
        localStorage.setItem('refresh_token', authData.refresh_token)
        localStorage.setItem('token_expires_at', expiresAt.toString())
        localStorage.setItem('user', JSON.stringify(authData.user))
      }
    },

    loadFromStorage() {
      if (import.meta.client) {
        const token = localStorage.getItem('token')
        const refreshToken = localStorage.getItem('refresh_token')
        const tokenExpiresAt = localStorage.getItem('token_expires_at')
        const userStr = localStorage.getItem('user')

        if (token && userStr) {
          this.token = token
          this.refreshToken = refreshToken
          this.tokenExpiresAt = tokenExpiresAt ? Number.parseInt(tokenExpiresAt) : null
          this.user = JSON.parse(userStr)
          this.isAuthenticated = true
        }
      }
    },

    logout() {
      this.user = null
      this.token = null
      this.refreshToken = null
      this.tokenExpiresAt = null
      this.isAuthenticated = false

      if (import.meta.client) {
        localStorage.removeItem('token')
        localStorage.removeItem('refresh_token')
        localStorage.removeItem('token_expires_at')
        localStorage.removeItem('user')
      }

      navigateTo('/login')
    },
  },

  getters: {
    isAdmin: state => state.user?.role === 'admin',
  },
})
