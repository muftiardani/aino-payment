<template>
  <div>
    <div class="mb-6">
      <h1
        class="text-3xl font-bold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent"
      >
        Edit Payment
      </h1>
      <p class="text-muted-foreground mt-1">Update payment information</p>
    </div>

    <Loading v-if="loading" :loading="loading" />

    <Card v-else>
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Amount -->
          <div class="md:col-span-2">
            <Input
              id="amount"
              v-model="form.amount"
              type="number"
              label="Amount"
              placeholder="100000"
              required
              :error="errors.amount"
              hint="Enter amount in IDR"
            />
          </div>

          <!-- Status -->
          <div>
            <label class="block text-sm font-medium text-foreground mb-2">
              Status
              <span class="text-destructive">*</span>
            </label>
            <select v-model="form.status" class="input w-full" required>
              <option value="pending">Pending</option>
              <option value="completed">Completed</option>
              <option value="failed">Failed</option>
              <option value="refunded">Refunded</option>
            </select>
          </div>

          <!-- Category -->
          <div>
            <label class="block text-sm font-medium text-foreground mb-2">
              Category
              <span class="text-destructive">*</span>
            </label>
            <select v-model="form.category_id" class="input w-full" required>
              <option value="">Select category</option>
              <option v-for="category in categories" :key="category.id" :value="category.id">
                {{ category.name }}
              </option>
            </select>
          </div>

          <!-- Payment Method -->
          <div>
            <label class="block text-sm font-medium text-foreground mb-2">
              Payment Method
              <span class="text-destructive">*</span>
            </label>
            <select v-model="form.payment_method_id" class="input w-full" required>
              <option value="">Select payment method</option>
              <option v-for="method in paymentMethods" :key="method.id" :value="method.id">
                {{ method.name }}
              </option>
            </select>
          </div>

          <!-- Transaction Date -->
          <div>
            <Input
              id="transaction_date"
              v-model="form.transaction_date"
              type="datetime-local"
              label="Transaction Date"
              required
            />
          </div>

          <!-- Description -->
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-foreground mb-2">Description</label>
            <textarea
              v-model="form.description"
              class="input w-full min-h-[100px]"
              placeholder="Enter payment description..."
            ></textarea>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center justify-end gap-4 pt-4 border-t">
          <Button type="button" variant="outline" @click="navigateTo(`/payments/${id}`)">
            Cancel
          </Button>
          <Button type="submit" variant="primary" :loading="uiStore.loading">Update Payment</Button>
        </div>
      </form>
    </Card>
  </div>
</template>

<script setup lang="ts">
import type { Category, PaymentMethod } from '~/types/payment'
import dayjs from 'dayjs'

definePageMeta({
  middleware: 'auth',
})

const route = useRoute()
const { getPaymentById, updatePayment, getCategories, getPaymentMethods } = usePayment()
const uiStore = useUIStore()

const id = route.params.id as string
const loading = ref(true)
const categories = ref<Category[]>([])
const paymentMethods = ref<PaymentMethod[]>([])

const form = reactive({
  amount: '',
  status: 'pending',
  category_id: '',
  payment_method_id: '',
  transaction_date: '',
  description: '',
})

const errors = reactive({
  amount: '',
})

const handleSubmit = async () => {
  if (!form.amount || parseFloat(form.amount) <= 0) {
    errors.amount = 'Amount must be greater than 0'
    return
  }

  const response = await updatePayment(id, {
    amount: parseFloat(form.amount),
    status: form.status as 'pending' | 'completed' | 'failed' | 'refunded',
    category_id: form.category_id,
    payment_method_id: form.payment_method_id,
    transaction_date: new Date(form.transaction_date).toISOString(),
    description: form.description,
  })

  if (response.success) {
    navigateTo(`/payments/${id}`)
  }
}

onMounted(async () => {
  const [paymentRes, categoriesRes, methodsRes] = await Promise.all([
    getPaymentById(id),
    getCategories(),
    getPaymentMethods(),
  ])

  if (paymentRes.success && paymentRes.data) {
    const payment = paymentRes.data
    form.amount = payment.amount.toString()
    form.status = payment.status
    form.category_id = payment.category_id
    form.payment_method_id = payment.payment_method_id
    form.transaction_date = dayjs(payment.transaction_date).format('YYYY-MM-DDTHH:mm')
    form.description = payment.description
  }

  if (categoriesRes.success && categoriesRes.data) {
    categories.value = categoriesRes.data
  }

  if (methodsRes.success && methodsRes.data) {
    paymentMethods.value = methodsRes.data
  }

  loading.value = false
})
</script>
