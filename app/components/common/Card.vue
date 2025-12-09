<template>
  <div class="glass-card shadow-sm hover:shadow-md transition-shadow duration-300 overflow-hidden">
    <div
      v-if="$slots.header || title"
      class="px-6 py-4 border-b border-white/10 dark:border-white/5"
    >
      <slot name="header">
        <h3 v-if="title" class="text-lg font-semibold leading-none tracking-tight">
          {{ title }}
        </h3>
      </slot>
    </div>
    <div class="p-6 pt-0" :class="{ 'pt-6': !$slots.header && !title }">
      <slot />
    </div>
    <div v-if="$slots.footer" class="flex items-center p-6 pt-0">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup lang="ts">
withDefaults(
  defineProps<{
    padding?: boolean
    title?: string
  }>(),
  {
    padding: true,
    title: undefined,
  }
)
</script>
