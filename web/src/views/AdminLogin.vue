<template>
  <div class="min-h-screen flex items-center justify-center bg-stone-50 dark:bg-stone-950 transition-colors duration-500">
    <div class="w-full max-w-sm mx-auto">
      <div class="bg-white dark:bg-stone-900 rounded-2xl shadow-sm border border-stone-200 dark:border-stone-800 p-8 transition-colors duration-500">
        <div class="text-center mb-8">
          <h1 class="font-display text-2xl font-bold text-stone-900 dark:text-stone-100 mb-1">Admin</h1>
          <p class="text-sm text-stone-400">Sign in to manage articles</p>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-5">
          <div>
            <label class="block text-sm font-medium text-stone-600 dark:text-stone-300 mb-1.5">Username</label>
            <input
              v-model="username"
              type="text"
              required
              autocomplete="username"
              class="w-full px-4 py-2.5 border border-stone-200 dark:border-stone-700 bg-white dark:bg-stone-900 rounded-lg focus:ring-2 focus:ring-amber-500/20 focus:border-amber-600 dark:focus:border-amber-500 outline-none transition-all duration-300 text-sm text-stone-900 dark:text-stone-100 placeholder:text-stone-400"
              placeholder="Enter username"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-stone-600 dark:text-stone-300 mb-1.5">Password</label>
            <input
              v-model="password"
              type="password"
              required
              autocomplete="current-password"
              class="w-full px-4 py-2.5 border border-stone-200 dark:border-stone-700 bg-white dark:bg-stone-900 rounded-lg focus:ring-2 focus:ring-amber-500/20 focus:border-amber-600 dark:focus:border-amber-500 outline-none transition-all duration-300 text-sm text-stone-900 dark:text-stone-100 placeholder:text-stone-400"
              placeholder="Enter password"
            />
          </div>

          <p v-if="errorMsg" class="text-red-500 text-sm">{{ errorMsg }}</p>

          <button
            type="submit"
            :disabled="loading"
            class="w-full py-2.5 bg-amber-700 hover:bg-amber-800 text-white rounded-lg disabled:opacity-50 transition-colors duration-500 text-sm font-medium"
          >
            {{ loading ? 'Signing in...' : 'Sign in' }}
          </button>
        </form>
      </div>

      <p class="text-center text-xs text-stone-400 mt-6">
        <router-link to="/" class="hover:text-stone-600 dark:hover:text-stone-400 transition-colors duration-500">&larr; Back to site</router-link>
      </p>
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
