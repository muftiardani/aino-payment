<template>
  <NuxtLayout name="auth">
    <Card class="w-full max-w-md">
      <template #header>
        <h2 class="text-2xl font-bold text-center">Login to AinoPay</h2>
        <p class="text-sm text-muted-foreground text-center mt-1">
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
        />

        <Input
          id="password"
          v-model="form.password"
          type="password"
          label="Password"
          placeholder="••••••••"
          required
          :error="errors.password"
        />

        <Button
          type="submit"
          variant="primary"
          size="lg"
          class="w-full"
          :loading="uiStore.loading"
        >
          Login
        </Button>
      </form>

      <template #footer>
        <p class="text-sm text-center text-muted-foreground">
          Don't have an account?
          <NuxtLink
            to="/register"
            class="text-primary hover:underline font-medium"
          >
            Register here
          </NuxtLink>
        </p>
      </template>
    </Card>
  </NuxtLayout>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: "guest",
  layout: false,
});

const { login } = useAuth();
const uiStore = useUIStore();

const form = reactive({
  email: "",
  password: "",
});

const errors = reactive({
  email: "",
  password: "",
});

const handleLogin = async () => {
  // Reset errors
  errors.email = "";
  errors.password = "";

  // Validate
  if (!form.email) {
    errors.email = "Email is required";
    return;
  }
  if (!form.password) {
    errors.password = "Password is required";
    return;
  }

  await login(form.email, form.password);
};
</script>
