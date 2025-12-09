export default defineNuxtRouteMiddleware((_to, _from) => {
  const authStore = useAuthStore()

  // Load auth from storage if not already loaded
  if (!authStore.isAuthenticated && import.meta.client) {
    authStore.loadFromStorage()
  }

  // If not authenticated, redirect to login
  if (!authStore.isAuthenticated) {
    return navigateTo('/login')
  }
})
