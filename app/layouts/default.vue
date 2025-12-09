<template>
  <div class="min-h-screen bg-background">
    <Navbar />
    <Sidebar />

    <main class="lg:ml-64 mt-16 p-4 sm:p-6 lg:p-8">
      <slot />
    </main>

    <Toast />
  </div>
</template>

<script setup lang="ts">
import { useTokenRefresh } from '~/composables/useTokenRefresh'

const uiStore = useUIStore()
const authStore = useAuthStore()

onMounted(() => {
  uiStore.loadDarkMode()

  // Start auto-refresh if user is authenticated
  if (authStore.isAuthenticated && authStore.tokenExpiresAt) {
    const { startAutoRefresh } = useTokenRefresh()
    startAutoRefresh()
  }
})

onUnmounted(() => {
  // Cleanup auto-refresh timer
  if (authStore.isAuthenticated) {
    const { stopAutoRefresh } = useTokenRefresh()
    stopAutoRefresh()
  }
})
</script>
