export default defineNuxtRouteMiddleware((_to, _from) => {
  const authStore = useAuthStore()

  // Load auth from storage
  if (import.meta.client) {
    authStore.loadFromStorage()
  }

  // If already authenticated, redirect to home
  if (authStore.isAuthenticated) {
    return navigateTo('/')
  }
})
