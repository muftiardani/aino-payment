<template>
  <NuxtLayout name="auth">
    <Card class="w-full max-w-md">
      <template #header>
        <h2 class="text-2xl font-bold text-center">Create Account</h2>
        <p class="text-sm text-muted-foreground text-center mt-1">
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
        />

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
          hint="Minimum 6 characters"
        />

        <Input
          id="confirm_password"
          v-model="form.confirm_password"
          type="password"
          label="Confirm Password"
          placeholder="••••••••"
          required
          :error="errors.confirm_password"
        />

        <Button
          type="submit"
          variant="primary"
          size="lg"
          class="w-full"
          :loading="uiStore.loading"
        >
          Register
        </Button>
      </form>

      <template #footer>
        <p class="text-sm text-center text-muted-foreground">
          Already have an account?
          <NuxtLink
            to="/login"
            class="text-primary hover:underline font-medium"
          >
            Login here
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

const { register } = useAuth();
const uiStore = useUIStore();

const form = reactive({
  full_name: "",
  email: "",
  password: "",
  confirm_password: "",
});

const errors = reactive({
  full_name: "",
  email: "",
  password: "",
  confirm_password: "",
});

const handleRegister = async () => {
  // Reset errors
  Object.keys(errors).forEach((key) => {
    errors[key as keyof typeof errors] = "";
  });

  // Validate
  if (!form.full_name) {
    errors.full_name = "Full name is required";
    return;
  }
  if (!form.email) {
    errors.email = "Email is required";
    return;
  }
  if (!form.password) {
    errors.password = "Password is required";
    return;
  }
  if (form.password.length < 6) {
    errors.password = "Password must be at least 6 characters";
    return;
  }
  if (form.password !== form.confirm_password) {
    errors.confirm_password = "Passwords do not match";
    return;
  }

  await register(form.email, form.password, form.full_name);
};
</script>
