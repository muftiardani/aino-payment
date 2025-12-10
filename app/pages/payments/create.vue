<template>
  <div>
    <div class="mb-6">
      <h1
        class="text-3xl font-bold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent"
      >
        Create Payment
      </h1>
      <p class="text-muted-foreground mt-1">Add a new payment transaction</p>
    </div>

    <Card>
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
            <p v-if="errors.category_id" class="text-sm text-destructive mt-1">
              {{ errors.category_id }}
            </p>
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
            <p v-if="errors.payment_method_id" class="text-sm text-destructive mt-1">
              {{ errors.payment_method_id }}
            </p>
          </div>

          <!-- Transaction Date -->
          <div class="md:col-span-2">
            <Input
              id="transaction_date"
              v-model="form.transaction_date"
              type="datetime-local"
              label="Transaction Date"
              required
              :error="errors.transaction_date"
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
            <p v-if="errors.description" class="text-sm text-destructive mt-1">
              {{ errors.description }}
            </p>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center justify-end gap-4 pt-4 border-t">
          <Button type="button" variant="outline" @click="navigateTo('/payments')">Cancel</Button>
          <Button type="submit" variant="primary" :loading="uiStore.loading">Create Payment</Button>
        </div>
      </form>
    </Card>
  </div>
</template>

<script setup lang="ts">
import type { Category, PaymentMethod } from '~/types/payment'
import dayjs from 'dayjs'
import { createPaymentSchema } from '~/utils/validation-schemas'

definePageMeta({
  middleware: 'auth',
})

const { createPayment, getCategories, getPaymentMethods } = usePayment()
const uiStore = useUIStore()
const { errors, validate } = useFormValidation(createPaymentSchema)

const categories = ref<Category[]>([])
const paymentMethods = ref<PaymentMethod[]>([])

const form = reactive({
  amount: '',
  category_id: '',
  payment_method_id: '',
  transaction_date: dayjs().format('YYYY-MM-DDTHH:mm'),
  description: '',
})

const handleSubmit = async () => {
  // Validate form
  if (!validate(form)) {
    return
  }

  // Submit
  const response = await createPayment({
    amount: Number.parseFloat(form.amount.toString()),
    category_id: form.category_id,
    payment_method_id: form.payment_method_id,
    transaction_date: new Date(form.transaction_date).toISOString(),
    description: form.description,
  })

  if (response.success) {
    navigateTo('/payments')
  }
}

onMounted(async () => {
  const [categoriesRes, methodsRes] = await Promise.all([getCategories(), getPaymentMethods()])

  if (categoriesRes.success && categoriesRes.data) {
    categories.value = categoriesRes.data
  }

  if (methodsRes.success && methodsRes.data) {
    paymentMethods.value = methodsRes.data
  }
})
</script>
