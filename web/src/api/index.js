import axios from 'axios'

const apiClient = axios.create({
  baseURL: '/api/v1',
  headers: {
    'Content-Type': 'application/json'
  }
})

// Attach JWT token if available
apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('admin_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Redirect to login on 401
apiClient.interceptors.response.use(
  (res) => res,
  (err) => {
    if (err.response?.status === 401 && !err.config.url?.includes('/auth/login')) {
      localStorage.removeItem('admin_token')
      localStorage.removeItem('admin_user')
      window.location.href = '/admin/login'
    }
    return Promise.reject(err)
  }
)

export default {
  // Auth
  login(username, password) {
    return apiClient.post('/auth/login', { username, password })
  },

  // Articles
  getArticles(page = 1, size = 10, categoryId, tagId) {
    return apiClient.get('/articles', {
      params: { page, size, category_id: categoryId, tag_id: tagId }
    })
  },
  getArticleDetails(id) {
    return apiClient.get(`/articles/${id}`)
  },
  createArticle(data) {
    return apiClient.post('/articles', data)
  },
  updateArticle(id, data) {
    return apiClient.put(`/articles/${id}`, data)
  },
  deleteArticle(id) {
    return apiClient.delete(`/articles/${id}`)
  },

  // Categories & Tags
  getCategories() {
    return apiClient.get('/categories')
  },
  getTags() {
    return apiClient.get('/tags')
  },

  // Upload
  uploadImage(file) {
    const formData = new FormData()
    formData.append('file', file)
    return apiClient.post('/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  }
}
