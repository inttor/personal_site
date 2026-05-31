<template>
  <div>
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-stone-900 dark:text-stone-100 font-display">{{ isEditing ? 'Edit Article' : 'New Article' }}</h1>
      <p class="text-sm text-stone-400 dark:text-stone-500 mt-1">{{ isEditing ? 'Update your article content' : 'Write a new article' }}</p>
    </div>

    <div v-if="loadingArticle" class="text-center py-16 text-stone-400">Loading article...</div>

    <form v-else @submit.prevent="handleSubmit" class="space-y-6 max-w-4xl">
      <!-- Title -->
      <div>
        <label class="block text-sm font-medium text-stone-600 dark:text-stone-300 mb-1.5">Title</label>
        <input
          v-model="form.title"
          type="text"
          required
          class="w-full px-4 py-2.5 border border-stone-200 dark:border-stone-700 bg-white dark:bg-stone-900 rounded-lg focus:ring-2 focus:ring-amber-500/20 focus:border-amber-600 dark:focus:border-amber-500 outline-none transition-all duration-300 text-sm text-stone-900 dark:text-stone-100 placeholder:text-stone-400"
          placeholder="Enter article title..."
        />
      </div>

      <!-- Summary -->
      <div>
        <label class="block text-sm font-medium text-stone-600 dark:text-stone-300 mb-1.5">Summary</label>
        <textarea
          v-model="form.summary"
          rows="2"
          class="w-full px-4 py-2.5 border border-stone-200 dark:border-stone-700 bg-white dark:bg-stone-900 rounded-lg focus:ring-2 focus:ring-amber-500/20 focus:border-amber-600 dark:focus:border-amber-500 outline-none transition-all duration-300 text-sm text-stone-900 dark:text-stone-100 placeholder:text-stone-400 resize-none"
          placeholder="Brief summary of the article..."
        ></textarea>
      </div>

      <!-- Cover Image -->
      <div>
        <label class="block text-sm font-medium text-stone-600 dark:text-stone-300 mb-1.5">Cover Image</label>
        <div class="flex gap-2">
          <input
            v-model="form.cover_url"
            type="text"
            class="flex-1 px-4 py-2.5 border border-stone-200 dark:border-stone-700 bg-white dark:bg-stone-900 rounded-lg focus:ring-2 focus:ring-amber-500/20 focus:border-amber-600 dark:focus:border-amber-500 outline-none transition-all duration-300 text-sm text-stone-900 dark:text-stone-100 placeholder:text-stone-400"
            placeholder="Paste URL or click Upload →"
          />
          <label class="cursor-pointer shrink-0 inline-flex items-center gap-1.5 px-4 py-2.5 bg-stone-100 dark:bg-stone-800 hover:bg-stone-200 dark:hover:bg-stone-700 text-stone-600 dark:text-stone-400 rounded-lg transition-colors duration-300 text-sm">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
            Upload
            <input type="file" accept="image/*" class="hidden" @change="handleCoverUpload">
          </label>
        </div>
        <div v-if="coverUploading" class="mt-3 w-full h-32 rounded-lg border border-stone-200 dark:border-stone-700 bg-stone-100 dark:bg-stone-800 flex items-center justify-center text-sm text-stone-400">Uploading...</div>
        <img v-else-if="form.cover_url" :src="form.cover_url" class="mt-3 w-full h-32 object-cover rounded-lg border border-stone-200 dark:border-stone-700" />
        <p class="text-[11px] text-stone-400 mt-1">Paste an image URL or upload from your computer. Leave empty for the default background.</p>
      </div>

      <!-- Category -->
      <div>
        <label class="block text-sm font-medium text-stone-600 dark:text-stone-300 mb-1.5">Category</label>
        <select
          v-model="form.category_id"
          class="w-full px-4 py-2.5 border border-stone-200 dark:border-stone-700 bg-white dark:bg-stone-900 rounded-lg focus:ring-2 focus:ring-amber-500/20 focus:border-amber-600 dark:focus:border-amber-500 outline-none transition-all duration-300 text-sm text-stone-900 dark:text-stone-100"
        >
          <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
        </select>
      </div>

      <!-- Content -->
      <div>
        <div class="flex items-center justify-between mb-1.5">
          <label class="block text-sm font-medium text-stone-600 dark:text-stone-300">Content (Markdown)</label>
          <label class="cursor-pointer text-xs bg-stone-100 dark:bg-stone-800 hover:bg-stone-200 dark:hover:bg-stone-700 text-stone-500 py-1.5 px-3 rounded-md inline-flex items-center transition-colors duration-300">
            <svg class="w-3.5 h-3.5 mr-1.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
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
              class="w-full px-4 py-3 border border-stone-200 dark:border-stone-700 bg-white dark:bg-stone-900 rounded-lg focus:ring-2 focus:ring-amber-500/20 focus:border-amber-600 dark:focus:border-amber-500 outline-none transition-all duration-300 font-mono text-sm resize-y text-stone-900 dark:text-stone-100 placeholder:text-stone-400"
              placeholder="Write your article content here in Markdown format..."
            ></textarea>
          </div>
          <div class="border border-stone-200 dark:border-stone-700 rounded-lg p-4 bg-stone-50/50 dark:bg-stone-900/50 overflow-y-auto prose prose-stone dark:prose-invert prose-sm max-w-none markdown-body" style="height: 480px;" v-html="compiledContent"></div>
        </div>
      </div>

      <!-- Actions -->
      <div class="flex items-center gap-3 pt-4 border-t border-stone-100 dark:border-stone-800">
        <button
          type="submit"
          :disabled="submitting"
          class="px-5 py-2.5 bg-amber-700 hover:bg-amber-800 text-white rounded-lg disabled:opacity-50 transition-colors duration-500 text-sm font-medium"
        >
          {{ submitting ? (isEditing ? 'Updating...' : 'Publishing...') : (isEditing ? 'Update Article' : 'Publish Article') }}
        </button>
        <router-link
          to="/admin"
          class="px-5 py-2.5 text-sm font-medium text-stone-500 hover:text-stone-800 dark:hover:text-stone-200 hover:bg-stone-100 dark:hover:bg-stone-800 rounded-lg transition-colors duration-300"
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
  cover_url: '',
  category_id: 1
})

const categories = ref([])
const loadingArticle = ref(false)
const submitting = ref(false)
const coverUploading = ref(false)
const errorMsg = ref('')
const successMsg = ref('')

const compiledContent = computed(() => {
  if (!form.value.content) return '<p class="text-stone-400 italic">Preview will appear here...</p>'
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

const handleCoverUpload = async (e) => {
  const file = e.target.files[0]
  if (!file) return
  e.target.value = ''

  coverUploading.value = true
  try {
    const res = await api.uploadImage(file)
    if (res.data.code === 200) {
      form.value.cover_url = res.data.data.url
    } else {
      alert(res.data.message || 'Image upload failed')
    }
  } catch (err) {
    console.error(err)
    alert('An error occurred during image upload')
  } finally {
    coverUploading.value = false
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
      } else {
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
  try {
    const res = await api.getCategories()
    if (res.data.code === 200) {
      categories.value = res.data.data || []
    }
  } catch (e) {
    console.error('Failed to load categories', e)
  }

  if (!isEditing.value) return
  loadingArticle.value = true
  try {
    const res = await api.getArticleDetails(route.params.id)
    if (res.data.code === 200) {
      const article = res.data.data
      form.value.title = article.title || ''
      form.value.summary = article.summary || ''
      form.value.content = article.content || ''
      form.value.cover_url = article.cover_url || ''
      form.value.category_id = article.category_id || 1
    }
  } catch (err) {
    errorMsg.value = err.response?.data?.message || err.message || 'Failed to load article'
  } finally {
    loadingArticle.value = false
  }
})
</script>
