<template>
  <!-- Admin routes get sidebar layout -->
  <router-view v-if="isAdminRoute" />

  <!-- Public site layout -->
  <div v-else class="min-h-screen bg-stone-50 dark:bg-stone-950 text-stone-700 dark:text-stone-300 font-sans flex flex-col transition-colors duration-500 grain-overlay" :style="{ backgroundImage: `url(${TEXTURE_DOTS})` }">
    <header class="w-full max-w-5xl mx-auto px-5 sm:px-6 pt-8 sm:pt-12 md:pt-16 pb-6 md:pb-10">
      <div class="flex items-end justify-between">
        <router-link to="/" class="group leading-none shrink-0">
          <span class="block font-display text-2xl sm:text-3xl font-bold tracking-tight text-stone-900 dark:text-stone-100 transition-colors duration-500">
            Karrl
          </span>
          <span class="block text-[10px] sm:text-[11px] font-mono tracking-[0.2em] uppercase text-stone-400 dark:text-stone-500 mt-1 sm:mt-1.5 transition-colors duration-500 group-hover:text-amber-600 dark:group-hover:text-amber-500">
            文码之间
          </span>
        </router-link>

        <!-- Desktop nav -->
        <div class="hidden md:flex items-center gap-8">
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

        <!-- Mobile toggles -->
        <div class="flex md:hidden items-center gap-2">
          <button
            @click="toggleDark"
            class="w-9 h-9 flex items-center justify-center rounded-full text-stone-400 dark:text-stone-500 hover:text-amber-600 dark:hover:text-amber-500 hover:bg-stone-200/60 dark:hover:bg-stone-800 transition-all duration-500"
            :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
          >
            <svg v-if="isDark" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/></svg>
            <svg v-else class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"/></svg>
          </button>
          <button
            @click="mobileMenuOpen = !mobileMenuOpen"
            class="w-9 h-9 flex items-center justify-center rounded-full text-stone-400 dark:text-stone-500 hover:text-amber-600 dark:hover:text-amber-500 hover:bg-stone-200/60 dark:hover:bg-stone-800 transition-all duration-500"
            :aria-label="mobileMenuOpen ? 'Close menu' : 'Open menu'"
            :aria-expanded="mobileMenuOpen"
          >
            <svg v-if="mobileMenuOpen" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 18L18 6M6 6l12 12"/></svg>
            <svg v-else class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16M4 12h16M4 18h16"/></svg>
          </button>
        </div>
      </div>

      <!-- Mobile menu -->
      <Transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="opacity-0 -translate-y-2"
        enter-to-class="opacity-100 translate-y-0"
        leave-active-class="transition duration-150 ease-in"
        leave-from-class="opacity-100 translate-y-0"
        leave-to-class="opacity-0 -translate-y-2"
      >
        <nav v-if="mobileMenuOpen" class="md:hidden mt-5 pt-4 border-t border-stone-200 dark:border-stone-800 flex flex-col gap-1">
          <router-link
            to="/"
            @click="mobileMenuOpen = false"
            class="py-2.5 px-2 -mx-2 rounded-lg text-sm tracking-wide text-stone-500 dark:text-stone-400 hover:text-stone-900 dark:hover:text-stone-100 hover:bg-stone-200/60 dark:hover:bg-stone-800 transition-colors duration-200"
            active-class="!text-amber-600 dark:!text-amber-500 !bg-amber-50 dark:!bg-amber-950/40"
          >
            Writing
          </router-link>
          <router-link
            to="/projects"
            @click="mobileMenuOpen = false"
            class="py-2.5 px-2 -mx-2 rounded-lg text-sm tracking-wide text-stone-500 dark:text-stone-400 hover:text-stone-900 dark:hover:text-stone-100 hover:bg-stone-200/60 dark:hover:bg-stone-800 transition-colors duration-200"
            active-class="!text-amber-600 dark:!text-amber-500 !bg-amber-50 dark:!bg-amber-950/40"
          >
            Projects
          </router-link>
        </nav>
      </Transition>
    </header>

    <main class="flex-grow w-full max-w-5xl mx-auto px-6 pb-32">
      <router-view></router-view>
    </main>

    <footer class="w-full max-w-5xl mx-auto px-6 pb-10 relative">
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

const mobileMenuOpen = ref(false)
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
watch(() => route.path, () => { mobileMenuOpen.value = false })
</script>
