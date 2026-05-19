<template>
  <!-- Admin routes get sidebar layout -->
  <router-view v-if="isAdminRoute" />

  <!-- Public site layout -->
  <div v-else class="min-h-screen bg-white text-zinc-800 font-sans selection:bg-zinc-900 selection:text-white flex flex-col">
    <header class="w-full max-w-5xl mx-auto px-6 py-12 flex items-center justify-between">
      <router-link to="/" class="text-xl font-bold tracking-tight text-zinc-900 hover:opacity-70 transition-opacity">
        Karrl.
      </router-link>
      <nav class="space-x-8 text-sm font-medium text-zinc-500">
        <router-link to="/" class="hover:text-zinc-900 transition-colors duration-300">Blog</router-link>
        <router-link to="/projects" class="hover:text-zinc-900 transition-colors duration-300">Projects</router-link>
      </nav>
    </header>

    <main class="flex-grow w-full max-w-5xl mx-auto px-6 pb-24">
      <router-view></router-view>
    </main>

    <footer class="w-full text-center py-6 text-sm text-zinc-500">
      <a href="https://beian.miit.gov.cn" target="_blank" rel="noopener noreferrer" class="hover:text-zinc-900 transition-colors">
        粤ICP备2023073336号-1
      </a>
    </footer>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const isAdminRoute = computed(() => route.path.startsWith('/admin'))
</script>
