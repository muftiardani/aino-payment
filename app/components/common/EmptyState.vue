<template>
  <div class="flex flex-col items-center justify-center py-12 px-4 text-center">
    <!-- Icon -->
    <div class="mb-6 relative">
      <div class="absolute inset-0 bg-primary/10 rounded-full blur-2xl"></div>
      <div class="relative bg-gradient-to-br from-primary/20 to-secondary/20 p-8 rounded-full">
        <slot name="icon">
          <svg
            class="w-16 h-16 text-muted-foreground/50"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"
            />
          </svg>
        </slot>
      </div>
    </div>

    <!-- Title -->
    <h3 class="text-xl font-bold mb-2 text-foreground">
      {{ title }}
    </h3>

    <!-- Description -->
    <p class="text-muted-foreground mb-6 max-w-md">
      {{ description }}
    </p>

    <!-- Action Button -->
    <slot name="action">
      <Button v-if="actionText" @click="$emit('action')" variant="primary">
        {{ actionText }}
      </Button>
    </slot>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string
  description: string
  actionText?: string
}>()

defineEmits<{
  action: []
}>()
</script>
