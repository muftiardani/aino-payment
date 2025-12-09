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
const uiStore = useUIStore();

onMounted(() => {
  uiStore.loadDarkMode();
});
</script>
