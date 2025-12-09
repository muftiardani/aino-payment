export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();

  // Load auth from storage if not already loaded
  if (!authStore.isAuthenticated && process.client) {
    authStore.loadFromStorage();
  }

  // If not authenticated, redirect to login
  if (!authStore.isAuthenticated) {
    return navigateTo("/login");
  }
});
