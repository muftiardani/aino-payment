import type {
  Payment,
  PaymentMethod,
  Category,
  CreatePaymentRequest,
  UpdatePaymentRequest,
  PaymentListResponse,
  DashboardStats,
} from '~/types/payment'

export const usePayment = () => {
  const { $api } = useNuxtApp()
  const uiStore = useUIStore()

  const getPayments = async (page = 1, limit = 10, status = '', search = '') => {
    const params: Record<string, any> = { page, limit }
    if (status) params.status = status
    if (search) params.search = search

    return await $api<PaymentListResponse>('/payments', { params })
  }

  const getPaymentById = async (id: string) => {
    return await $api<Payment>(`/payments/${id}`)
  }

  const createPayment = async (data: CreatePaymentRequest) => {
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
    return await $api<Category[]>('/categories')
  }

  const getPaymentMethods = async () => {
    return await $api<PaymentMethod[]>('/payment-methods')
  }

  const getDashboardStats = async () => {
    return await $api<DashboardStats>('/dashboard/stats')
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
    getRecentPayments,
  }
}
