import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ArticleView from '../views/ArticleView.vue'
import ProjectsView from '../views/ProjectsView.vue'
import AdminLogin from '../views/AdminLogin.vue'
import AdminLayout from '../views/AdminLayout.vue'
import AdminArticles from '../views/AdminArticles.vue'
import AdminArticleEditor from '../views/AdminArticleEditor.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/projects',
      name: 'projects',
      component: ProjectsView
    },
    {
      path: '/articles/:id',
      name: 'article',
      component: ArticleView,
      props: true
    },
    // Redirect /write to admin create page
    {
      path: '/write',
      redirect: '/admin/articles/create'
    },
    // Catch-all 404
    {
      path: '/:pathMatch(.*)*',
      redirect: '/'
    },
    // Admin routes
    {
      path: '/admin/login',
      name: 'admin-login',
      component: AdminLogin,
      meta: { title: 'Admin Login' }
    },
    {
      path: '/admin',
      component: AdminLayout,
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'admin',
          component: AdminArticles,
          meta: { title: 'Articles' }
        },
        {
          path: 'articles/create',
          name: 'admin-article-create',
          component: AdminArticleEditor,
          meta: { title: 'New Article' }
        },
        {
          path: 'articles/:id/edit',
          name: 'admin-article-edit',
          component: AdminArticleEditor,
          meta: { title: 'Edit Article' }
        }
      ]
    }
  ]
})

// Auth guard
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('admin_token')
  if (to.meta.requiresAuth && !token) {
    next({ name: 'admin-login' })
  } else if (to.name === 'admin-login' && token) {
    next({ name: 'admin' })
  } else {
    next()
  }
})

export default router
