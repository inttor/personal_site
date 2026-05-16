<template>
  <div class="min-h-screen flex items-center justify-center bg-zinc-50">
    <div class="w-full max-w-sm mx-auto">
      <div class="bg-white rounded-2xl shadow-sm border border-zinc-100 p-8">
        <div class="text-center mb-8">
          <h1 class="text-2xl font-bold text-zinc-900">Admin</h1>
          <p class="text-sm text-zinc-500 mt-1">Sign in to manage articles</p>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-5">
          <div>
            <label class="block text-sm font-medium text-zinc-700 mb-1.5">Username</label>
            <input
              v-model="username"
              type="text"
              required
              autocomplete="username"
              class="w-full px-4 py-2.5 border border-zinc-200 rounded-lg focus:ring-2 focus:ring-zinc-900/10 focus:border-zinc-900 outline-none transition-all text-sm"
              placeholder="Enter username"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-zinc-700 mb-1.5">Password</label>
            <input
              v-model="password"
              type="password"
              required
              autocomplete="current-password"
              class="w-full px-4 py-2.5 border border-zinc-200 rounded-lg focus:ring-2 focus:ring-zinc-900/10 focus:border-zinc-900 outline-none transition-all text-sm"
              placeholder="Enter password"
            />
          </div>

          <p v-if="errorMsg" class="text-red-500 text-sm">{{ errorMsg }}</p>

          <button
            type="submit"
            :disabled="loading"
            class="w-full py-2.5 bg-zinc-900 text-white rounded-lg hover:bg-zinc-800 disabled:opacity-50 transition-colors text-sm font-medium"
          >
            {{ loading ? 'Signing in...' : 'Sign in' }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api'

const router = useRouter()
const username = ref('')
const password = ref('')
const loading = ref(false)
const errorMsg = ref('')

const handleLogin = async () => {
  loading.value = true
  errorMsg.value = ''

  try {
    const res = await api.login(username.value, password.value)
    if (res.data.code === 200) {
      const { token, user } = res.data.data
      localStorage.setItem('admin_token', token)
      localStorage.setItem('admin_user', JSON.stringify(user))
      router.push('/admin')
    } else {
      errorMsg.value = res.data.message || 'Login failed'
    }
  } catch (err) {
    errorMsg.value = err.response?.data?.message || err.message || 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>
