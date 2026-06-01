<template>
  <div class="animate-fade-in-up">
    <!-- Hero / Profile -->
    <section class="relative mb-12">
      <!-- Ambient background orbs -->
      <div class="absolute -top-32 -right-24 w-96 h-96 rounded-full bg-amber-100/30 dark:bg-amber-900/10 blur-3xl -z-10 pointer-events-none"></div>
      <div class="absolute -top-16 -left-24 w-72 h-72 rounded-full bg-stone-200/40 dark:bg-stone-800/20 blur-3xl -z-10 pointer-events-none"></div>

      <!-- Hero background image (subtle overlay) -->
      <div class="absolute inset-0 -z-20 pointer-events-none opacity-[0.04] dark:opacity-[0.03] rounded-2xl overflow-hidden">
        <img :src="HERO_BG" alt="" class="w-full h-full object-cover grayscale" loading="lazy" />
      </div>

      <div class="flex flex-col md:flex-row items-center md:items-start gap-8 mb-6">
        <img
          src="/img/accountPicture.jpg"
          alt="Karrl"
          class="w-20 h-20 rounded-full grayscale hover:grayscale-0 transition-all duration-700 ease-out ring-1 ring-stone-200 dark:ring-stone-800 hover:ring-amber-300/60 dark:hover:ring-amber-600/40"
        />
        <div class="text-center md:text-left pt-1">
          <h1 class="font-display text-4xl md:text-5xl font-bold text-stone-900 dark:text-stone-100 tracking-tight mb-3">
            Karrl
          </h1>
          <p class="text-stone-500 dark:text-stone-400 font-mono text-xs tracking-widest uppercase">
            Words and code, thoughts and craft
          </p>
        </div>
      </div>
      <p class="text-stone-500 dark:text-stone-400 leading-relaxed max-w-xl text-base italic font-display">
        竹杖芒鞋轻胜马，谁怕？一蓑烟雨任平生。
      </p>
    </section>

    <div class="divider"><span>&sect;</span></div>

    <!-- Filters + Article List -->
    <section>
      <div class="flex flex-col md:flex-row gap-8 md:gap-16">
        <!-- Filters -->
        <aside class="md:w-44 shrink-0">
          <div class="md:sticky md:top-16">
            <!-- Year filter -->
            <div v-if="uniqueYears.length" class="mb-6 md:mb-10">
              <h4 class="text-[10px] font-mono tracking-[0.2em] uppercase text-stone-400 dark:text-stone-500 mb-3 md:mb-5 select-none">
                By Year
              </h4>
              <div class="flex flex-row flex-wrap gap-1.5 md:flex-col md:gap-0">
                <button
                  class="text-sm cursor-pointer px-3 py-1 md:py-0 md:px-0 rounded-full md:rounded-none transition-all duration-500 md:flex md:items-center md:gap-2 md:w-full md:py-2.5"
                  :class="selectedYear === '' ? 'bg-amber-100 dark:bg-amber-900/30 text-amber-800 dark:text-amber-400 md:bg-transparent md:dark:bg-transparent md:text-amber-700 md:dark:text-amber-500 font-medium' : 'bg-stone-100 dark:bg-stone-800 text-stone-500 dark:text-stone-400 md:bg-transparent md:dark:bg-transparent md:hover:text-stone-700 md:dark:hover:text-stone-300'"
                  @click="selectedYear = ''; selectedTag = ''"
                >
                  All
                </button>
                <button
                  v-for="year in uniqueYears" :key="year"
                  class="text-sm cursor-pointer px-3 py-1 md:py-0 md:px-0 rounded-full md:rounded-none transition-all duration-500 md:flex md:items-center md:gap-2 md:w-full md:py-2.5"
                  :class="selectedYear === year ? 'bg-amber-100 dark:bg-amber-900/30 text-amber-800 dark:text-amber-400 md:bg-transparent md:dark:bg-transparent md:text-amber-700 md:dark:text-amber-500 font-medium' : 'bg-stone-100 dark:bg-stone-800 text-stone-500 dark:text-stone-400 md:bg-transparent md:dark:bg-transparent md:hover:text-stone-700 md:dark:hover:text-stone-300'"
                  @click="toggleYear(year)"
                >
                  {{ year }}
                </button>
              </div>
            </div>

            <!-- Topic filter -->
            <div v-if="uniqueTags.length">
              <h4 class="text-[10px] font-mono tracking-[0.2em] uppercase text-stone-400 dark:text-stone-500 mb-3 md:mb-5 select-none">
                Topics
              </h4>
              <div class="flex flex-wrap gap-1.5">
                <button
                  v-for="tag in uniqueTags" :key="tag"
                  @click="toggleTag(tag)"
                  class="text-xs cursor-pointer px-3 py-1 rounded-full transition-all duration-500"
                  :class="selectedTag === tag ? 'bg-amber-100 dark:bg-amber-900/30 text-amber-800 dark:text-amber-400 ring-1 ring-amber-300/60 dark:ring-amber-600/30' : 'bg-stone-100 dark:bg-stone-800 text-stone-500 dark:text-stone-400 hover:bg-stone-200 dark:hover:bg-stone-700 hover:text-stone-700 dark:hover:text-stone-300'"
                >
                  {{ tag }}
                </button>
              </div>
            </div>
          </div>
        </aside>

        <!-- Articles -->
        <div class="flex-1 min-w-0">
          <!-- Loading -->
          <div v-if="loading" class="space-y-0">
            <div v-for="i in 5" :key="i" class="flex flex-col md:flex-row md:items-baseline justify-between py-6 border-b border-stone-100 dark:border-stone-800 last:border-0">
              <div class="skeleton h-6 w-72 rounded mb-2 md:mb-0"></div>
              <div class="flex items-center gap-6">
                <div class="skeleton h-4 w-14 rounded"></div>
                <div class="skeleton h-4 w-24 rounded"></div>
              </div>
            </div>
          </div>

          <div v-else-if="error" class="text-red-500 dark:text-red-400 py-12 text-sm">{{ error }}</div>

          <div v-else class="space-y-0">
            <p v-if="filteredArticles.length === 0" class="text-stone-400 dark:text-stone-500 py-12 text-sm">
              No articles found.
            </p>

            <article
              v-for="article in filteredArticles" :key="article.id"
              class="group cursor-pointer flex flex-col md:flex-row md:items-baseline justify-between py-6 border-b border-stone-100 dark:border-stone-800 last:border-0 hover:bg-stone-50/60 dark:hover:bg-stone-900/30 -mx-4 px-4 rounded-lg transition-all duration-500"
              @click="$router.push({ name: 'article', params: { id: article.id } })"
            >
              <h3 class="text-lg text-stone-800 dark:text-stone-200 group-hover:text-amber-700 dark:group-hover:text-amber-500 font-medium transition-colors duration-500 line-clamp-2 md:line-clamp-1 font-display">
                {{ article.title }}
              </h3>

              <div class="text-xs text-stone-400 dark:text-stone-500 flex items-center gap-5 shrink-0 mt-1.5 md:mt-0 tracking-wide">
                <span v-if="article.category?.name || article.category_name" class="text-stone-500 dark:text-stone-400">
                  {{ article.category?.name || article.category_name }}
                </span>
                <time>{{ new Date(article.created_at).toLocaleDateString('en-US', { year: 'numeric', month: 'short', day: 'numeric' }) }}</time>
              </div>
            </article>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '@/api'
import { HERO_BG } from '@/config/images.js'

const articles = ref([])
const loading = ref(true)
const error = ref(null)

const selectedTag = ref('')
const selectedYear = ref('')

const uniqueYears = computed(() => {
  const years = new Set()
  articles.value.forEach(article => {
    if (article.created_at) {
      years.add(new Date(article.created_at).getFullYear())
    }
  })
  return Array.from(years).sort((a, b) => b - a)
})

const uniqueTags = computed(() => {
  const tags = new Set()
  articles.value.forEach(article => {
    if (article.tags && Array.isArray(article.tags)) {
      article.tags.forEach(tag => tags.add(tag))
    }
  })
  return Array.from(tags).sort()
})

const toggleTag = (tag) => {
  selectedTag.value = selectedTag.value === tag ? '' : tag
}

const toggleYear = (year) => {
  selectedYear.value = selectedYear.value === year ? '' : year
  selectedTag.value = ''
}

const filteredArticles = computed(() => {
  return articles.value.filter(article => {
    if (selectedTag.value && (!article.tags || !article.tags.includes(selectedTag.value))) {
      return false
    }
    if (selectedYear.value && new Date(article.created_at).getFullYear() !== selectedYear.value) {
      return false
    }
    return true
  })
})

onMounted(async () => {
  try {
    const response = await api.getArticles()
    if (response.data.code === 200) {
      articles.value = response.data.data.list || []
    } else {
      error.value = response.data.message || 'Failed to fetch articles'
    }
  } catch (err) {
    error.value = 'Failed to fetch articles'
    console.error(err)
  } finally {
    loading.value = false
  }
})
</script>
