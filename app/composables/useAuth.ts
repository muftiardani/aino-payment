export const useAuth = () => {
  const authStore = useAuthStore()
  const router = useRouter()
  const uiStore = useUIStore()

  const login = async (email: string, password: string) => {
    uiStore.setLoading(true)
    try {
      const response = await authStore.login({ email, password })

      if (response.success) {
        uiStore.showToast('Login successful!', 'success')
        router.push('/')
      } else {
        uiStore.showToast(response.error || 'Login failed', 'error')
      }

      return response
    } finally {
      uiStore.setLoading(false)
    }
  }

  const register = async (email: string, password: string, full_name: string) => {
    uiStore.setLoading(true)
    try {
      const response = await authStore.register({ email, password, full_name })

      if (response.success) {
        uiStore.showToast('Registration successful!', 'success')
        router.push('/')
      } else {
        uiStore.showToast(response.error || 'Registration failed', 'error')
      }

      return response
    } finally {
      uiStore.setLoading(false)
    }
  }

  const logout = () => {
    authStore.logout()
    uiStore.showToast('Logged out successfully', 'info')
  }

  return {
    user: computed(() => authStore.user),
    isAuthenticated: computed(() => authStore.isAuthenticated),
    isAdmin: computed(() => authStore.isAdmin),
    login,
    register,
    logout,
  }
}
