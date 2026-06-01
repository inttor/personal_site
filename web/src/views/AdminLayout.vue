<template>
  <div class="min-h-screen bg-stone-50 dark:bg-stone-950 flex transition-colors duration-500">
    <!-- Mobile sidebar backdrop -->
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="sidebarOpen"
        class="fixed inset-0 z-40 bg-stone-900/40 dark:bg-black/60 md:hidden"
        @click="sidebarOpen = false"
      />
    </Transition>

    <!-- Sidebar -->
    <aside
      class="fixed md:static inset-y-0 left-0 z-50 w-56 bg-white dark:bg-stone-900 border-r border-stone-200 dark:border-stone-800 shrink-0 flex flex-col transition-transform duration-300 md:translate-x-0"
      :class="sidebarOpen ? 'translate-x-0' : '-translate-x-full'"
    >
      <div class="px-6 py-6 border-b border-stone-100 dark:border-stone-800 flex items-center justify-between">
        <router-link to="/admin" class="group leading-none" @click="sidebarOpen = false">
          <span class="font-display text-xl font-bold tracking-tight text-stone-900 dark:text-stone-100">
            Karrl
          </span>
          <span class="block text-[10px] font-mono tracking-wider uppercase text-stone-400 dark:text-stone-500 mt-1">
            Admin
          </span>
        </router-link>
        <button
          class="md:hidden w-8 h-8 flex items-center justify-center rounded-lg text-stone-400 hover:text-stone-600 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors"
          @click="sidebarOpen = false"
          aria-label="Close sidebar"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
      </div>

      <nav class="flex-1 px-3 py-4 space-y-1">
        <router-link
          to="/admin"
          :class="navClass('/admin', true)"
          @click="sidebarOpen = false"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"/></svg>
          Articles
        </router-link>
      </nav>

      <div class="px-3 py-4 border-t border-stone-100 dark:border-stone-800">
        <button
          @click="handleLogout"
          class="flex items-center gap-2.5 w-full px-3 py-2 text-sm text-stone-400 dark:text-stone-500 hover:text-red-600 hover:bg-red-50 dark:hover:bg-red-950/30 rounded-lg transition-colors duration-500"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/></svg>
          Sign out
        </button>
      </div>
    </aside>

    <!-- Main -->
    <div class="flex-1 flex flex-col min-w-0">
      <header class="bg-white dark:bg-stone-900 border-b border-stone-200 dark:border-stone-800 px-4 md:px-8 py-4 flex items-center justify-between transition-colors duration-500">
        <div class="flex items-center gap-3">
          <button
            class="md:hidden w-8 h-8 flex items-center justify-center rounded-lg text-stone-400 hover:text-stone-600 hover:bg-stone-100 dark:hover:bg-stone-800 transition-colors"
            @click="sidebarOpen = true"
            aria-label="Open sidebar"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16M4 12h16M4 18h16"/></svg>
          </button>
          <h2 class="text-sm text-stone-400 dark:text-stone-500 tracking-wide">
            {{ route.meta?.title || 'Dashboard' }}
          </h2>
        </div>
        <div class="flex items-center gap-3 text-sm text-stone-400 dark:text-stone-500">
          <router-link to="/" class="hover:text-stone-700 dark:hover:text-stone-300 transition-colors duration-500">View Site</router-link>
        </div>
      </header>

      <main class="flex-1 p-4 md:p-8 overflow-auto">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const sidebarOpen = ref(false)

const navClass = (path, exact = false) => {
  const match = exact ? route.path === path : route.path.startsWith(path)
  return [
    'flex items-center gap-2.5 px-3 py-2 text-sm rounded-lg transition-all duration-500',
    match ? 'bg-amber-50 dark:bg-amber-900/20 text-amber-700 dark:text-amber-500 font-medium' : 'text-stone-400 dark:text-stone-500 hover:text-stone-700 dark:hover:text-stone-300 hover:bg-stone-50 dark:hover:bg-stone-800/50'
  ].join(' ')
}

const handleLogout = () => {
  localStorage.removeItem('admin_token')
  localStorage.removeItem('admin_user')
  router.push('/admin/login')
}
</script>
