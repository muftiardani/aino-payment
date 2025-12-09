<template>
  <div
    class="min-h-screen bg-gradient-to-br from-primary/10 via-background to-secondary/10 flex items-center justify-center p-4"
  >
    <slot />
    <Toast />
  </div>
</template>

<script setup lang="ts">
const uiStore = useUIStore()

onMounted(() => {
  uiStore.loadDarkMode()
})
</script>
