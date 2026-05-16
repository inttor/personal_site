<template>
  <div>
    <!-- Header -->
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-2xl font-bold text-zinc-900">Articles</h1>
        <p class="text-sm text-zinc-500 mt-1">Manage all your articles</p>
      </div>
      <router-link
        to="/admin/articles/create"
        class="inline-flex items-center gap-2 px-4 py-2 bg-zinc-900 text-white rounded-lg hover:bg-zinc-800 transition-colors text-sm font-medium"
      >
        <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/></svg>
        New Article
      </router-link>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-center py-16 text-zinc-400">Loading articles...</div>

    <!-- Error -->
    <div v-else-if="errorMsg" class="text-center py-16 text-red-500">{{ errorMsg }}</div>

    <!-- Empty -->
    <div v-else-if="articles.length === 0" class="text-center py-16">
      <p class="text-zinc-400">No articles yet.</p>
      <router-link to="/admin/articles/create" class="text-sm text-zinc-900 underline hover:no-underline mt-2 inline-block">
        Write your first article
      </router-link>
    </div>

    <!-- Table -->
    <div v-else class="bg-white rounded-xl border border-zinc-100 shadow-sm overflow-hidden">
      <table class="w-full">
        <thead>
          <tr class="border-b border-zinc-100 bg-zinc-50/50">
            <th class="text-left px-5 py-3 text-xs font-medium text-zinc-500 uppercase tracking-wider">Title</th>
            <th class="text-left px-5 py-3 text-xs font-medium text-zinc-500 uppercase tracking-wider hidden md:table-cell">Category</th>
            <th class="text-left px-5 py-3 text-xs font-medium text-zinc-500 uppercase tracking-wider hidden sm:table-cell">Date</th>
            <th class="text-left px-5 py-3 text-xs font-medium text-zinc-500 uppercase tracking-wider hidden sm:table-cell">Views</th>
            <th class="text-right px-5 py-3 text-xs font-medium text-zinc-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-zinc-50">
          <tr v-for="article in articles" :key="article.id" class="hover:bg-zinc-50/50 transition-colors">
            <td class="px-5 py-4">
              <router-link :to="`/admin/articles/${article.id}/edit`" class="text-sm font-medium text-zinc-900 hover:text-zinc-600 transition-colors line-clamp-1">
                {{ article.title }}
              </router-link>
            </td>
            <td class="px-5 py-4 text-sm text-zinc-500 hidden md:table-cell">
              <span class="inline-block bg-zinc-100 px-2 py-0.5 rounded text-xs">{{ article.category_name || 'Uncategorized' }}</span>
            </td>
            <td class="px-5 py-4 text-sm text-zinc-500 hidden sm:table-cell">{{ new Date(article.created_at).toLocaleDateString('en-US', { year: 'numeric', month: 'short', day: 'numeric' }) }}</td>
            <td class="px-5 py-4 text-sm text-zinc-500 hidden sm:table-cell">{{ article.view_count ?? 0 }}</td>
            <td class="px-5 py-4 text-right">
              <div class="flex items-center justify-end gap-2">
                <router-link
                  :to="`/admin/articles/${article.id}/edit`"
                  class="px-3 py-1.5 text-xs font-medium text-zinc-600 hover:text-zinc-900 hover:bg-zinc-100 rounded-md transition-colors"
                >
                  Edit
                </router-link>
                <button
                  @click="confirmDelete(article)"
                  class="px-3 py-1.5 text-xs font-medium text-red-500 hover:text-red-700 hover:bg-red-50 rounded-md transition-colors"
                >
                  Delete
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="flex items-center justify-between px-5 py-4 border-t border-zinc-100 bg-zinc-50/50">
        <p class="text-xs text-zinc-500">
          Page {{ currentPage }} of {{ totalPages }}
        </p>
        <div class="flex gap-2">
          <button
            :disabled="currentPage <= 1"
            @click="changePage(currentPage - 1)"
            class="px-3 py-1.5 text-xs font-medium rounded-md transition-colors disabled:opacity-30 disabled:cursor-not-allowed"
            :class="currentPage > 1 ? 'text-zinc-700 hover:bg-zinc-200' : 'text-zinc-300'"
          >
            Previous
          </button>
          <button
            :disabled="currentPage >= totalPages"
            @click="changePage(currentPage + 1)"
            class="px-3 py-1.5 text-xs font-medium rounded-md transition-colors disabled:opacity-30 disabled:cursor-not-allowed"
            :class="currentPage < totalPages ? 'text-zinc-700 hover:bg-zinc-200' : 'text-zinc-300'"
          >
            Next
          </button>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="deleting" class="fixed inset-0 z-50 flex items-center justify-center bg-black/20 backdrop-blur-sm" @click.self="deleting = null">
      <div class="bg-white rounded-2xl shadow-xl border border-zinc-100 p-6 max-w-sm w-full mx-4">
        <h3 class="text-lg font-semibold text-zinc-900 mb-2">Delete Article</h3>
        <p class="text-sm text-zinc-500 mb-1">Are you sure you want to delete:</p>
        <p class="text-sm font-medium text-zinc-800 mb-6">"{{ deleting.title }}"</p>
        <p v-if="deleteError" class="text-red-500 text-sm mb-4">{{ deleteError }}</p>
        <div class="flex justify-end gap-3">
          <button
            @click="deleting = null"
            class="px-4 py-2 text-sm font-medium text-zinc-600 hover:text-zinc-900 hover:bg-zinc-100 rounded-lg transition-colors"
          >
            Cancel
          </button>
          <button
            @click="handleDelete"
            :disabled="deletingLoading"
            class="px-4 py-2 text-sm font-medium bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 transition-colors"
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
