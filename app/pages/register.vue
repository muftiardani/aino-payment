<template>
  <NuxtLayout name="auth">
    <Card class="w-full max-w-md">
      <template #header>
        <h2
          class="text-3xl font-bold text-center bg-clip-text text-transparent bg-gradient-to-r from-primary to-violet-600 mb-2"
        >
          Create Account
        </h2>
        <p class="text-sm text-muted-foreground text-center">
          Sign up to start managing your payments
        </p>
      </template>

      <form @submit.prevent="handleRegister" class="space-y-4">
        <Input
          id="full_name"
          v-model="form.full_name"
          type="text"
          label="Full Name"
          placeholder="John Doe"
          required
          :error="errors.full_name"
          class="animate-fade-up"
          style="animation-delay: 100ms"
        >
          <template #prefix>
            <IconUser class="w-5 h-5" />
          </template>
        </Input>

        <Input
          id="email"
          v-model="form.email"
          type="email"
          label="Email"
          placeholder="your@email.com"
          required
          :error="errors.email"
          class="animate-fade-up"
          style="animation-delay: 200ms"
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
          hint="Minimum 6 characters"
          class="animate-fade-up"
          style="animation-delay: 300ms"
        >
          <template #prefix>
            <IconLock class="w-5 h-5" />
          </template>
        </Input>

        <Input
          id="confirm_password"
          v-model="form.confirm_password"
          type="password"
          label="Confirm Password"
          placeholder="••••••••"
          required
          :error="errors.confirm_password"
          class="animate-fade-up"
          style="animation-delay: 400ms"
        >
          <template #prefix>
            <IconLock class="w-5 h-5" />
          </template>
        </Input>

        <Button
          type="submit"
          variant="primary"
          size="lg"
          class="w-full animate-fade-up"
          style="animation-delay: 500ms"
          :loading="uiStore.loading"
        >
          Register
        </Button>
      </form>

      <template #footer>
        <p class="text-sm text-center text-muted-foreground">
          Already have an account?
          <NuxtLink to="/login" class="text-primary hover:underline font-medium">
            Login here
          </NuxtLink>
        </p>
      </template>
    </Card>
  </NuxtLayout>
</template>

<script setup lang="ts">
import IconUser from '~/components/icons/IconUser.vue'
import IconEmail from '~/components/icons/IconEmail.vue'
import IconLock from '~/components/icons/IconLock.vue'
import { registerSchema } from '~/utils/validation-schemas'

definePageMeta({
  middleware: 'guest',
  layout: false,
})

const { register } = useAuth()
const uiStore = useUIStore()
const { errors, validate } = useFormValidation(registerSchema)

const form = reactive({
  full_name: '',
  email: '',
  password: '',
  confirm_password: '',
})

const handleRegister = async () => {
  // Validate form
  if (!validate(form)) {
    return
  }

  await register(form.email, form.password, form.full_name)
}
</script>
