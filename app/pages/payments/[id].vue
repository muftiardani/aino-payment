<template>
  <div>
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1
          class="text-3xl font-bold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent"
        >
          Payment Details
        </h1>
        <p class="text-muted-foreground mt-1">View and manage payment information</p>
      </div>
      <div class="flex gap-2">
        <NuxtLink :to="`/payments/${id}/edit`">
          <Button variant="outline">
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
              />
            </svg>
            Edit
          </Button>
        </NuxtLink>
        <Button variant="destructive" @click="showDeleteModal = true">
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
            />
          </svg>
          Delete
        </Button>
      </div>
    </div>

    <Loading v-if="loading" :loading="loading" />

    <div v-else-if="payment" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Main Info -->
      <Card class="lg:col-span-2">
        <template #header>
          <h2 class="text-xl font-bold">Payment Information</h2>
        </template>

        <div class="space-y-6">
          <!-- Amount -->
          <div>
            <label class="text-sm text-muted-foreground">Amount</label>
            <p class="text-3xl font-bold text-primary">Rp {{ formatCurrency(payment.amount) }}</p>
          </div>

          <!-- Description -->
          <div>
            <label class="text-sm text-muted-foreground">Description</label>
            <p class="text-lg">
              {{ payment.description || 'No description provided' }}
            </p>
          </div>

          <!-- Details Grid -->
          <div class="grid grid-cols-2 gap-4 pt-4 border-t">
            <div>
              <label class="text-sm text-muted-foreground">Category</label>
              <p class="font-medium">{{ payment.category?.name }}</p>
            </div>
            <div>
              <label class="text-sm text-muted-foreground">Payment Method</label>
              <p class="font-medium">{{ payment.payment_method?.name }}</p>
            </div>
            <div>
              <label class="text-sm text-muted-foreground">Status</label>
              <div class="mt-1">
                <Badge :variant="getStatusVariant(payment.status)">
                  {{ payment.status }}
                </Badge>
              </div>
            </div>
            <div>
              <label class="text-sm text-muted-foreground">Transaction Date</label>
              <p class="font-medium">
                {{ formatDate(payment.transaction_date) }}
              </p>
            </div>
          </div>
        </div>
      </Card>

      <!-- Metadata -->
      <Card>
        <template #header>
          <h2 class="text-xl font-bold">Metadata</h2>
        </template>

        <div class="space-y-4">
          <div>
            <label class="text-sm text-muted-foreground">Payment ID</label>
            <p class="text-xs font-mono bg-muted p-2 rounded mt-1 break-all">
              {{ payment.id }}
            </p>
          </div>
          <div>
            <label class="text-sm text-muted-foreground">Created At</label>
            <p class="font-medium">{{ formatDate(payment.created_at) }}</p>
          </div>
          <div>
            <label class="text-sm text-muted-foreground">Last Updated</label>
            <p class="font-medium">{{ formatDate(payment.updated_at) }}</p>
          </div>
        </div>
      </Card>
    </div>

    <!-- Delete Confirmation Modal -->
    <Modal :show="showDeleteModal" title="Delete Payment" @close="showDeleteModal = false">
      <div class="space-y-4">
        <p class="text-muted-foreground">
          Are you sure you want to delete this payment? This action cannot be undone.
        </p>
        <div class="p-4 bg-destructive/10 rounded-lg border border-destructive/20">
          <p class="font-medium">{{ payment?.description || 'Payment' }}</p>
          <p class="text-sm text-muted-foreground">
            Amount: Rp {{ formatCurrency(payment?.amount || 0) }}
          </p>
        </div>
      </div>

      <template #footer>
        <div class="flex justify-end gap-3">
          <Button variant="outline" @click="showDeleteModal = false">Cancel</Button>
          <Button variant="destructive" @click="handleDelete" :loading="uiStore.loading">
            Delete Payment
          </Button>
        </div>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import type { Payment } from '~/types/payment'
import dayjs from 'dayjs'

definePageMeta({
  middleware: 'auth',
})

const route = useRoute()
const { getPaymentById, deletePayment } = usePayment()
const uiStore = useUIStore()

const id = route.params.id as string
const payment = ref<Payment | null>(null)
const loading = ref(true)
const showDeleteModal = ref(false)

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('id-ID').format(amount)
}

const formatDate = (date: string) => {
  return dayjs(date).format('DD MMMM YYYY, HH:mm')
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

const handleDelete = async () => {
  const response = await deletePayment(id)
  if (response.success) {
    showDeleteModal.value = false
    navigateTo('/payments')
  }
}

onMounted(async () => {
  const response = await getPaymentById(id)
  if (response.success && response.data) {
    payment.value = response.data
  }
  loading.value = false
})
</script>
