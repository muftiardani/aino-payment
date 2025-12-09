<template>
  <aside
    :class="sidebarClasses"
    class="fixed left-0 top-16 h-[calc(100vh-4rem)] bg-card border-r border-border transition-transform duration-300 z-30"
  >
    <nav class="p-4 space-y-2">
      <NuxtLink
        v-for="item in menuItems"
        :key="item.path"
        :to="item.path"
        class="flex items-center gap-3 px-4 py-3 rounded-lg hover:bg-accent transition-colors"
        active-class="bg-primary text-primary-foreground hover:bg-primary/90"
      >
        <component :is="item.icon" class="w-5 h-5" />
        <span class="font-medium">{{ item.label }}</span>
      </NuxtLink>
    </nav>
  </aside>

  <!-- Overlay for mobile -->
  <div
    v-if="uiStore.sidebarOpen"
    class="fixed inset-0 bg-black/50 z-20 lg:hidden"
    @click="uiStore.toggleSidebar"
  />
</template>

<script setup lang="ts">
import IconDashboard from '~/components/icons/IconDashboard.vue'
import IconPayments from '~/components/icons/IconPayments.vue'
import IconCategories from '~/components/icons/IconCategories.vue'

const uiStore = useUIStore()

const menuItems = [
  {
    label: 'Dashboard',
    path: '/',
    icon: IconDashboard,
  },
  {
    label: 'Payments',
    path: '/payments',
    icon: IconPayments,
  },
  {
    label: 'Categories',
    path: '/categories',
    icon: IconCategories,
  },
]

const sidebarClasses = computed(() => {
  return uiStore.sidebarOpen ? 'w-64 translate-x-0' : 'w-64 -translate-x-full lg:translate-x-0'
})
</script>
