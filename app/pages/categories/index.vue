<template>
  <div>
    <h1 class="text-3xl font-bold mb-6">Categories</h1>

    <Card>
      <Loading v-if="loading" :loading="loading" />

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div
          v-for="category in categories"
          :key="category.id"
          class="p-4 rounded-lg border hover:bg-accent transition-colors"
        >
          <h3 class="font-bold text-lg mb-1">{{ category.name }}</h3>
          <p class="text-sm text-muted-foreground">
            {{ category.description }}
          </p>
        </div>
      </div>
    </Card>
  </div>
</template>

<script setup lang="ts">
import type { Category } from "~/types/payment";

definePageMeta({
  middleware: "auth",
});

const { getCategories } = usePayment();

const categories = ref<Category[]>([]);
const loading = ref(true);

onMounted(async () => {
  const response = await getCategories();
  if (response.success && response.data) {
    categories.value = response.data;
  }
  loading.value = false;
});
</script>
