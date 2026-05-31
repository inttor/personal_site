<template>
  <div class="min-h-screen bg-stone-50 dark:bg-stone-950 flex transition-colors duration-500">
    <!-- Sidebar -->
    <aside class="w-56 bg-white dark:bg-stone-900 border-r border-stone-200 dark:border-stone-800 shrink-0 flex flex-col transition-colors duration-500">
      <div class="px-6 py-6 border-b border-stone-100 dark:border-stone-800">
        <router-link to="/admin" class="group leading-none">
          <span class="font-display text-xl font-bold tracking-tight text-stone-900 dark:text-stone-100">
            Karrl
          </span>
          <span class="block text-[10px] font-mono tracking-wider uppercase text-stone-400 dark:text-stone-500 mt-1">
            Admin
          </span>
        </router-link>
      </div>

      <nav class="flex-1 px-3 py-4 space-y-1">
        <router-link
          to="/admin"
          :class="navClass('/admin', true)"
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
    <div class="flex-1 flex flex-col">
      <header class="bg-white dark:bg-stone-900 border-b border-stone-200 dark:border-stone-800 px-8 py-4 flex items-center justify-between transition-colors duration-500">
        <h2 class="text-sm text-stone-400 dark:text-stone-500 tracking-wide">
          {{ route.meta?.title || 'Dashboard' }}
        </h2>
        <div class="flex items-center gap-3 text-sm text-stone-400 dark:text-stone-500">
          <router-link to="/" class="hover:text-stone-700 dark:hover:text-stone-300 transition-colors duration-500">View Site</router-link>
        </div>
      </header>

      <main class="flex-1 p-8">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

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
