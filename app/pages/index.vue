<template>
  <div>
    <div class="mb-6 animate-fade-in">
      <h1
        class="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-gray-900 to-gray-600 dark:from-white dark:to-gray-400"
      >
        Dashboard
      </h1>
      <p class="text-muted-foreground mt-1">Welcome back, {{ authStore.user?.full_name }}!</p>
    </div>

    <!-- Statistics Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8 animate-fade-up">
      <Card class="group hover:-translate-y-1 transition-transform duration-300">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-muted-foreground">Total Payments</p>
            <p
              class="text-3xl font-bold mt-1 bg-clip-text text-transparent bg-gradient-to-r from-primary to-violet-500"
            >
              {{ stats?.total_payments || 0 }}
            </p>
          </div>
          <div
            class="p-4 bg-primary/10 rounded-2xl group-hover:scale-110 transition-transform duration-300"
          >
            <svg class="w-6 h-6 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
              />
            </svg>
          </div>
        </div>
      </Card>

      <Card class="group hover:-translate-y-1 transition-transform duration-300">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-muted-foreground">Completed</p>
            <p
              class="text-3xl font-bold mt-1 bg-clip-text text-transparent bg-gradient-to-r from-emerald-500 to-teal-500"
            >
              {{ stats?.completed_count || 0 }}
            </p>
          </div>
          <div
            class="p-4 bg-emerald-100 dark:bg-emerald-900/30 rounded-2xl group-hover:scale-110 transition-transform duration-300"
          >
            <svg
              class="w-6 h-6 text-emerald-600 dark:text-emerald-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
        </div>
      </Card>

      <Card class="group hover:-translate-y-1 transition-transform duration-300">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-muted-foreground">Pending</p>
            <p
              class="text-3xl font-bold mt-1 bg-clip-text text-transparent bg-gradient-to-r from-amber-500 to-orange-500"
            >
              {{ stats?.pending_count || 0 }}
            </p>
          </div>
          <div
            class="p-4 bg-amber-100 dark:bg-amber-900/30 rounded-2xl group-hover:scale-110 transition-transform duration-300"
          >
            <svg
              class="w-6 h-6 text-amber-600 dark:text-amber-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
        </div>
      </Card>

      <Card class="group hover:-translate-y-1 transition-transform duration-300">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-muted-foreground">Total Amount</p>
            <p
              class="text-3xl font-bold mt-1 bg-clip-text text-transparent bg-gradient-to-r from-blue-600 to-cyan-500"
            >
              Rp {{ formatCurrency(stats?.total_amount || 0) }}
            </p>
          </div>
          <div
            class="p-4 bg-blue-100 dark:bg-blue-900/30 rounded-2xl group-hover:scale-110 transition-transform duration-300"
          >
            <svg
              class="w-6 h-6 text-blue-600 dark:text-blue-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
        </div>
      </Card>
    </div>

    <!-- Earnings Chart -->
    <div class="mb-8 animate-fade-up" style="animation-delay: 0.1s">
      <EarningsChart />
    </div>

    <!-- Recent Payments -->
    <div class="animate-fade-up" style="animation-delay: 0.2s">
      <Card>
        <template #header>
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-bold">Recent Payments</h2>
            <NuxtLink to="/payments" class="text-sm text-primary hover:underline">
              View all
            </NuxtLink>
          </div>
        </template>

        <Loading v-if="loading" :loading="loading" />

        <div v-else-if="recentPayments.length === 0" class="text-center py-8 text-muted-foreground">
          No payments yet. Create your first payment!
        </div>

        <div v-else class="space-y-4">
          <div
            v-for="payment in recentPayments"
            :key="payment.id"
            class="flex items-center justify-between p-4 rounded-lg border hover:bg-accent transition-colors"
          >
            <div class="flex-1">
              <p class="font-medium">
                {{ payment.description || 'No description' }}
              </p>
              <p class="text-sm text-muted-foreground">
                {{ payment.category?.name }} • {{ payment.payment_method?.name }}
              </p>
              <p class="text-xs text-muted-foreground">
                {{ formatDate(payment.transaction_date) }}
              </p>
            </div>
            <div class="text-right">
              <p class="font-bold">Rp {{ formatCurrency(payment.amount) }}</p>
              <span :class="getStatusClass(payment.status)" class="badge">
                {{ payment.status }}
              </span>
            </div>
          </div>
        </div>

        <template #footer>
          <NuxtLink to="/payments/create">
            <Button variant="primary" class="w-full">Create New Payment</Button>
          </NuxtLink>
        </template>
      </Card>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { DashboardStats, Payment } from '~/types/payment'
import dayjs from 'dayjs'
import EarningsChart from '~/components/dashboard/EarningsChart.vue'

definePageMeta({
  middleware: 'auth',
})

const authStore = useAuthStore()
const { getDashboardStats, getRecentPayments } = usePayment()

const stats = ref<DashboardStats | null>(null)
const recentPayments = ref<Payment[]>([])
const loading = ref(true)

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('id-ID').format(amount)
}

const formatDate = (date: string) => {
  return dayjs(date).format('DD MMM YYYY, HH:mm')
}

const getStatusClass = (status: string) => {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200',
    completed: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200',
    failed: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200',
    refunded: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
  }
  return classes[status as keyof typeof classes] || ''
}

onMounted(async () => {
  loading.value = true

  const [statsRes, paymentsRes] = await Promise.all([getDashboardStats(), getRecentPayments()])

  if (statsRes.success && statsRes.data) {
    stats.value = statsRes.data
  }

  if (paymentsRes.success && paymentsRes.data) {
    recentPayments.value = paymentsRes.data
  }

  loading.value = false
})
</script>
