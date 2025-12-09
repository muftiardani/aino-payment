<template>
  <div v-if="loading" class="flex items-center justify-center p-8">
    <div class="relative">
      <div class="w-12 h-12 rounded-full border-4 border-muted"></div>
      <div
        class="w-12 h-12 rounded-full border-4 border-primary border-t-transparent animate-spin absolute top-0 left-0"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  loading?: boolean;
}>();
</script>
