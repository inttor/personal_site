<template>
  <div class="min-h-screen bg-zinc-50 flex">
    <!-- Sidebar -->
    <aside class="w-56 bg-white border-r border-zinc-100 shrink-0 flex flex-col">
      <div class="px-6 py-6 border-b border-zinc-100">
        <router-link to="/admin" class="text-lg font-bold tracking-tight text-zinc-900">
          Karrl. Admin
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

      <div class="px-3 py-4 border-t border-zinc-100">
        <button
          @click="handleLogout"
          class="flex items-center gap-2.5 w-full px-3 py-2 text-sm text-zinc-500 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/></svg>
          Sign out
        </button>
      </div>
    </aside>

    <!-- Main -->
    <div class="flex-1 flex flex-col">
      <header class="bg-white border-b border-zinc-100 px-8 py-4 flex items-center justify-between">
        <h2 class="text-sm font-medium text-zinc-500">
          {{ route.meta?.title || 'Dashboard' }}
        </h2>
        <div class="flex items-center gap-3 text-sm text-zinc-500">
          <router-link to="/" class="hover:text-zinc-900 transition-colors">View Site</router-link>
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
    'flex items-center gap-2.5 px-3 py-2 text-sm rounded-lg transition-colors',
    match ? 'bg-zinc-100 text-zinc-900 font-medium' : 'text-zinc-500 hover:text-zinc-900 hover:bg-zinc-50'
  ].join(' ')
}

const handleLogout = () => {
  localStorage.removeItem('admin_token')
  localStorage.removeItem('admin_user')
  router.push('/admin/login')
}
</script>
