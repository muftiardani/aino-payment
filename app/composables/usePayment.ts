import type {
  Payment,
  PaymentMethod,
  Category,
  CreatePaymentRequest,
  UpdatePaymentRequest,
  PaymentListResponse,
  DashboardStats,
  MonthlyStats,
} from '~/types/payment'

import type { ApiResponse } from '~/types/api'
import { useCache } from './useCache'

export const usePayment = () => {
  const { $api } = useNuxtApp()
  const uiStore = useUIStore()

  const getPayments = async (
    page = 1,
    limit = 10,
    status = '',
    search = '',
    minAmount?: number,
    maxAmount?: number,
    startDate?: string,
    endDate?: string
  ) => {
    // Build query params
    const params: Record<string, string | number> = { page, limit }
    if (status) params.status = status
    if (search) params.search = search
    if (minAmount !== undefined) params.min_amount = minAmount
    if (maxAmount !== undefined) params.max_amount = maxAmount
    if (startDate) params.start_date = startDate
    if (endDate) params.end_date = endDate

    return await $api<PaymentListResponse>('/payments', { params })
  }

  const getPaymentById = async (id: string) => {
    return await $api<Payment>(`/payments/${id}`)
  }

  const createPayment = async (data: CreatePaymentRequest): Promise<ApiResponse<Payment>> => {
    uiStore.setLoading(true)
    try {
      const response = await $api<Payment>('/payments', {
        method: 'POST',
        body: JSON.stringify(data),
      })

      if (response.success) {
        uiStore.showToast('Payment created successfully!', 'success')
      } else {
        uiStore.showToast(response.error || 'Failed to create payment', 'error')
      }

      return response
    } finally {
      uiStore.setLoading(false)
    }
  }

  const updatePayment = async (id: string, data: UpdatePaymentRequest) => {
    uiStore.setLoading(true)
    try {
      const response = await $api<Payment>(`/payments/${id}`, {
        method: 'PUT',
        body: JSON.stringify(data),
      })

      if (response.success) {
        uiStore.showToast('Payment updated successfully!', 'success')
      } else {
        uiStore.showToast(response.error || 'Failed to update payment', 'error')
      }

      return response
    } finally {
      uiStore.setLoading(false)
    }
  }

  const deletePayment = async (id: string) => {
    uiStore.setLoading(true)
    try {
      const response = await $api(`/payments/${id}`, {
        method: 'DELETE',
      })

      if (response.success) {
        uiStore.showToast('Payment deleted successfully!', 'success')
      } else {
        uiStore.showToast(response.error || 'Failed to delete payment', 'error')
      }

      return response
    } finally {
      uiStore.setLoading(false)
    }
  }

  const getCategories = async () => {
    const cache = useCache<Category[]>('categories', 600000) // 10 minutes

    const cached = cache.get()
    if (cached) {
      return { success: true, data: cached }
    }

    const response = await $api<Category[]>('/categories')

    if (response.success && response.data) {
      cache.set(response.data)
    }

    return response
  }

  const getPaymentMethods = async () => {
    const cache = useCache<PaymentMethod[]>('payment_methods', 600000) // 10 minutes

    const cached = cache.get()
    if (cached) {
      return { success: true, data: cached }
    }

    const response = await $api<PaymentMethod[]>('/payment-methods')

    if (response.success && response.data) {
      cache.set(response.data)
    }

    return response
  }

  const getDashboardStats = async () => {
    return await $api<DashboardStats>('/dashboard/stats')
  }

  const getMonthlyStats = async (year?: number) => {
    const params: Record<string, string | number> = {}
    if (year) params.year = year
    return await $api<MonthlyStats[]>('/dashboard/chart', { params })
  }

  const exportPayments = async (
    status = '',
    search = '',
    minAmount?: number,
    maxAmount?: number,
    startDate?: string,
    endDate?: string
  ) => {
    const params: Record<string, string | number | boolean> = {}
    if (status) params.status = status
    if (search) params.search = search
    if (minAmount !== undefined) params.min_amount = minAmount
    if (maxAmount !== undefined) params.max_amount = maxAmount
    if (startDate) params.start_date = startDate
    if (endDate) params.end_date = endDate

    const response = await $api<Blob>('/payments/export', {
      params,
      responseType: 'blob',
    })

    // Create download link
    const blob = response as unknown as Blob
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `payments-${new Date().toISOString().split('T')[0]}.csv`)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }

  const getRecentPayments = async () => {
    return await $api<Payment[]>('/dashboard/recent')
  }

  return {
    getPayments,
    getPaymentById,
    createPayment,
    updatePayment,
    deletePayment,
    getCategories,
    getPaymentMethods,
    getDashboardStats,
    getMonthlyStats,
    exportPayments,
    getRecentPayments,
  }
}
