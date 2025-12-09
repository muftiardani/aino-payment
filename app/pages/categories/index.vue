<template>
  <div>
    <div class="mb-6">
      <h1
        class="text-3xl font-bold bg-gradient-to-r from-primary to-purple-600 bg-clip-text text-transparent"
      >
        Categories
      </h1>
      <p class="text-muted-foreground mt-1">Browse all payment categories</p>
    </div>

    <Card :padding="false">
      <SkeletonCard v-if="loading" :count="3" />

      <EmptyState
        v-else-if="categories.length === 0"
        title="No categories found"
        description="There are no payment categories available at the moment."
      >
        <template #icon>
          <svg
            class="w-16 h-16 text-muted-foreground/50"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"
            />
          </svg>
        </template>
      </EmptyState>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-6">
        <div
          v-for="(category, index) in categories"
          :key="category.id"
          class="group p-6 rounded-xl border-2 border-border hover:border-primary hover:shadow-xl hover:shadow-primary/10 transition-all duration-300 cursor-pointer bg-gradient-to-br from-card to-card/50 fade-in-card"
          :style="{ animationDelay: `${index * 100}ms` }"
        >
          <div class="flex items-start gap-4">
            <div
              class="p-3 rounded-full bg-gradient-to-br from-primary/20 to-primary/10 group-hover:from-primary/30 group-hover:to-primary/20 group-hover:scale-110 transition-all duration-300"
            >
              <svg
                class="w-6 h-6 text-primary"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"
                />
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="font-bold text-lg mb-1 group-hover:text-primary transition-colors">
                {{ category.name }}
              </h3>
              <p class="text-sm text-muted-foreground leading-relaxed">
                {{ category.description }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </Card>
  </div>
</template>

<script setup lang="ts">
import type { Category } from '~/types/payment'

definePageMeta({
  middleware: 'auth',
})

const { getCategories } = usePayment()

const categories = ref<Category[]>([])
const loading = ref(true)

onMounted(async () => {
  const response = await getCategories()
  if (response.success && response.data) {
    categories.value = response.data
  }
  loading.value = false
})
</script>

<style scoped>
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.fade-in-card {
  animation: fadeInUp 0.5s ease-out forwards;
  opacity: 0;
}
</style>
