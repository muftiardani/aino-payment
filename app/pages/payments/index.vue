<template>
  <div>
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1
          class="text-3xl font-bold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent"
        >
          Payments
        </h1>
        <p class="text-muted-foreground mt-1">Manage all your payment transactions</p>
      </div>
      <NuxtLink to="/payments/create">
        <Button variant="primary" class="group shadow-lg shadow-primary/20">
          <svg
            class="w-5 h-5 mr-2 group-hover:rotate-90 transition-transform duration-300"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 4v16m8-8H4"
            />
          </svg>
          New Payment
        </Button>
      </NuxtLink>
    </div>

    <!-- Filters -->
    <Card class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <!-- Search -->
        <div class="md:col-span-2">
          <div class="relative">
            <svg
              class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
            <input
              v-model="filters.search"
              type="text"
              placeholder="Search payments by description..."
              class="input pl-10 w-full"
              @input="debouncedSearch"
            />
          </div>
        </div>

        <!-- Status Filter -->
        <div>
          <select v-model="filters.status" class="input w-full" @change="loadPayments">
            <option value="">All Status</option>
            <option value="pending">Pending</option>
            <option value="completed">Completed</option>
            <option value="failed">Failed</option>
            <option value="refunded">Refunded</option>
          </select>
        </div>
      </div>

      <!-- Active Filters -->
      <div v-if="hasActiveFilters" class="mt-4 flex items-center gap-2 flex-wrap">
        <span class="text-sm text-muted-foreground">Active filters:</span>
        <Badge
          v-if="filters.search"
          variant="info"
          class="cursor-pointer"
          @click="removeSearchFilter"
        >
          Search: "{{ filters.search }}"
          <svg class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </Badge>
        <Badge
          v-if="filters.status"
          variant="info"
          class="cursor-pointer"
          @click="removeStatusFilter"
        >
          Status: {{ filters.status }}
          <svg class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </Badge>
        <button class="text-sm text-primary hover:underline" @click="clearFilters">
          Clear all
        </button>
      </div>
    </Card>

    <!-- Payments List -->
    <Card>
      <!-- Skeleton Loader -->
      <SkeletonCard v-if="loading" :count="5" />

      <!-- Empty State -->
      <EmptyState
        v-else-if="payments.length === 0"
        :title="hasActiveFilters ? 'No payments match your filters' : 'No payments found'"
        :description="
          hasActiveFilters
            ? 'Try adjusting your search or filters to find what you\'re looking for.'
            : 'You haven\'t created any payments yet. Click the button below to create your first payment transaction.'
        "
        :action-text="hasActiveFilters ? '' : 'Create First Payment'"
        @action="navigateTo('/payments/create')"
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
              d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
            />
          </svg>
        </template>
      </EmptyState>

      <!-- Payments Table -->
      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead class="border-b bg-muted/30">
            <tr class="text-left">
              <th class="p-4 font-semibold text-sm">Description</th>
              <th class="p-4 font-semibold text-sm">Amount</th>
              <th class="p-4 font-semibold text-sm">Category</th>
              <th class="p-4 font-semibold text-sm">Method</th>
              <th class="p-4 font-semibold text-sm">Status</th>
              <th class="p-4 font-semibold text-sm">Date</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(payment, index) in payments"
              :key="payment.id"
              class="border-b hover:bg-accent/50 transition-all duration-200 cursor-pointer group fade-in"
              :style="{ animationDelay: `${index * 50}ms` }"
              @click="navigateTo(`/payments/${payment.id}`)"
            >
              <td class="p-4">
                <div class="flex items-center gap-3">
                  <div
                    class="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 group-hover:scale-110 transition-all"
                  >
                    <svg
                      class="w-5 h-5 text-primary"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                      />
                    </svg>
                  </div>
                  <span class="font-medium">{{ payment.description || 'No description' }}</span>
                </div>
              </td>
              <td class="p-4">
                <span class="font-bold text-lg">Rp {{ formatCurrency(payment.amount) }}</span>
              </td>
              <td class="p-4">
                <Badge variant="default">
                  {{ payment.category?.name }}
                </Badge>
              </td>
              <td class="p-4 text-muted-foreground">
                {{ payment.payment_method?.name }}
              </td>
              <td class="p-4">
                <Badge :variant="getStatusVariant(payment.status)">
                  {{ payment.status }}
                </Badge>
              </td>
              <td class="p-4 text-sm text-muted-foreground">
                {{ formatDate(payment.transaction_date) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <template v-if="total > limit" #footer>
        <div class="flex items-center justify-between">
          <p class="text-sm text-muted-foreground">
            Showing {{ (page - 1) * limit + 1 }} to {{ Math.min(page * limit, total) }} of
            {{ total }} payments
          </p>
          <div class="flex gap-2">
            <Button variant="outline" size="sm" :disabled="page === 1" @click="page--">
              <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M15 19l-7-7 7-7"
                />
              </svg>
              Previous
            </Button>
            <Button variant="outline" size="sm" :disabled="page * limit >= total" @click="page++">
              Next
              <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 5l7 7-7 7"
                />
              </svg>
            </Button>
          </div>
        </div>
      </template>
    </Card>
  </div>
</template>

<script setup lang="ts">
import type { Payment } from '~/types/payment'
import dayjs from 'dayjs'

definePageMeta({
  middleware: 'auth',
})

const { getPayments } = usePayment()

const payments = ref<Payment[]>([])
const loading = ref(true)
const page = ref(1)
const limit = ref(10)
const total = ref(0)

const filters = reactive({
  search: '',
  status: '',
})

const hasActiveFilters = computed(() => {
  return filters.search !== '' || filters.status !== ''
})

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('id-ID').format(amount)
}

const formatDate = (date: string) => {
  return dayjs(date).format('DD MMM YYYY')
}

const getStatusVariant = (status: string): 'success' | 'warning' | 'error' | 'info' | 'default' => {
  const variants: Record<string, 'success' | 'warning' | 'error' | 'info'> = {
    completed: 'success',
    pending: 'warning',
    failed: 'error',
    refunded: 'info',
  }
  return variants[status] || 'default'
}

const loadPayments = async () => {
  loading.value = true
  const response = await getPayments(page.value, limit.value, filters.status, filters.search)

  if (response.success && response.data) {
    payments.value = response.data.payments
    total.value = response.data.total
  }

  loading.value = false
}

const clearFilters = () => {
  filters.search = ''
  filters.status = ''
  page.value = 1
  loadPayments()
}

const removeSearchFilter = () => {
  filters.search = ''
  loadPayments()
}

const removeStatusFilter = () => {
  filters.status = ''
  loadPayments()
}

// Debounced search
let searchTimeout: ReturnType<typeof setTimeout>
const debouncedSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    page.value = 1
    loadPayments()
  }, 500)
}

watch(page, loadPayments)

onMounted(loadPayments)
</script>

<style scoped>
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.fade-in {
  animation: fadeInUp 0.3s ease-out forwards;
  opacity: 0;
}
</style>
