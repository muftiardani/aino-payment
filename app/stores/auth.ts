import { defineStore } from 'pinia'
import type { User, LoginRequest, RegisterRequest, AuthResponse } from '~/types/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null,
    token: null as string | null,
    isAuthenticated: false,
  }),

  actions: {
    async register(data: RegisterRequest) {
      const { $api } = useNuxtApp()
      const response = await $api<AuthResponse>('/auth/register', {
        method: 'POST',
        body: data,
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
        body: data,
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

    setAuth(authData: AuthResponse) {
      this.user = authData.user
      this.token = authData.token
      this.isAuthenticated = true

      // Save to localStorage
      if (process.client) {
        localStorage.setItem('token', authData.token)
        localStorage.setItem('user', JSON.stringify(authData.user))
      }
    },

    loadFromStorage() {
      if (process.client) {
        const token = localStorage.getItem('token')
        const userStr = localStorage.getItem('user')

        if (token && userStr) {
          this.token = token
          this.user = JSON.parse(userStr)
          this.isAuthenticated = true
        }
      }
    },

    logout() {
      this.user = null
      this.token = null
      this.isAuthenticated = false

      if (process.client) {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
      }

      navigateTo('/login')
    },
  },

  getters: {
    isAdmin: state => state.user?.role === 'admin',
  },
})
