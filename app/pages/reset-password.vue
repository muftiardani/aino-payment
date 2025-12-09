<template>
  <NuxtLayout name="auth">
    <div class="w-full max-w-md space-y-8">
      <!-- Logo/Header -->
      <div class="text-center">
        <h1
          class="text-4xl font-extrabold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent mb-2"
        >
          AinoPay
        </h1>
        <h2 class="text-2xl font-bold text-foreground">Set New Password</h2>
        <p class="mt-2 text-muted-foreground">Enter your new password below</p>
      </div>

      <Card class="p-8 backdrop-blur-sm bg-card/80 border-muted/20 shadow-xl">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <Input
            id="password"
            v-model="password"
            type="password"
            label="New Password"
            placeholder="••••••••"
            required
            :error="errors.password"
            class="animate-fade-up"
            style="animation-delay: 100ms"
          >
            <template #prefix>
              <IconLock class="w-5 h-5 text-muted-foreground" />
            </template>
          </Input>

          <Input
            id="confirmPassword"
            v-model="confirmPassword"
            type="password"
            label="Confirm Password"
            placeholder="••••••••"
            required
            :error="errors.confirmPassword"
            class="animate-fade-up"
            style="animation-delay: 200ms"
          >
            <template #prefix>
              <IconLock class="w-5 h-5 text-muted-foreground" />
            </template>
          </Input>

          <Button
            type="submit"
            variant="primary"
            size="lg"
            class="w-full animate-fade-up"
            style="animation-delay: 300ms"
            :loading="loading"
          >
            Reset Password
          </Button>

          <div class="text-center text-sm animate-fade-up" style="animation-delay: 400ms">
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
  </NuxtLayout>
</template>

<script setup lang="ts">
import IconLock from '~/components/icons/IconLock.vue'

definePageMeta({
  layout: false,
  middleware: 'guest',
})

const route = useRoute()
const authStore = useAuthStore()
const uiStore = useUIStore()

const password = ref('')
const confirmPassword = ref('')
const loading = ref(false)
const token = route.query.token as string

const errors = reactive({
  password: '',
  confirmPassword: '',
})

onMounted(() => {
  if (!token) {
    uiStore.showToast('Invalid reset link', 'error')
    navigateTo('/forgot-password')
  }
})

const handleSubmit = async () => {
  // Validate
  let hasError = false
  if (password.value.length < 6) {
    errors.password = 'Password must be at least 6 characters'
    hasError = true
  } else {
    errors.password = ''
  }

  if (password.value !== confirmPassword.value) {
    errors.confirmPassword = 'Passwords do not match'
    hasError = true
  } else {
    errors.confirmPassword = ''
  }

  if (hasError) return

  loading.value = true

  try {
    const response = await authStore.resetPassword({
      token,
      new_password: password.value,
    })

    if (response.success) {
      uiStore.showToast('Password reset successfully! Please login.', 'success')
      navigateTo('/login')
    } else {
      uiStore.showToast(response.error || 'Failed to reset password', 'error')
    }
  } catch (err) {
    console.error(err)
    uiStore.showToast('An unexpected error occurred', 'error')
  } finally {
    loading.value = false
  }
}
</script>
