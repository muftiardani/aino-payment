<template>
  <div>
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold">Payments</h1>
        <p class="text-muted-foreground">
          Manage all your payment transactions
        </p>
      </div>
      <NuxtLink to="/payments/create">
        <Button variant="primary">
          <svg
            class="w-5 h-5 mr-2"
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

    <Card>
      <Loading v-if="loading" :loading="loading" />

      <div
        v-else-if="payments.length === 0"
        class="text-center py-12 text-muted-foreground"
      >
        <svg
          class="w-16 h-16 mx-auto mb-4 opacity-50"
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
        <p class="text-lg font-medium mb-2">No payments found</p>
        <p class="text-sm">Create your first payment to get started</p>
      </div>

      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead class="border-b">
            <tr class="text-left">
              <th class="p-4 font-medium">Description</th>
              <th class="p-4 font-medium">Amount</th>
              <th class="p-4 font-medium">Category</th>
              <th class="p-4 font-medium">Method</th>
              <th class="p-4 font-medium">Status</th>
              <th class="p-4 font-medium">Date</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="payment in payments"
              :key="payment.id"
              class="border-b hover:bg-accent transition-colors cursor-pointer"
              @click="navigateTo(`/payments/${payment.id}`)"
            >
              <td class="p-4">{{ payment.description || "No description" }}</td>
              <td class="p-4 font-medium">
                Rp {{ formatCurrency(payment.amount) }}
              </td>
              <td class="p-4">{{ payment.category?.name }}</td>
              <td class="p-4">{{ payment.payment_method?.name }}</td>
              <td class="p-4">
                <span :class="getStatusClass(payment.status)" class="badge">
                  {{ payment.status }}
                </span>
              </td>
              <td class="p-4 text-sm text-muted-foreground">
                {{ formatDate(payment.transaction_date) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <template #footer v-if="total > limit">
        <div class="flex items-center justify-between">
          <p class="text-sm text-muted-foreground">
            Showing {{ (page - 1) * limit + 1 }} to
            {{ Math.min(page * limit, total) }} of {{ total }} payments
          </p>
          <div class="flex gap-2">
            <Button
              variant="outline"
              size="sm"
              :disabled="page === 1"
              @click="page--"
            >
              Previous
            </Button>
            <Button
              variant="outline"
              size="sm"
              :disabled="page * limit >= total"
              @click="page++"
            >
              Next
            </Button>
          </div>
        </div>
      </template>
    </Card>
  </div>
</template>

<script setup lang="ts">
import type { Payment } from "~/types/payment";
import dayjs from "dayjs";

definePageMeta({
  middleware: "auth",
});

const { getPayments } = usePayment();

const payments = ref<Payment[]>([]);
const loading = ref(true);
const page = ref(1);
const limit = ref(10);
const total = ref(0);

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat("id-ID").format(amount);
};

const formatDate = (date: string) => {
  return dayjs(date).format("DD MMM YYYY");
};

const getStatusClass = (status: string) => {
  const classes = {
    pending:
      "bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200",
    completed:
      "bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200",
    failed: "bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200",
    refunded: "bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200",
  };
  return classes[status as keyof typeof classes] || "";
};

const loadPayments = async () => {
  loading.value = true;
  const response = await getPayments(page.value, limit.value);

  if (response.success && response.data) {
    payments.value = response.data.payments;
    total.value = response.data.total;
  }

  loading.value = false;
};

watch(page, loadPayments);

onMounted(loadPayments);
</script>
