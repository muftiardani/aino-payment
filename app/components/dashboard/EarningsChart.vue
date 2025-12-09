<template>
  <Card class="p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="font-semibold text-lg">Earnings Overview</h3>
        <p class="text-sm text-muted-foreground">Monthly revenue for the current year</p>
      </div>
    </div>

    <div class="h-[300px] w-full">
      <div v-if="loading" class="h-full flex items-center justify-center">
        <IconLoading class="w-8 h-8 text-primary animate-spin" />
      </div>
      <div v-else-if="error" class="h-full flex items-center justify-center text-red-500">
        {{ error }}
      </div>
      <Bar v-else :data="chartData" :options="chartOptions" />
    </div>
  </Card>
</template>

<script setup lang="ts">
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  type TooltipItem,
  type ChartOptions,
} from 'chart.js'
import { Bar } from 'vue-chartjs'
import IconLoading from '~/components/icons/IconLoading.vue'
import type { MonthlyStats } from '~/types/payment'

// Register ChartJS components
ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

const { getMonthlyStats } = usePayment()
const loading = ref(true)
const error = ref('')
const stats = ref<MonthlyStats[]>([])

// Chart Colors
const primaryColor = '#4f46e5' // Indigo 600
const gridColor = '#e5e7eb' // Gray 200

const chartData = computed(() => {
  // Sort stats by month index if needed, or rely on backend order.
  // Backend returns formatted month "Jan", "Feb" etc.
  const labels = stats.value.map(s => s.month)
  const data = stats.value.map(s => s.total_amount)

  return {
    labels,
    datasets: [
      {
        label: 'Earnings',
        backgroundColor: primaryColor,
        borderRadius: 4,
        data,
      },
    ],
  }
})

const chartOptions: ChartOptions<'bar'> = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false,
    },
    tooltip: {
      callbacks: {
        label: (context: TooltipItem<'bar'>) => {
          return new Intl.NumberFormat('id-ID', {
            style: 'currency',
            currency: 'IDR',
          }).format(context.raw as number)
        },
      },
    },
  },
  scales: {
    y: {
      beginAtZero: true,
      grid: {
        color: gridColor,
      },
      ticks: {
        callback: (value: number | string) => {
          if (typeof value === 'number') {
            if (value >= 1000000) return `Rp${value / 1000000}M`
            if (value >= 1000) return `Rp${value / 1000}k`
          }
          return value
        },
      },
    },
    x: {
      grid: {
        display: false,
      },
    },
  },
}

const loadData = async () => {
  loading.value = true
  try {
    const response = await getMonthlyStats()
    if (response.success && response.data) {
      stats.value = response.data
    } else {
      error.value = 'Failed to load chart data'
    }
  } catch (err) {
    console.error(err)
    error.value = 'An error occurred'
  } finally {
    loading.value = false
  }
}

onMounted(loadData)
</script>
