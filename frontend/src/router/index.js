import { createRouter, createWebHashHistory } from 'vue-router'
import { useDatabaseStore } from '@/stores/database'

// Views
import DatabaseSelector from '@/views/DatabaseSelector.vue'
import Dashboard from '@/views/Dashboard.vue'

const routes = [
  {
    path: '/',
    name: 'DatabaseSelector',
    component: DatabaseSelector
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresDB: true }
  },
  {
    path: '/products',
    name: 'Products',
    component: () => import('@/views/Products.vue'),
    meta: { requiresDB: true }
  },
  {
    path: '/categories',
    name: 'Categories',
    component: () => import('@/views/Categories.vue'),
    meta: { requiresDB: true }
  },
  {
    path: '/movements',
    name: 'Movements',
    component: () => import('@/views/Movements.vue'),
    meta: { requiresDB: true }
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const dbStore = useDatabaseStore()
  
  if (to.meta.requiresDB && !dbStore.isConnected) {
    next({ name: 'DatabaseSelector' })
  } else {
    next()
  }
})

export default router
