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
            :error="errors.email"
            class="animate-fade-up"
            style="animation-delay: 100ms"
          >
            <template #prefix>
              <IconEmail class="w-5 h-5 text-muted-foreground" />
            </template>
          </Input>

          <Button
            type="submit"
            variant="primary"
            size="lg"
            class="w-full animate-fade-up"
            style="animation-delay: 200ms"
            :loading="loading"
          >
            Send Reset Link
          </Button>

          <div class="text-center text-sm animate-fade-up" style="animation-delay: 300ms">
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
import IconEmail from '~/components/icons/IconEmail.vue'
import { forgotPasswordSchema } from '~/utils/validation-schemas'

definePageMeta({
  layout: false,
  middleware: 'guest',
})

const authStore = useAuthStore()
const uiStore = useUIStore()
const { errors, validate } = useFormValidation(forgotPasswordSchema)

const email = ref('')
const loading = ref(false)

const handleSubmit = async () => {
  // Validate form
  if (!validate({ email: email.value })) {
    return
  }

  loading.value = true

  try {
    const response = await authStore.forgotPassword({ email: email.value })

    if (response.success) {
      uiStore.showToast('Reset link sent! Please check your email.', 'success')
      email.value = ''
    } else {
      uiStore.showToast(response.error || 'Failed to send reset link', 'error')
    }
  } catch (err) {
    console.error(err)
    uiStore.showToast('An unexpected error occurred', 'error')
  } finally {
    loading.value = false
  }
}
</script>
