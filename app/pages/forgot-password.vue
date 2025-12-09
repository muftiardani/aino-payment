<template>
  <div class="min-h-screen flex items-center justify-center bg-background px-4">
    <div class="w-full max-w-md space-y-8">
      <!-- Logo/Header -->
      <div class="text-center">
        <h1
          class="text-4xl font-extrabold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent mb-2"
        >
          AinoPay
        </h1>
        <h2 class="text-2xl font-bold text-foreground">Reset Password</h2>
        <p class="mt-2 text-muted-foreground">Enter your email to receive password reset link</p>
      </div>

      <Card class="p-8 backdrop-blur-sm bg-card/80 border-muted/20 shadow-xl">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <Input
            id="email"
            v-model="email"
            type="email"
            label="Email Address"
            placeholder="name@company.com"
            required
            :error="error"
          >
            <template #prefix>
              <IconEmail class="w-5 h-5 text-muted-foreground" />
            </template>
          </Input>

          <Button type="submit" variant="primary" size="lg" class="w-full" :loading="loading">
            Send Reset Link
          </Button>

          <div class="text-center text-sm">
            <NuxtLink
              to="/login"
              class="font-medium text-primary hover:text-primary/90 transition-colors"
            >
              Back to Login
            </NuxtLink>
          </div>
        </form>
      </Card>
    </div>
  </div>
</template>

<script setup lang="ts">
import IconEmail from '~/components/icons/IconEmail.vue'

definePageMeta({
  layout: false,
  middleware: 'guest',
})

const authStore = useAuthStore()
const uiStore = useUIStore()

const email = ref('')
const loading = ref(false)
const error = ref('')

const handleSubmit = async () => {
  if (!email.value) return

  loading.value = true
  error.value = ''

  try {
    const response = await authStore.forgotPassword({ email: email.value })

    if (response.success) {
      uiStore.showToast('Reset link sent! Please check your email.', 'success')
      // Optional: navigate to login or show success message
      email.value = ''
    } else {
      error.value = response.error || 'Failed to send reset link'
    }
  } catch (err) {
    console.error(err)
    error.value = 'An unexpected error occurred'
  } finally {
    loading.value = false
  }
}
</script>
