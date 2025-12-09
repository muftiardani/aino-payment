<template>
  <div :class="cardClasses">
    <div v-if="$slots.header" class="p-6 border-b">
      <slot name="header" />
    </div>
    <div :class="contentClasses">
      <slot />
    </div>
    <div v-if="$slots.footer" class="p-6 border-t">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    padding?: boolean
  }>(),
  {
    padding: true,
  }
)

const cardClasses = 'card'
const contentClasses = computed(() => (props.padding ? 'p-6' : ''))
</script>
