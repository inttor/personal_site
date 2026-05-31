<template>
  <!-- Admin routes get sidebar layout -->
  <router-view v-if="isAdminRoute" />

  <!-- Public site layout -->
  <div v-else class="min-h-screen bg-stone-50 dark:bg-stone-950 text-stone-700 dark:text-stone-300 font-sans flex flex-col transition-colors duration-500 grain-overlay" :style="{ backgroundImage: `url(${TEXTURE_DOTS})` }">
    <header class="w-full max-w-3xl mx-auto px-6 pt-16 pb-10 flex items-end justify-between">
      <router-link to="/" class="group leading-none">
        <span class="block font-display text-3xl font-bold tracking-tight text-stone-900 dark:text-stone-100 transition-colors duration-500">
          Karrl
        </span>
        <span class="block text-[11px] font-mono tracking-[0.2em] uppercase text-stone-400 dark:text-stone-500 mt-1.5 transition-colors duration-500 group-hover:text-amber-600 dark:group-hover:text-amber-500">
          文码之间
        </span>
      </router-link>
      <div class="flex items-center gap-8">
        <nav class="space-x-8 text-sm tracking-wide text-stone-400 dark:text-stone-500">
          <router-link to="/" class="hover:text-stone-800 dark:hover:text-stone-200 transition-colors duration-500 pb-1 border-b border-transparent hover:border-amber-400/40">
            Writing
          </router-link>
          <router-link to="/projects" class="hover:text-stone-800 dark:hover:text-stone-200 transition-colors duration-500 pb-1 border-b border-transparent hover:border-amber-400/40">
            Projects
          </router-link>
        </nav>
        <button
          @click="toggleDark"
          class="w-8 h-8 flex items-center justify-center rounded-full text-stone-400 dark:text-stone-500 hover:text-amber-600 dark:hover:text-amber-500 hover:bg-stone-200/60 dark:hover:bg-stone-800 transition-all duration-500"
          :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
        >
          <svg v-if="isDark" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/></svg>
          <svg v-else class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"/></svg>
        </button>
      </div>
    </header>

    <main class="flex-grow w-full max-w-3xl mx-auto px-6 pb-32">
      <router-view></router-view>
    </main>

    <footer class="w-full max-w-3xl mx-auto px-6 pb-10 relative">
      <!-- Decorative line image -->
      <img :src="DECO_LINE" alt="" class="w-48 h-3 mx-auto mb-6 opacity-40 dark:opacity-30 pointer-events-none" aria-hidden="true" />
      <div class="border-t border-stone-200 dark:border-stone-800 pt-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-stone-400 dark:text-stone-600 tracking-wide">
        <span>&copy; {{ year }} Karrl</span>
        <div class="flex items-center gap-4">
          <a href="https://beian.miit.gov.cn" target="_blank" rel="noopener noreferrer" class="hover:text-stone-600 dark:hover:text-stone-400 transition-colors duration-500">
            粤ICP备2023073336号-1
          </a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { TEXTURE_DOTS, DECO_LINE } from '@/config/images.js'

const route = useRoute()
const isAdminRoute = computed(() => route.path.startsWith('/admin'))

const isDark = ref(false)
const year = new Date().getFullYear()

const toggleDark = () => {
  isDark.value = !isDark.value
}

const applyTheme = (dark) => {
  document.documentElement.classList.toggle('dark', dark)
  localStorage.setItem('theme', dark ? 'dark' : 'light')
}

onMounted(() => {
  const saved = localStorage.getItem('theme')
  if (saved === 'dark' || (!saved && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    isDark.value = true
  }
  applyTheme(isDark.value)
})

watch(isDark, applyTheme)
</script>
