export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();

  // Load auth from storage
  if (process.client) {
    authStore.loadFromStorage();
  }

  // If already authenticated, redirect to home
  if (authStore.isAuthenticated) {
    return navigateTo("/");
  }
});
