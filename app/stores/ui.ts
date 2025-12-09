import { defineStore } from 'pinia'

export const useUIStore = defineStore('ui', {
  state: () => ({
    sidebarOpen: true,
    darkMode: false,
    loading: false,
    toast: {
      show: false,
      message: '',
      type: 'info' as 'success' | 'error' | 'warning' | 'info',
    },
  }),

  actions: {
    toggleSidebar() {
      this.sidebarOpen = !this.sidebarOpen
    },

    toggleDarkMode() {
      this.darkMode = !this.darkMode
      if (import.meta.client) {
        localStorage.setItem('darkMode', String(this.darkMode))
        document.documentElement.classList.toggle('dark', this.darkMode)
      }
    },

    loadDarkMode() {
      if (import.meta.client) {
        const saved = localStorage.getItem('darkMode')
        this.darkMode = saved === 'true'
        document.documentElement.classList.toggle('dark', this.darkMode)
      }
    },

    setLoading(value: boolean) {
      this.loading = value
    },

    showToast(message: string, type: 'success' | 'error' | 'warning' | 'info' = 'info') {
      this.toast = { show: true, message, type }
      setTimeout(() => {
        this.toast.show = false
      }, 3000)
    },
  },
})
