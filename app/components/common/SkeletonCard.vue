<template>
  <div class="space-y-4">
    <div v-for="i in count" :key="i" class="animate-pulse">
      <div class="card p-6">
        <div class="flex items-center justify-between">
          <div class="flex-1 space-y-3">
            <div class="h-4 bg-muted rounded w-3/4"></div>
            <div class="h-3 bg-muted rounded w-1/2"></div>
            <div class="h-3 bg-muted rounded w-1/4"></div>
          </div>
          <div class="ml-4">
            <div class="h-12 w-12 bg-muted rounded-full"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  count?: number
}>()
</script>
