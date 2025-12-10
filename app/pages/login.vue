<template>
  <NuxtLayout name="auth">
    <Card class="w-full max-w-md">
      <template #header>
        <h2
          class="text-3xl font-bold text-center bg-clip-text text-transparent bg-gradient-to-r from-primary to-violet-600 mb-2"
        >
          Login to AinoPay
        </h2>
        <p class="text-sm text-muted-foreground text-center">
          Enter your credentials to access your account
        </p>
      </template>

      <form @submit.prevent="handleLogin" class="space-y-4">
        <Input
          id="email"
          v-model="form.email"
          type="email"
          label="Email"
          placeholder="your@email.com"
          required
          :error="errors.email"
          class="animate-fade-up"
          style="animation-delay: 100ms"
        >
          <template #prefix>
            <IconEmail class="w-5 h-5" />
          </template>
        </Input>

        <Input
          id="password"
          v-model="form.password"
          type="password"
          label="Password"
          placeholder="••••••••"
          required
          :error="errors.password"
          class="animate-fade-up"
          style="animation-delay: 200ms"
        >
          <template #prefix>
            <IconLock class="w-5 h-5" />
          </template>
        </Input>

        <div class="animate-fade-up" style="animation-delay: 300ms">
          <NuxtLink
            to="/forgot-password"
            class="text-sm font-medium text-primary hover:underline block text-right"
          >
            Forgot password?
          </NuxtLink>
        </div>

        <Button
          type="submit"
          variant="primary"
          size="lg"
          class="w-full animate-fade-up"
          style="animation-delay: 400ms"
          :loading="uiStore.loading"
        >
          Login
        </Button>

        <!-- Social Login -->
        <div class="relative animate-fade-up" style="animation-delay: 500ms">
          <div class="absolute inset-0 flex items-center">
            <span class="w-full border-t border-muted"></span>
          </div>
          <div class="relative flex justify-center text-xs uppercase">
            <span class="bg-background px-2 text-muted-foreground">Or continue with</span>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4 animate-fade-up" style="animation-delay: 600ms">
          <Button variant="outline" class="w-full" type="button">
            <IconGoogle class="mr-2 h-4 w-4" />
            Google
          </Button>
          <Button variant="outline" class="w-full" type="button">
            <IconGithub class="mr-2 h-4 w-4" />
            Github
          </Button>
        </div>
      </form>

      <template #footer>
        <p class="text-sm text-center text-muted-foreground">
          Don't have an account?
          <NuxtLink to="/register" class="text-primary hover:underline font-medium">
            Register here
          </NuxtLink>
        </p>
      </template>
    </Card>
  </NuxtLayout>
</template>

<script setup lang="ts">
import IconEmail from '~/components/icons/IconEmail.vue'
import IconLock from '~/components/icons/IconLock.vue'
import IconGoogle from '~/components/icons/IconGoogle.vue'
import IconGithub from '~/components/icons/IconGithub.vue'
import { loginSchema } from '~/utils/validation-schemas'

definePageMeta({
  middleware: 'guest',
  layout: false,
})

const { login } = useAuth()
const uiStore = useUIStore()
const { errors, validate } = useFormValidation(loginSchema)

const form = reactive({
  email: '',
  password: '',
})

const handleLogin = async () => {
  // Validate form
  if (!validate(form)) {
    return
  }

  await login(form.email, form.password)
}
</script>
