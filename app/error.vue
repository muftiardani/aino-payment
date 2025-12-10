<template>
  <NuxtLayout name="auth">
    <div class="min-h-screen flex items-center justify-center p-4">
      <Card class="w-full max-w-md text-center">
        <!-- Error Icon -->
        <div class="mb-6">
          <div
            v-if="is404"
            class="inline-flex items-center justify-center w-20 h-20 rounded-full bg-primary/10 text-primary mb-4"
          >
            <svg class="w-10 h-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
          </div>
          <div
            v-else
            class="inline-flex items-center justify-center w-20 h-20 rounded-full bg-destructive/10 text-destructive mb-4"
          >
            <svg class="w-10 h-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
              />
            </svg>
          </div>

          <!-- Error Code -->
          <h1 class="text-6xl font-bold text-foreground mb-2">
            {{ error?.statusCode || 500 }}
          </h1>

          <!-- Error Title -->
          <h2 class="text-2xl font-semibold text-foreground mb-2">
            {{ errorTitle }}
          </h2>

          <!-- Error Message -->
          <p class="text-muted-foreground mb-6">
            {{ errorMessage }}
          </p>
        </div>

        <!-- Action Buttons -->
        <div class="flex flex-col sm:flex-row gap-3 justify-center">
          <Button variant="primary" @click="handleClearError">
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
              />
            </svg>
            Go Home
          </Button>

          <Button v-if="!is404" variant="outline" @click="handleRetry">
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

          <Button v-else variant="outline" @click="goBack">
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M10 19l-7-7m0 0l7-7m-7 7h18"
              />
            </svg>
            Go Back
          </Button>
        </div>

        <!-- Debug Info (Development Only) -->
        <div v-if="isDev && error?.stack" class="mt-6 pt-6 border-t border-border">
          <details class="text-left">
            <summary
              class="text-sm font-medium text-muted-foreground cursor-pointer hover:text-foreground"
            >
              Error Details (Development)
            </summary>
            <pre class="mt-2 p-3 bg-muted rounded-md text-xs overflow-auto max-h-40">{{
              error.stack
            }}</pre>
          </details>
        </div>
      </Card>
    </div>
  </NuxtLayout>
</template>

<script setup lang="ts">
const props = defineProps<{
  error: {
    statusCode?: number
    message?: string
    stack?: string
  }
}>()

const isDev = import.meta.dev
const router = useRouter()

const is404 = computed(() => props.error?.statusCode === 404)

const errorTitle = computed(() => {
  const code = props.error?.statusCode || 500

  const titles: Record<number, string> = {
    404: 'Page Not Found',
    401: 'Unauthorized',
    403: 'Forbidden',
    500: 'Internal Server Error',
    502: 'Bad Gateway',
    503: 'Service Unavailable',
  }

  return titles[code] || 'Something Went Wrong'
})

const errorMessage = computed(() => {
  if (props.error?.message) {
    return props.error.message
  }

  const code = props.error?.statusCode || 500

  const messages: Record<number, string> = {
    404: "The page you're looking for doesn't exist or has been moved.",
    401: 'You need to be logged in to access this page.',
    403: "You don't have permission to access this resource.",
    500: "We're working on fixing this. Please try again later.",
    502: 'The server is temporarily unavailable. Please try again in a moment.',
    503: 'The service is currently under maintenance. Please try again later.',
  }

  return messages[code] || 'An unexpected error occurred. Please try again.'
})

const handleClearError = () => {
  clearError({ redirect: '/' })
}

const handleRetry = () => {
  // Reload the current page
  window.location.reload()
}

const goBack = () => {
  router.back()
}
</script>
