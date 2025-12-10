<template>
  <div class="flex flex-col items-center justify-center p-6 text-center" :class="sizeClasses">
    <!-- Error Icon -->
    <div
      class="inline-flex items-center justify-center rounded-full bg-destructive/10 text-destructive mb-4"
      :class="iconSizeClasses"
    >
      <svg class="w-1/2 h-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
    </div>

    <!-- Error Message -->
    <h3 class="font-semibold text-foreground mb-2" :class="titleSizeClasses">
      {{ displayTitle }}
    </h3>
    <p class="text-muted-foreground mb-4" :class="textSizeClasses">
      {{ displayMessage }}
    </p>

    <!-- Retry Button -->
    <Button v-if="retry" variant="outline" size="sm" @click="handleRetry">
      <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
        />
      </svg>
      Try Again
    </Button>
  </div>
</template>

<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    error?: unknown
    title?: string
    message?: string
    retry?: () => void | Promise<void>
    size?: 'sm' | 'md' | 'lg'
  }>(),
  {
    title: 'Something went wrong',
    message: 'An error occurred while loading this content.',
    size: 'md',
  }
)

const { getErrorMessage } = useErrorHandler()

// Get display message from props or error
const displayTitle = computed(() => props.title || 'Something went wrong')
const displayMessage = computed(() => {
  if (props.message) return props.message
  if (props.error) return getErrorMessage(props.error)
  return 'An error occurred while loading this content.'
})

const sizeClasses = computed(() => {
  const sizes = {
    sm: 'min-h-[200px]',
    md: 'min-h-[300px]',
    lg: 'min-h-[400px]',
  }
  return sizes[props.size]
})

const iconSizeClasses = computed(() => {
  const sizes = {
    sm: 'w-12 h-12',
    md: 'w-16 h-16',
    lg: 'w-20 h-20',
  }
  return sizes[props.size]
})

const titleSizeClasses = computed(() => {
  const sizes = {
    sm: 'text-base',
    md: 'text-lg',
    lg: 'text-xl',
  }
  return sizes[props.size]
})

const textSizeClasses = computed(() => {
  const sizes = {
    sm: 'text-sm',
    md: 'text-base',
    lg: 'text-lg',
  }
  return sizes[props.size]
})

const handleRetry = async () => {
  if (props.retry) {
    await props.retry()
  }
}
</script>
