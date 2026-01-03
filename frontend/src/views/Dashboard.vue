<template>
  <div class="h-screen flex flex-col bg-gray-50 dark:bg-gray-900">
    <header class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4">
            <h1 class="text-2xl font-bold text-gray-800 dark:text-gray-100">Stok Takip Sistemi</h1>
            <span class="px-3 py-1 text-sm bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full">
              {{ dbStore.activeDatabaseName }}
            </span>
          </div>
          
          <div class="flex items-center space-x-4">
            <button
              @click="themeStore.toggleTheme()"
              class="p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
              title="Tema Değiştir"
            >
              <svg v-if="themeStore.isDark" class="w-6 h-6 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z" clip-rule="evenodd"></path>
              </svg>
              <svg v-else class="w-6 h-6 text-gray-600 dark:text-gray-300" fill="currentColor" viewBox="0 0 20 20">
                <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"></path>
              </svg>
            </button>

            <button
              @click="$router.push('/db-selector')"
              class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors flex items-center space-x-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2 1 3 3 3h10c2 0 3-1 3-3V7M4 7V5c0-2 1-3 3-3h10c2 0 3 1 3 3v2M4 7h16M10 11v6m4-6v6"></path>
              </svg>
              <span>Veritabanı Değiştir</span>
            </button>

            <nav class="flex space-x-1">
              <router-link
                v-for="item in navItems"
                :key="item.name"
                :to="{ name: item.route }"
                class="px-4 py-2 rounded-lg text-sm font-medium transition-colors"
                :class="currentRoute === item.route 
                  ? 'bg-blue-600 text-white' 
                  : 'text-gray-600 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'"
              >
                {{ item.label }}
              </router-link>
            </nav>
          </div>
        </div>
      </div>
    </header>

    <main class="flex-1 overflow-auto p-6">
      <div class="max-w-7xl mx-auto">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
          <div class="card bg-white dark:bg-gray-800">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400">Toplam Ürün</p>
                <p class="text-3xl font-bold text-gray-800 dark:text-gray-100">{{ totalProducts }}</p>
              </div>
              <div class="p-3 bg-blue-100 dark:bg-blue-900 rounded-full">
                <svg class="w-8 h-8 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card bg-white dark:bg-gray-800">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400">Kritik Stok</p>
                <p class="text-3xl font-bold text-red-600 dark:text-red-400">{{ lowStockCount }}</p>
              </div>
              <div class="p-3 bg-red-100 dark:bg-red-900 rounded-full">
                <svg class="w-8 h-8 text-red-600 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card bg-white dark:bg-gray-800">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400">Kategoriler</p>
                <p class="text-3xl font-bold text-gray-800 dark:text-gray-100">{{ totalCategories }}</p>
              </div>
              <div class="p-3 bg-green-100 dark:bg-green-900 rounded-full">
                <svg class="w-8 h-8 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card bg-white dark:bg-gray-800">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400">Bugünkü Hareketler</p>
                <p class="text-3xl font-bold text-gray-800 dark:text-gray-100">{{ todayMovements }}</p>
              </div>
              <div class="p-3 bg-purple-100 dark:bg-purple-900 rounded-full">
                <svg class="w-8 h-8 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>

        <div class="card bg-white dark:bg-gray-800 mb-6">
          <h2 class="text-2xl font-bold text-gray-800 dark:text-gray-100 mb-4">Hoş Geldiniz!</h2>
          <p class="text-gray-600 dark:text-gray-400 mb-6">
            Stok takip sisteminize başarıyla bağlandınız. Yukarıdaki menüden istediğiniz sayfaya geçebilirsiniz.
          </p>
          
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <router-link
              :to="{ name: 'Products' }"
              class="p-4 border-2 border-gray-200 dark:border-gray-700 rounded-lg hover:border-blue-500 hover:bg-blue-50 dark:hover:bg-blue-900/20 transition-all cursor-pointer group"
            >
              <div class="flex items-center space-x-3">
                <div class="p-2 bg-blue-100 dark:bg-blue-900 rounded-lg group-hover:bg-blue-200 dark:group-hover:bg-blue-800 transition-colors">
                  <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-800 dark:text-gray-100">Ürünler</p>
                  <p class="text-sm text-gray-500 dark:text-gray-400">Ürün yönetimi</p>
                </div>
              </div>
            </router-link>

            <router-link
              :to="{ name: 'Movements' }"
              class="p-4 border-2 border-gray-200 dark:border-gray-700 rounded-lg hover:border-green-500 hover:bg-green-50 dark:hover:bg-green-900/20 transition-all cursor-pointer group"
            >
              <div class="flex items-center space-x-3">
                <div class="p-2 bg-green-100 dark:bg-green-900 rounded-lg group-hover:bg-green-200 dark:group-hover:bg-green-800 transition-colors">
                  <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-800 dark:text-gray-100">Stok Hareketi</p>
                  <p class="text-sm text-gray-500 dark:text-gray-400">Giriş/Çıkış işlemleri</p>
                </div>
              </div>
            </router-link>

            <router-link
              :to="{ name: 'Categories' }"
              class="p-4 border-2 border-gray-200 dark:border-gray-700 rounded-lg hover:border-purple-500 hover:bg-purple-50 dark:hover:bg-purple-900/20 transition-all cursor-pointer group"
            >
              <div class="flex items-center space-x-3">
                <div class="p-2 bg-purple-100 dark:bg-purple-900 rounded-lg group-hover:bg-purple-200 dark:group-hover:bg-purple-800 transition-colors">
                  <svg class="w-6 h-6 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-800 dark:text-gray-100">Kategoriler</p>
                  <p class="text-sm text-gray-500 dark:text-gray-400">Kategori yönetimi</p>
                </div>
              </div>
            </router-link>
          </div>
        </div>

        <div class="card bg-white dark:bg-gray-800">
          <h3 class="text-lg font-semibold mb-4 text-gray-800 dark:text-gray-100">Son Hareketler</h3>
          
          <div v-if="recentMovements.length === 0" class="text-center py-8 text-gray-500 dark:text-gray-400">
            <p>Henüz hareket kaydı yok</p>
          </div>
          
          <div v-else class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
              <thead>
                <tr>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Tarih</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Ürün</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Tip</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Miktar</th>
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Not</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200 dark:divide-gray-700">
                <tr v-for="movement in recentMovements" :key="movement.id" class="hover:bg-gray-50 dark:hover:bg-gray-700">
                  <td class="px-4 py-3 text-sm text-gray-600 dark:text-gray-300">
                    {{ formatDate(movement.created_at) }}
                  </td>
                  <td class="px-4 py-3 text-sm font-medium text-gray-800 dark:text-gray-200">
                    {{ getProductName(movement.product_id) }}
                  </td>
                  <td class="px-4 py-3">
                    <span
                      class="px-2 py-1 text-xs font-medium rounded-full"
                      :class="movement.type === 'IN'
                        ? 'bg-green-100 dark:bg-green-900 text-green-800 dark:text-green-200'
                        : 'bg-red-100 dark:bg-red-900 text-red-800 dark:text-red-200'"
                    >
                      {{ movement.type === 'IN' ? 'Giriş' : 'Çıkış' }}
                    </span>
                  </td>
                  <td class="px-4 py-3 text-sm text-gray-600 dark:text-gray-300">
                    {{ movement.quantity }}
                  </td>
                  <td class="px-4 py-3 text-sm text-gray-500 dark:text-gray-400">
                    {{ movement.note || '-' }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useDatabaseStore } from '@/stores/database'
import { useProductStore } from '@/stores/products'
import { useCategoryStore } from '@/stores/categories'
import { useMovementStore } from '@/stores/movements'
import { useThemeStore } from '@/stores/theme'

const route = useRoute()
const dbStore = useDatabaseStore()
const productStore = useProductStore()
const categoryStore = useCategoryStore()
const movementStore = useMovementStore()
const themeStore = useThemeStore()

const currentRoute = computed(() => route.name)

const navItems = [
  { label: 'Ana Sayfa', route: 'Dashboard' },
  { label: 'Ürünler', route: 'Products' },
  { label: 'Kategoriler', route: 'Categories' },
  { label: 'Hareketler', route: 'Movements' }
]

const totalProducts = computed(() => productStore.products.length)
const lowStockCount = computed(() => productStore.lowStockProducts.length)
const totalCategories = computed(() => categoryStore.categories.length)
const todayMovements = computed(() => {
  const today = new Date().toDateString()
  return movementStore.movements.filter(m => {
    const movementDate = new Date(m.created_at).toDateString()
    return movementDate === today
  }).length
})

const recentMovements = computed(() => {
  return [...movementStore.movements]
    .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
    .slice(0, 10)
})

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  return date.toLocaleDateString('tr-TR', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getProductName = (productId) => {
  const product = productStore.products.find(p => p.id === productId)
  return product ? product.name : 'Bilinmeyen Ürün'
}

onMounted(async () => {
  try {
    await Promise.all([
      productStore.loadProducts(),
      categoryStore.loadCategories(),
      movementStore.loadMovements()
    ])
  } catch (err) {
    console.error('Failed to load dashboard data:', err)
  }
})
</script>
