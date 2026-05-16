<template>
  <div>
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-zinc-900">{{ isEditing ? 'Edit Article' : 'New Article' }}</h1>
      <p class="text-sm text-zinc-500 mt-1">{{ isEditing ? 'Update your article content' : 'Write a new article' }}</p>
    </div>

    <div v-if="loadingArticle" class="text-center py-16 text-zinc-400">Loading article...</div>

    <form v-else @submit.prevent="handleSubmit" class="space-y-6 max-w-4xl">
      <!-- Title -->
      <div>
        <label class="block text-sm font-medium text-zinc-700 mb-1.5">Title</label>
        <input
          v-model="form.title"
          type="text"
          required
          class="w-full px-4 py-2.5 border border-zinc-200 rounded-lg focus:ring-2 focus:ring-zinc-900/10 focus:border-zinc-900 outline-none transition-all text-sm"
          placeholder="Enter article title..."
        />
      </div>

      <!-- Summary -->
      <div>
        <label class="block text-sm font-medium text-zinc-700 mb-1.5">Summary</label>
        <textarea
          v-model="form.summary"
          rows="2"
          class="w-full px-4 py-2.5 border border-zinc-200 rounded-lg focus:ring-2 focus:ring-zinc-900/10 focus:border-zinc-900 outline-none transition-all text-sm resize-none"
          placeholder="Brief summary of the article..."
        ></textarea>
      </div>

      <!-- Category -->
      <div>
        <label class="block text-sm font-medium text-zinc-700 mb-1.5">Category</label>
        <select
          v-model="form.category_id"
          class="w-full px-4 py-2.5 border border-zinc-200 rounded-lg focus:ring-2 focus:ring-zinc-900/10 focus:border-zinc-900 outline-none transition-all text-sm bg-white"
        >
          <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
        </select>
      </div>

      <!-- Content -->
      <div>
        <div class="flex items-center justify-between mb-1.5">
          <label class="block text-sm font-medium text-zinc-700">Content (Markdown)</label>
          <label class="cursor-pointer text-xs bg-zinc-100 hover:bg-zinc-200 text-zinc-600 py-1.5 px-3 rounded-md inline-flex items-center transition-colors">
            <svg class="w-3.5 h-3.5 mr-1.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>
            Insert Image
            <input type="file" accept="image/*" class="hidden" @change="handleImageUpload">
          </label>
        </div>
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <div>
            <textarea
              v-model="form.content"
              ref="editorRef"
              required
              rows="22"
              class="w-full px-4 py-3 border border-zinc-200 rounded-lg focus:ring-2 focus:ring-zinc-900/10 focus:border-zinc-900 outline-none transition-all font-mono text-sm resize-y"
              placeholder="Write your article content here in Markdown format..."
            ></textarea>
          </div>
          <div class="border border-zinc-200 rounded-lg p-4 bg-zinc-50/50 overflow-y-auto prose prose-zinc prose-sm max-w-none markdown-body" style="height: 480px;" v-html="compiledContent"></div>
        </div>
      </div>

      <!-- Actions -->
      <div class="flex items-center gap-3 pt-4 border-t border-zinc-100">
        <button
          type="submit"
          :disabled="submitting"
          class="px-5 py-2.5 bg-zinc-900 text-white rounded-lg hover:bg-zinc-800 disabled:opacity-50 transition-colors text-sm font-medium"
        >
          {{ submitting ? (isEditing ? 'Updating...' : 'Publishing...') : (isEditing ? 'Update Article' : 'Publish Article') }}
        </button>
        <router-link
          to="/admin"
          class="px-5 py-2.5 text-sm font-medium text-zinc-600 hover:text-zinc-900 hover:bg-zinc-100 rounded-lg transition-colors"
        >
          Cancel
        </router-link>
        <span v-if="successMsg" class="text-sm text-green-600 ml-2">{{ successMsg }}</span>
        <span v-if="errorMsg" class="text-sm text-red-500 ml-2">{{ errorMsg }}</span>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { marked } from 'marked'
import api from '@/api'

const router = useRouter()
const route = useRoute()

const isEditing = computed(() => !!route.params.id)
const editorRef = ref(null)

const form = ref({
  title: '',
  summary: '',
  content: '',
  category_id: 1
})

const categories = ref([])
const loadingArticle = ref(false)
const submitting = ref(false)
const errorMsg = ref('')
const successMsg = ref('')

const compiledContent = computed(() => {
  if (!form.value.content) return '<p class="text-zinc-400 italic">Preview will appear here...</p>'
  return marked(form.value.content)
})

const handleImageUpload = async (e) => {
  const file = e.target.files[0]
  if (!file) return
  e.target.value = ''

  try {
    const res = await api.uploadImage(file)
    if (res.data.code === 200) {
      const imageUrl = res.data.data.url
      const html = `<img src="${imageUrl}" alt="" class="w-full md:w-2/3 mx-auto rounded-lg" />`
      insertIntoEditor(html, { placeCursorInAlt: true })
    } else {
      alert(res.data.message || 'Image upload failed')
    }
  } catch (err) {
    console.error(err)
    alert('An error occurred during image upload')
  }
}

const insertIntoEditor = (text, opts = {}) => {
  const textarea = editorRef.value
  if (!textarea) {
    form.value.content += `\n${text}\n`
    return
  }
  const start = textarea.selectionStart
  const end = textarea.selectionEnd
  form.value.content = form.value.content.substring(0, start) + text + form.value.content.substring(end)
  setTimeout(() => {
    textarea.focus()
    let cursorPos = start + text.length
    if (opts.placeCursorInAlt) {
      const altIndex = text.indexOf('alt="')
      if (altIndex !== -1) cursorPos = start + altIndex + 5
    }
    textarea.selectionStart = cursorPos
    textarea.selectionEnd = cursorPos
  }, 0)
}

const handleSubmit = async () => {
  submitting.value = true
  errorMsg.value = ''
  successMsg.value = ''

  try {
    if (isEditing.value) {
      const res = await api.updateArticle(route.params.id, form.value)
      if (res.data.code === 200) {
        successMsg.value = 'Article updated successfully!'
        setTimeout(() => router.push('/admin'), 1000)
        errorMsg.value = res.data.message || 'Update failed'
      }
    } else {
      const res = await api.createArticle(form.value)
      if (res.data.code === 200) {
        successMsg.value = 'Article published successfully!'
        setTimeout(() => router.push('/admin'), 1000)
      } else {
        errorMsg.value = res.data.message || 'Publish failed'
      }
    }
  } catch (err) {
    errorMsg.value = err.response?.data?.message || err.message || 'An error occurred'
  } finally {
    submitting.value = false
  }
}

onMounted(async () => {
  // Load categories
  try {
    const res = await api.getCategories()
    if (res.data.code === 200) {
      categories.value = res.data.data || []
    }
  } catch (e) {
    console.error('Failed to load categories', e)
  }

  // Load article for editing
  if (!isEditing.value) return
  loadingArticle.value = true
  try {
    const res = await api.getArticleDetails(route.params.id)
    if (res.data.code === 200) {
      const article = res.data.data
      form.value.title = article.title || ''
      form.value.summary = article.summary || ''
      form.value.content = article.content || ''
      form.value.category_id = article.category_id || 1
    }
  } catch (err) {
    errorMsg.value = err.response?.data?.message || err.message || 'Failed to load article'
  } finally {
    loadingArticle.value = false
  }
})
</script>
