<template>
  <div class="animate-fade-in-up mt-10">
    <div v-if="loading" class="text-zinc-400 py-10 tracking-wide font-light">Loading article...</div>
    <div v-else-if="error" class="text-red-400 py-10">{{ error }}</div>
    
    <div v-else class="flex flex-col lg:flex-row gap-8 xl:gap-16 max-w-5xl mx-auto px-6">
      <!-- Main Content -->
      <article class="flex-1 min-w-0 max-w-3xl">
        <!-- Article Header -->
        <header class="relative mb-12">
          <!-- Ambient Background Texture (Full width bleed, fading out) -->
          <div class="absolute -top-24 -left-32 -right-32 h-72 z-[-1] pointer-events-none" style="mask-image: linear-gradient(to bottom, black 10%, transparent 100%); -webkit-mask-image: linear-gradient(to bottom, black 10%, transparent 100%);">
            <img src="https://images.unsplash.com/photo-1557682250-33bd709cbe85?q=80&w=2000&auto=format&fit=crop" class="w-full h-full object-cover opacity-[0.06] grayscale" alt="" />
          </div>
          
          <h1 class="text-4xl md:text-5xl font-extrabold tracking-tight text-zinc-900 mb-6 leading-tight">{{ article.title }}</h1>
          
          <div class="flex flex-wrap items-center gap-4 text-zinc-500 text-sm font-light">
            <time class="flex items-center gap-1.5">
              <svg class="w-4 h-4 text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>
              {{ new Date(article.created_at).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' }) }}
            </time>
            <span v-if="article.category_name" class="flex items-center gap-1.5 text-zinc-600">
              <span class="w-1 h-1 rounded-full bg-zinc-300"></span>
              {{ article.category_name }}
            </span>
            <div class="flex gap-2 ml-auto">
              <span v-for="tag in article.tags" :key="tag" class="bg-zinc-50 text-zinc-600 text-xs px-2.5 py-1 rounded-md hover:bg-zinc-100 hover:text-zinc-900 transition-colors cursor-pointer border border-zinc-100">
                #{{ tag }}
              </span>
            </div>
          </div>
        </header>

        <div class="w-full h-px bg-gradient-to-r from-zinc-200/60 to-transparent mb-12"></div>
        
        <!-- Rendered HTML -->
        <div class="prose prose-zinc prose-h1:text-2xl prose-h1:font-bold prose-h2:text-xl prose-h2:font-bold  prose-h3:text-lg prose-p:text-justify prose-p:text-zinc-600 prose-a:text-zinc-900 hover:prose-a:text-zinc-600 prose-headings:text-zinc-900 prose-blockquote:border-l-zinc-300 prose-blockquote:bg-zinc-50 prose-blockquote:py-0.5 prose-blockquote:px-5 prose-blockquote:rounded-r-lg prose-blockquote:font-normal prose-blockquote:not-italic prose-blockquote:text-zinc-600 prose-pre:bg-zinc-900/95 prose-pre:shadow-sm prose-img:rounded-xl prose-img:shadow-sm markdown-body selection:bg-zinc-200" v-html="compiledMarkdown" ref="contentRef"></div>
      </article>

      <!-- Sidebar TOC -->
      <aside v-if="toc.length > 0" class="hidden lg:block w-64 shrink-0">
        <div class="sticky top-20 bg-zinc-50/50 border border-zinc-100/60 rounded-2xl p-6 backdrop-blur-sm shadow-sm max-h-[calc(100vh-6rem)] overflow-y-auto">
          <div class="flex items-center gap-2 mb-6">
            <svg class="w-4 h-4 text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16M4 12h16M4 18h7"></path></svg>
            <h4 class="text-[11px] font-mono tracking-wider uppercase text-zinc-500 select-none m-0">Table of Contents</h4>
          </div>
          <ul class="space-y-3 relative text-sm font-light">
            <!-- Timeline vertical line -->
            <div class="absolute left-0 top-2 bottom-2 w-px bg-zinc-100"></div>
            
            <li v-for="item in toc" :key="item.id" 
                class="transition-colors cursor-pointer block pr-3 relative"
                :class="activeHeading === item.id ? 'text-zinc-900 font-medium' : 'text-zinc-500 hover:text-zinc-800'"
                @click="scrollTo(item.id)">
              <div class="flex items-center">
                <!-- Highlight Dot/Indicator -->
                <div 
                  class="absolute -left-[5px] top-1.5 w-[5px] h-[5px] rounded-full transition-all duration-300"
                  :class="activeHeading === item.id ? 'bg-zinc-800 ring-4 ring-zinc-100' : 'bg-transparent'"
                ></div>
                
                <span :style="{ marginLeft: `${(item.level > 1 ? (item.level - 1) * 0.75 : 0) + 0.75}rem` }" class="line-clamp-2 leading-relaxed">
                  {{ item.text }}
                </span>
              </div>
            </li>
          </ul>
        </div>
      </aside>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed, nextTick, watch } from 'vue'
import { useRoute } from 'vue-router'
import { marked } from 'marked'
import api from '@/api'

const route = useRoute()
const article = ref(null)
const loading = ref(true)
const error = ref(null)

const contentRef = ref(null)
const toc = ref([])
const activeHeading = ref('')

const compiledMarkdown = computed(() => {
  if (!article.value || !article.value.content) return ''
  return marked(article.value.content)
})

const extractTOC = () => {
  if (!contentRef.value) return
  // Extract h1, h2, h3 and h4 headers
  const headings = Array.from(contentRef.value.querySelectorAll('h1, h2, h3, h4'))
  toc.value = headings.map((el, i) => {
    // Generate unique ID if none exists (marked might generate them automatically)
    if (!el.id) el.id = `heading-${i}`
    return {
      id: el.id,
      text: el.innerText,
      level: parseInt(el.tagName[1]) // "H2" -> 2
    }
  })
}

const scrollTo = (id) => {
  const el = document.getElementById(id)
  if (el) {
    // Scroll smoothly and offset slightly for visual padding
    const y = el.getBoundingClientRect().top + window.scrollY - 30
    window.scrollTo({ top: y, behavior: 'smooth' })
  }
}

const handleScroll = () => {
  if (!contentRef.value || toc.value.length === 0) return
  
  const headings = Array.from(contentRef.value.querySelectorAll('h1, h2, h3, h4'))
  let currentId = ''
  
  for (const el of headings) {
    const rect = el.getBoundingClientRect()
    // Define the trigger line 120px from top
    if (rect.top <= 120) {
      currentId = el.id
    } else {
      break
    }
  }
  
  // Default to first heading if at the very top
  activeHeading.value = currentId || (headings[0] ? headings[0].id : '')
}

// Extract TOC whenever article renders/updates
watch(article, async () => {
  if (article.value) {
    await nextTick() // Wait for Vue to inject v-html into the DOM
    extractTOC()
    handleScroll() // Trigger scroll check immediately
  }
})

onMounted(async () => {
  window.addEventListener('scroll', handleScroll, { passive: true })
  
  try {
    const id = route.params.id
    const response = await api.getArticleDetails(id)
    if (response.data.code === 200) {
      article.value = response.data.data
    } else {
      error.value = response.data.message || 'Failed to fetch article'
    }
  } catch (err) {
    error.value = 'Failed to fetch article'
    console.error(err)
  } finally {
    loading.value = false
  }
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<style scoped>
/* Ensure images in markdown are responsive and look polished */
.markdown-body img {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 1.25rem auto;
  border-radius: 0.5rem;
  box-shadow: 0 8px 24px rgba(15, 23, 42, 0.06);
}

/* Allow authors to use inline Tailwind-like classes in HTML img tags (e.g., class="w-80") */
.markdown-body img[class] {
  max-width: none; /* respect explicit width classes or inline styles */
}
</style>