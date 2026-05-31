<template>
  <div>
    <!-- Header -->
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-2xl font-bold text-stone-900 dark:text-stone-100 font-display">Articles</h1>
        <p class="text-sm text-stone-400 dark:text-stone-500 mt-1">Manage all your articles</p>
      </div>
      <router-link
        to="/admin/articles/create"
        class="inline-flex items-center gap-2 px-4 py-2 bg-amber-700 hover:bg-amber-800 text-white rounded-lg transition-colors duration-500 text-sm font-medium"
      >
        <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
        New Article
      </router-link>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-16 text-stone-400">Loading articles...</div>

    <!-- Error -->
    <div v-else-if="errorMsg" class="text-center py-16 text-red-500">{{ errorMsg }}</div>

    <!-- Empty -->
    <div v-else-if="articles.length === 0" class="text-center py-16">
      <p class="text-stone-400">No articles yet.</p>
      <router-link to="/admin/articles/create" class="text-sm text-amber-700 dark:text-amber-500 hover:underline mt-2 inline-block">
        Write your first article
      </router-link>
    </div>

    <!-- Table -->
    <div v-else class="bg-white dark:bg-stone-900 rounded-xl border border-stone-200 dark:border-stone-800 shadow-sm overflow-hidden transition-colors duration-500">
      <table class="w-full">
        <thead>
          <tr class="border-b border-stone-100 dark:border-stone-800 bg-stone-50/50 dark:bg-stone-900/50">
            <th class="text-left px-5 py-3 text-[10px] font-mono tracking-wider text-stone-400 dark:text-stone-500 uppercase">Title</th>
            <th class="text-left px-5 py-3 text-[10px] font-mono tracking-wider text-stone-400 dark:text-stone-500 uppercase hidden md:table-cell">Category</th>
            <th class="text-left px-5 py-3 text-[10px] font-mono tracking-wider text-stone-400 dark:text-stone-500 uppercase hidden sm:table-cell">Date</th>
            <th class="text-left px-5 py-3 text-[10px] font-mono tracking-wider text-stone-400 dark:text-stone-500 uppercase hidden sm:table-cell">Views</th>
            <th class="text-right px-5 py-3 text-[10px] font-mono tracking-wider text-stone-400 dark:text-stone-500 uppercase">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-stone-50 dark:divide-stone-800">
          <tr v-for="article in articles" :key="article.id" class="hover:bg-stone-50/50 dark:hover:bg-stone-800/30 transition-colors duration-300">
            <td class="px-5 py-4">
              <router-link :to="`/admin/articles/${article.id}/edit`" class="text-sm font-medium text-stone-900 dark:text-stone-100 hover:text-amber-700 dark:hover:text-amber-500 transition-colors duration-300 line-clamp-1">
                {{ article.title }}
              </router-link>
            </td>
            <td class="px-5 py-4 text-sm text-stone-400 hidden md:table-cell">
              <span class="inline-block bg-stone-100 dark:bg-stone-800 px-2 py-0.5 rounded text-xs">{{ article.category_name || 'Uncategorized' }}</span>
            </td>
            <td class="px-5 py-4 text-sm text-stone-400 hidden sm:table-cell">{{ new Date(article.created_at).toLocaleDateString('en-US', { year: 'numeric', month: 'short', day: 'numeric' }) }}</td>
            <td class="px-5 py-4 text-sm text-stone-400 hidden sm:table-cell">{{ article.view_count ?? 0 }}</td>
            <td class="px-5 py-4 text-right">
              <div class="flex items-center justify-end gap-2">
                <router-link
                  :to="`/admin/articles/${article.id}/edit`"
                  class="px-3 py-1.5 text-xs font-medium text-stone-500 hover:text-stone-800 dark:hover:text-stone-200 hover:bg-stone-100 dark:hover:bg-stone-800 rounded-md transition-colors duration-300"
                >
                  Edit
                </router-link>
                <button
                  @click="confirmDelete(article)"
                  class="px-3 py-1.5 text-xs font-medium text-red-500 hover:text-red-700 hover:bg-red-50 dark:hover:bg-red-950/30 rounded-md transition-colors duration-300"
                >
                  Delete
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="flex items-center justify-between px-5 py-4 border-t border-stone-100 dark:border-stone-800 bg-stone-50/50 dark:bg-stone-900/50">
        <p class="text-xs text-stone-400">
          Page {{ currentPage }} of {{ totalPages }}
        </p>
        <div class="flex gap-2">
          <button
            :disabled="currentPage <= 1"
            @click="changePage(currentPage - 1)"
            class="px-3 py-1.5 text-xs font-medium rounded-md transition-colors duration-300 disabled:opacity-30 disabled:cursor-not-allowed"
            :class="currentPage > 1 ? 'text-stone-600 dark:text-stone-300 hover:bg-stone-200 dark:hover:bg-stone-800' : 'text-stone-300 dark:text-stone-700'"
          >
            Previous
          </button>
          <button
            :disabled="currentPage >= totalPages"
            @click="changePage(currentPage + 1)"
            class="px-3 py-1.5 text-xs font-medium rounded-md transition-colors duration-300 disabled:opacity-30 disabled:cursor-not-allowed"
            :class="currentPage < totalPages ? 'text-stone-600 dark:text-stone-300 hover:bg-stone-200 dark:hover:bg-stone-800' : 'text-stone-300 dark:text-stone-700'"
          >
            Next
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="deleting" class="fixed inset-0 z-50 flex items-center justify-center bg-stone-950/40 backdrop-blur-sm" @click.self="deleting = null">
      <div class="bg-white dark:bg-stone-900 rounded-2xl shadow-xl border border-stone-200 dark:border-stone-800 p-6 max-w-sm w-full mx-4">
        <h3 class="text-lg font-semibold text-stone-900 dark:text-stone-100 mb-2 font-display">Delete Article</h3>
        <p class="text-sm text-stone-400 mb-1">Are you sure you want to delete:</p>
        <p class="text-sm font-medium text-stone-700 dark:text-stone-300 mb-6">"{{ deleting.title }}"</p>
        <p v-if="deleteError" class="text-red-500 text-sm mb-4">{{ deleteError }}</p>
        <div class="flex justify-end gap-3">
          <button
            @click="deleting = null"
            class="px-4 py-2 text-sm font-medium text-stone-500 hover:text-stone-800 dark:hover:text-stone-200 hover:bg-stone-100 dark:hover:bg-stone-800 rounded-lg transition-colors duration-300"
          >
            Cancel
          </button>
          <button
            @click="handleDelete"
            :disabled="deletingLoading"
            class="px-4 py-2 text-sm font-medium bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 transition-colors duration-300"
          >
            {{ deletingLoading ? 'Deleting...' : 'Delete' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '@/api'

const articles = ref([])
const loading = ref(true)
const errorMsg = ref('')
const currentPage = ref(1)
const totalPages = ref(1)
const pageSize = 10

const deleting = ref(null)
const deletingLoading = ref(false)
const deleteError = ref('')

const fetchArticles = async () => {
  loading.value = true
  errorMsg.value = ''
  try {
    const res = await api.getArticles(currentPage.value, pageSize)
    if (res.data.code === 200) {
      const d = res.data.data
      articles.value = d.list || []
      totalPages.value = Math.max(1, Math.ceil((d.total || 0) / pageSize))
    }
  } catch (err) {
    errorMsg.value = err.response?.data?.message || err.message || 'Failed to load articles'
  } finally {
    loading.value = false
  }
}

const changePage = (page) => {
  currentPage.value = page
  fetchArticles()
}

const confirmDelete = (article) => {
  deleting.value = article
  deleteError.value = ''
}

const handleDelete = async () => {
  if (!deleting.value) return
  deletingLoading.value = true
  try {
    const res = await api.deleteArticle(deleting.value.id)
    if (res.data.code === 200) {
      articles.value = articles.value.filter(a => a.id !== deleting.value.id)
      deleting.value = null
    } else {
      deleteError.value = res.data.message || 'Delete failed'
    }
  } catch (err) {
    deleteError.value = err.response?.data?.message || err.message || 'Delete failed'
  } finally {
    deletingLoading.value = false
  }
}

onMounted(fetchArticles)
</script>
