<template>
  <div class="animate-fade-in-up mt-8">
    <!-- Loading skeleton -->
    <div v-if="loading" class="space-y-10 py-8">
      <div class="skeleton h-12 w-3/4 rounded-lg"></div>
      <div class="flex gap-4">
        <div class="skeleton h-4 w-20 rounded"></div>
        <div class="skeleton h-4 w-16 rounded"></div>
        <div class="skeleton h-4 w-24 rounded"></div>
      </div>
      <div class="space-y-4 mt-12">
        <div class="skeleton h-4 w-full rounded"></div>
        <div class="skeleton h-4 w-full rounded"></div>
        <div class="skeleton h-4 w-5/6 rounded"></div>
        <div class="skeleton h-4 w-full rounded"></div>
        <div class="skeleton h-4 w-4/6 rounded"></div>
        <div class="skeleton h-4 w-full rounded"></div>
        <div class="skeleton h-4 w-3/6 rounded"></div>
      </div>
    </div>

    <div v-else-if="error" class="text-red-500 dark:text-red-400 py-16 text-center">{{ error }}</div>

    <article v-else class="flex flex-col lg:flex-row gap-10 xl:gap-16">
      <!-- Main Content -->
      <div class="flex-1 min-w-0">
        <!-- Header -->
        <header class="relative mb-14">
          <div class="absolute -top-16 -left-20 -right-20 h-64 -z-10 pointer-events-none bg-gradient-to-b from-stone-200/50 dark:from-stone-800/30 to-transparent rounded-full blur-3xl"></div>
          <div class="absolute top-8 right-0 w-48 h-48 rounded-full bg-amber-100/20 dark:bg-amber-900/8 blur-3xl -z-10 pointer-events-none"></div>

          <h1 class="font-display text-4xl md:text-5xl font-bold tracking-tight text-stone-900 dark:text-stone-100 mb-6 leading-tight text-center">
            {{ article.title }}
          </h1>

          <!-- Article cover image -->
          <div v-if="article.cover_url" class="relative w-full h-40 md:h-56 rounded-lg overflow-hidden mb-6">
            <img
              :src="article.cover_url"
              :alt="article.title"
              class="w-full h-full object-cover"
            />
            <div class="absolute inset-0 bg-gradient-to-b from-transparent via-transparent to-stone-50/70 dark:to-stone-950/70 pointer-events-none"></div>
            <div class="absolute inset-0 ring-1 ring-inset ring-stone-900/5 dark:ring-stone-100/5 rounded-lg pointer-events-none"></div>
          </div>
          <div v-else class="w-full h-32 md:h-40 rounded-lg mb-6 overflow-hidden relative">
            <img :src="ARTICLE_COVER_FALLBACK" alt="" class="w-full h-full object-cover opacity-[0.06] dark:opacity-[0.04] grayscale" loading="lazy" />
            <div class="absolute inset-0 bg-gradient-to-b from-transparent to-stone-50/80 dark:to-stone-950/80"></div>
          </div>

          <div class="flex flex-wrap items-center gap-5 text-sm text-stone-400 dark:text-stone-500 tracking-wide">
            <time class="flex items-center gap-2">
              <svg class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
              {{ new Date(article.created_at).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' }) }}
            </time>
            <span v-if="article.category_name" class="flex items-center gap-2 text-stone-500 dark:text-stone-400">
              <span class="w-1 h-1 rounded-full bg-amber-400/60"></span>
              {{ article.category_name }}
            </span>
            <div class="flex gap-2 ml-auto">
              <span v-for="tag in article.tags" :key="tag" class="text-xs px-3 py-1 rounded-full bg-stone-100 dark:bg-stone-800 text-stone-500 dark:text-stone-400 hover:bg-amber-50 dark:hover:bg-amber-900/20 hover:text-amber-700 dark:hover:text-amber-500 transition-all duration-500 cursor-pointer border border-transparent hover:border-amber-200 dark:hover:border-amber-800/40">
                {{ tag }}
              </span>
            </div>
          </div>
        </header>

        <div class="w-full h-px bg-gradient-to-r from-stone-200 dark:from-stone-800 to-transparent mb-14"></div>

        <!-- Rendered HTML -->
        <div
          class="prose prose-stone dark:prose-invert prose-lg prose-p:text-justify max-w-none markdown-body"
          v-html="compiledMarkdown"
          ref="contentRef"
        ></div>
      </div>

      <!-- TOC Sidebar -->
      <aside v-if="toc.length > 0" class="hidden lg:block w-56 shrink-0">
        <div class="sticky top-20 bg-stone-50/80 dark:bg-stone-900/60 border border-stone-200/60 dark:border-stone-800/60 rounded-xl p-5 backdrop-blur-sm max-h-[calc(100vh-6rem)] overflow-y-auto">
          <div class="flex items-center gap-2 mb-5">
            <svg class="w-3.5 h-3.5 text-stone-400 dark:text-stone-500" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16M4 12h16M4 18h7"/></svg>
            <h4 class="text-[10px] font-mono tracking-[0.2em] uppercase text-stone-400 dark:text-stone-500 select-none m-0">
              Contents
            </h4>
          </div>
          <ul class="space-y-2.5 relative text-sm font-light">
            <div class="absolute left-0 top-1.5 bottom-1.5 w-px bg-stone-200 dark:bg-stone-800"></div>

            <li v-for="item in toc" :key="item.id"
                class="transition-all duration-500 cursor-pointer block pr-2 relative"
                :class="activeHeading === item.id ? 'text-amber-700 dark:text-amber-500 font-medium' : 'text-stone-400 dark:text-stone-500 hover:text-stone-700 dark:hover:text-stone-300'"
                @click="scrollTo(item.id)">
              <div class="flex items-center">
                <div
                  class="absolute -left-[4px] top-1.5 w-[4px] h-[4px] rounded-full transition-all duration-500"
                  :class="activeHeading === item.id ? 'bg-amber-500 ring-4 ring-amber-100 dark:ring-amber-900/30 scale-150' : 'bg-transparent'"
                ></div>

                <span :style="{ marginLeft: `${(item.level > 1 ? (item.level - 1) * 0.65 : 0) + 0.65}rem` }" class="line-clamp-2 leading-relaxed text-xs">
                  {{ item.text }}
                </span>
              </div>
            </li>
          </ul>
        </div>
      </aside>
    </article>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed, nextTick, watch } from 'vue'
import { useRoute } from 'vue-router'
import { marked } from 'marked'
import api from '@/api'
import { ARTICLE_COVER_FALLBACK } from '@/config/images.js'

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
  const headings = Array.from(contentRef.value.querySelectorAll('h1, h2, h3, h4'))
  toc.value = headings.map((el, i) => {
    if (!el.id) el.id = `heading-${i}`
    return {
      id: el.id,
      text: el.innerText,
      level: parseInt(el.tagName[1])
    }
  })
}

const scrollTo = (id) => {
  const el = document.getElementById(id)
  if (el) {
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
    if (rect.top <= 120) {
      currentId = el.id
    } else {
      break
    }
  }

  activeHeading.value = currentId || (headings[0] ? headings[0].id : '')
}

watch(article, async () => {
  if (article.value) {
    await nextTick()
    extractTOC()
    handleScroll()
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
.markdown-body {
  @apply font-sans;
}

.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4 {
  @apply font-display tracking-tight;
}

.markdown-body h1 {
  @apply text-3xl font-bold mt-16 mb-6;
}

.markdown-body h2 {
  @apply text-2xl font-semibold mt-14 mb-5;
}

.markdown-body h3 {
  @apply text-xl font-medium mt-10 mb-4;
}

.markdown-body p {
  @apply leading-relaxed text-stone-600 dark:text-stone-400;
}

.markdown-body a {
  @apply text-amber-700 dark:text-amber-500 no-underline hover:underline decoration-amber-300 dark:decoration-amber-700 underline-offset-2;
}

.markdown-body blockquote {
  @apply border-l-2 border-amber-400/60 bg-stone-50 dark:bg-stone-900/50 py-4 px-6 rounded-r-lg font-normal not-italic text-stone-600 dark:text-stone-400 my-8;
}

.markdown-body blockquote p {
  @apply text-stone-600 dark:text-stone-400;
}

.markdown-body pre {
  @apply bg-stone-900 dark:bg-stone-800/90 shadow-sm rounded-xl my-8;
}

.markdown-body code {
  @apply text-stone-800 dark:text-stone-200 bg-stone-100 dark:bg-stone-800 rounded px-1.5 py-0.5 text-sm font-mono;
}

.markdown-body pre code {
  @apply bg-transparent p-0 text-stone-200;
}

.markdown-body img {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 2rem auto;
  border-radius: 0.75rem;
  box-shadow: 0 4px 32px rgba(28, 25, 23, 0.08);
}

.markdown-body img[class] {
  max-width: none;
}

.markdown-body hr {
  @apply border-stone-200 dark:border-stone-800 my-12;
}
</style>
