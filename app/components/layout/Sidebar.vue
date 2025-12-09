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
const uiStore = useUIStore();

const menuItems = [
  {
    label: "Dashboard",
    path: "/",
    icon: "IconDashboard",
  },
  {
    label: "Payments",
    path: "/payments",
    icon: "IconPayments",
  },
  {
    label: "Categories",
    path: "/categories",
    icon: "IconCategories",
  },
];

const sidebarClasses = computed(() => {
  return uiStore.sidebarOpen
    ? "w-64 translate-x-0"
    : "w-64 -translate-x-full lg:translate-x-0";
});
</script>

<script lang="ts">
// Icon components
const IconDashboard = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
    </svg>
  `,
};

const IconPayments = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
    </svg>
  `,
};

const IconCategories = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
    </svg>
  `,
};
</script>
