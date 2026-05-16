<template>
  <div class="animate-fade-in-up">
    <!-- Profile Hero Section: High-end minimalism -->
    <div class="mb-20 flex flex-col md:flex-row items-center md:items-start gap-10">
      <img src="\src\img\accountPicture.jpg" alt="Avatar" class="w-24 h-24 rounded-full grayscale hover:grayscale-0 transition duration-500 ease-in-out" />
      <div class="text-center md:text-left pt-2">
        <h1 class="text-3xl font-bold mb-2 text-zinc-900 tracking-tight">Karrl</h1>
        <p class="text-zinc-500 font-mono text-sm mb-4">
          > STDIN | Think >> /dev/Mind
        </p>
        <p class="text-zinc-600 leading-relaxed max-w-xl text-base">
          竹杖芒鞋轻胜马，谁怕？一蓑烟雨任平生。
        </p>
      </div>
    </div>

    <!-- Main Content with Sidebar -->
    <div class="flex flex-col md:flex-row gap-12">
      <!-- Sidebar Navigation (Tags & Dates) -->
      <aside class="md:w-48 shrink-0">
        <div class="sticky top-12 space-y-10">
          
          <!-- Filter by Year -->
          <div v-if="uniqueYears.length">
            <h4 class="text-[11px] font-mono tracking-wider uppercase text-zinc-400 mb-4 select-none">Archives</h4>
            <ul class="space-y-2 text-sm font-light">
              <li 
                class="cursor-pointer transition-colors duration-300 flex items-center gap-2"
                :class="selectedYear === '' ? 'text-zinc-900 font-medium' : 'text-zinc-500 hover:text-zinc-800'"
                @click="selectedYear = ''; selectedTag = ''"
              >
                All
              </li>
              <li 
                v-for="year in uniqueYears" :key="year"
                class="cursor-pointer transition-colors duration-300 flex items-center gap-2"
                :class="selectedYear === year ? 'text-zinc-900 font-medium' : 'text-zinc-500 hover:text-zinc-800'"
                @click="toggleYear(year)"
              >
                {{ year }}
              </li>
            </ul>
          </div>

          <!-- Filter by Tags -->
          <div v-if="uniqueTags.length">
            <h4 class="text-[11px] font-mono tracking-wider uppercase text-zinc-400 mb-4 select-none">Tags</h4>
            <div class="flex flex-wrap gap-2">
              <span 
                v-for="tag in uniqueTags" :key="tag"
                @click="toggleTag(tag)"
                class="text-xs cursor-pointer px-2.5 py-1 rounded-md transition-all duration-300"
                :class="selectedTag === tag ? 'bg-zinc-800 text-white' : 'bg-zinc-100 text-zinc-500 hover:bg-zinc-200 hover:text-zinc-800'"
              >
                #{{ tag }}
              </span>
            </div>
          </div>
          
        </div>
      </aside>

      <!-- Article List -->
      <div class="flex-1 min-w-0">
        <div v-if="loading" class="text-zinc-400 animate-pulse">Loading insights...</div>
        <div v-else-if="error" class="text-red-400">{{ error }}</div>
        
        <div v-else class="space-y-2">
          <div v-if="filteredArticles.length === 0" class="text-zinc-400 font-light py-8">
            No articles found for this filter.
          </div>
          
          <div v-for="article in filteredArticles" :key="article.id" 
               class="group cursor-pointer flex flex-col md:flex-row md:items-baseline justify-between py-5 border-b border-zinc-100 last:border-0 hover:bg-zinc-50/50 -mx-6 px-6 rounded-xl transition duration-300"
               @click="$router.push({ name: 'article', params: { id: article.id } })">
            
            <h3 class="text-lg text-zinc-800 group-hover:text-zinc-900 font-medium transition-colors duration-300 line-clamp-2 md:line-clamp-1">
              {{ article.title }}
            </h3>
            
            <div class="text-sm text-zinc-400 flex items-center gap-6 shrink-0 mt-1 md:mt-0 font-light tracking-wide">
              <span v-if="article.category?.name || article.category_name">
                {{ article.category?.name || article.category_name }}
              </span>
              <time>{{ new Date(article.created_at).toLocaleDateString('en-US', { year: 'numeric', month: 'short', day: 'numeric' }) }}</time>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '@/api'

const articles = ref([])
const loading = ref(true)
const error = ref(null)

const selectedTag = ref('')
const selectedYear = ref('')

// Compute unique years from articles
const uniqueYears = computed(() => {
  const years = new Set()
  articles.value.forEach(article => {
    if (article.created_at) {
      years.add(new Date(article.created_at).getFullYear())
    }
  })
  return Array.from(years).sort((a, b) => b - a)
})

// Compute unique tags from articles
const uniqueTags = computed(() => {
  const tags = new Set()
  articles.value.forEach(article => {
    if (article.tags && Array.isArray(article.tags)) {
      article.tags.forEach(tag => tags.add(tag))
    }
  })
  return Array.from(tags).sort()
})

// Filter actions
const toggleTag = (tag) => {
  selectedTag.value = selectedTag.value === tag ? '' : tag
  // Keep year filter active, or optionally reset it:
  // selectedYear.value = ''
}

const toggleYear = (year) => {
  selectedYear.value = selectedYear.value === year ? '' : year
  // Keep tag filter active, or optionally reset it:
  selectedTag.value = ''
}

// Filtered articles list
const filteredArticles = computed(() => {
  return articles.value.filter(article => {
    // Tag filter
    if (selectedTag.value && (!article.tags || !article.tags.includes(selectedTag.value))) {
      return false
    }
    // Year filter logic
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