<template>
  <div
    class="min-h-screen bg-background flex items-center justify-center p-4 relative overflow-hidden"
  >
    <!-- Decorative Background -->
    <div class="absolute inset-0 z-0">
      <div
        class="absolute top-0 -left-4 w-72 h-72 bg-primary/30 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob"
      ></div>
      <div
        class="absolute top-0 -right-4 w-72 h-72 bg-violet-500/30 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"
      ></div>
      <div
        class="absolute -bottom-8 left-20 w-72 h-72 bg-pink-500/30 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"
      ></div>
    </div>

    <div class="z-10 w-full max-w-md animate-fade-up">
      <slot />
    </div>

    <Toast />
  </div>
</template>

<script setup lang="ts">
const uiStore = useUIStore()

onMounted(() => {
  uiStore.loadDarkMode()
})
</script>
