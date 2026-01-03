<template>
  <div class="h-screen flex flex-col bg-gray-50 dark:bg-gray-900">
    <header class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4">
            <button
              @click="$router.push('/dashboard')"
              class="p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg transition-colors"
              title="Ana Sayfaya Dön"
            >
              <svg class="w-6 h-6 text-gray-600 dark:text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
            </button>
            <h1 class="text-2xl font-bold text-gray-800 dark:text-gray-100">Ürün Yönetimi</h1>
          </div>
          <button
            @click="openCreateModal"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            + Yeni Ürün
          </button>
        </div>
      </div>
    </header>
    
    <main class="flex-1 overflow-auto p-6">
      <div class="max-w-7xl mx-auto">
        <div class="card bg-white dark:bg-gray-800 mb-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <input
              v-model="productStore.searchQuery"
              type="text"
              placeholder="Ürün adı veya kodu ara..."
              class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
            
            <select
              v-model="productStore.selectedCategory"
              class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">Tüm Kategoriler</option>
              <option v-for="cat in categoryStore.categories" :key="cat.id" :value="cat.id">
                {{ cat.name }}
              </option>
            </select>
            
            <select
              v-model="productStore.stockFilter"
              class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="all">Tüm Ürünler</option>
              <option value="low">Kritik Stok</option>
              <option value="out">Stokta Yok</option>
            </select>
          </div>
        </div>

        <div class="card bg-white dark:bg-gray-800">
          <div v-if="productStore.loading" class="text-center py-12">
            <p class="text-gray-500 dark:text-gray-400">Yükleniyor...</p>
          </div>
          
          <div v-else-if="filteredProducts.length === 0" class="text-center py-12">
            <p class="text-gray-500 dark:text-gray-400">Ürün bulunamadı</p>
          </div>
          
          <div v-else>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead>
                  <tr>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Kod</th>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Ürün Adı</th>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Kategori</th>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Birim</th>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Stok</th>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Kritik Seviye</th>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">Fiyat</th>
                    <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">İşlemler</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-200 dark:divide-gray-700">
                  <tr v-for="product in paginatedProducts" :key="product.id" class="hover:bg-gray-50 dark:hover:bg-gray-700">
                  <td class="px-4 py-3 text-sm font-mono text-gray-600 dark:text-gray-300">{{ product.code }}</td>
                  <td class="px-4 py-3 text-sm font-medium text-gray-800 dark:text-gray-200">{{ product.name }}</td>
                  <td class="px-4 py-3 text-sm">
                    <span 
                      class="px-2 py-1 text-xs rounded-full"
                      :style="{ 
                        backgroundColor: getCategoryColor(product.category_id) + '20',
                        color: getCategoryColor(product.category_id)
                      }"
                    >
                      {{ getCategoryName(product.category_id) }}
                    </span>
                  </td>
                  <td class="px-4 py-3 text-sm text-gray-600 dark:text-gray-300">{{ product.unit }}</td>
                  <td class="px-4 py-3 text-sm">
                    <span 
                      class="font-medium"
                      :class="getStockClass(product.current_stock, product.critical_limit)"
                    >
                      {{ product.current_stock }}
                    </span>
                  </td>
                  <td class="px-4 py-3 text-sm text-gray-600 dark:text-gray-300">{{ product.critical_limit }}</td>
                  <td class="px-4 py-3 text-sm text-gray-600 dark:text-gray-300">{{ formatPrice(product.price) }}</td>
                  <td class="px-4 py-3 text-sm">
                    <div class="flex space-x-2">
                      <button
                        @click="openEditModal(product)"
                        class="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300"
                        title="Düzenle"
                      >
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                        </svg>
                      </button>
                      <button
                        @click="confirmDelete(product)"
                        class="text-red-600 dark:text-red-400 hover:text-red-800 dark:hover:text-red-300"
                        title="Sil"
                      >
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                        </svg>
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          
          <!-- Pagination Controls -->
          <div v-if="totalPages > 1" class="mt-4 flex items-center justify-between border-t border-gray-200 dark:border-gray-700 pt-4">
            <div class="text-sm text-gray-600 dark:text-gray-400">
              {{ (currentPage - 1) * itemsPerPage + 1 }}-{{ Math.min(currentPage * itemsPerPage, filteredProducts.length) }} / {{ filteredProducts.length }} ürün
            </div>
            
            <div class="flex items-center space-x-2">
              <button
                @click="prevPage"
                :disabled="currentPage === 1"
                :class="[
                  'px-3 py-1 rounded-lg border transition-colors',
                  currentPage === 1
                    ? 'border-gray-200 dark:border-gray-700 text-gray-400 dark:text-gray-600 cursor-not-allowed'
                    : 'border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'
                ]"
              >
                ← Önceki
              </button>
              
              <div class="flex items-center space-x-1">
                <button
                  v-for="page in totalPages"
                  :key="page"
                  @click="goToPage(page)"
                  :class="[
                    'px-3 py-1 rounded-lg transition-colors',
                    page === currentPage
                      ? 'bg-blue-600 text-white'
                      : 'text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'
                  ]"
                >
                  {{ page }}
                </button>
              </div>
              
              <button
                @click="nextPage"
                :disabled="currentPage === totalPages"
                :class="[
                  'px-3 py-1 rounded-lg border transition-colors',
                  currentPage === totalPages
                    ? 'border-gray-200 dark:border-gray-700 text-gray-400 dark:text-gray-600 cursor-not-allowed'
                    : 'border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'
                ]"
              >
                Sonraki →
              </button>
            </div>
          </div>
        </div>
      </div>
      </div>
    </main>

    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 max-w-md w-full mx-4">
        <h2 class="text-xl font-bold text-gray-800 dark:text-gray-100 mb-4">
          {{ editingProduct ? 'Ürün Düzenle' : 'Yeni Ürün' }}
        </h2>
        
        <form @submit.prevent="saveProduct" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Ürün Kodu *</label>
            <input
              v-model="formData.code"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Ürün Adı *</label>
            <input
              v-model="formData.name"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Kategori *</label>
            <select
              v-model="formData.category_id"
              required
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">Kategori Seçin</option>
              <option v-for="cat in categoryStore.categories" :key="cat.id" :value="cat.id">
                {{ cat.name }}
              </option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Birim *</label>
            <input
              v-model="formData.unit"
              type="text"
              required
              placeholder="Adet, Kg, Lt vb."
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Kritik Stok Seviyesi *</label>
            <input
              v-model.number="formData.critical_limit"
              type="number"
              required
              min="0"
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Birim Fiyat *</label>
            <input
              v-model.number="formData.price"
              type="number"
              required
              min="0"
              step="0.01"
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div class="flex space-x-3 pt-4">
            <button
              type="submit"
              class="flex-1 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              {{ editingProduct ? 'Güncelle' : 'Kaydet' }}
            </button>
            <button
              type="button"
              @click="closeModal"
              class="flex-1 px-4 py-2 bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
            >
              İptal
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useProductStore } from '@/stores/products'
import { useCategoryStore } from '@/stores/categories'

const productStore = useProductStore()
const categoryStore = useCategoryStore()

const showModal = ref(false)
const editingProduct = ref(null)
const formData = ref({
  code: '',
  name: '',
  category_id: '',
  unit: '',
  critical_limit: 0,
  price: 0
})

// Pagination
const currentPage = ref(1)
const itemsPerPage = 25

const filteredProducts = computed(() => productStore.filteredProducts)

const totalPages = computed(() => {
  return Math.ceil(filteredProducts.value.length / itemsPerPage)
})

const paginatedProducts = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredProducts.value.slice(start, end)
})

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

// Watch for filter changes and reset to page 1
watch([
  () => productStore.searchQuery,
  () => productStore.selectedCategory,
  () => productStore.stockFilter
], () => {
  currentPage.value = 1
})

const openCreateModal = () => {
  editingProduct.value = null
  formData.value = {
    code: '',
    name: '',
    category_id: '',
    unit: '',
    critical_limit: 0,
    price: 0
  }
  showModal.value = true
}

const openEditModal = (product) => {
  editingProduct.value = product
  formData.value = {
    code: product.code,
    name: product.name,
    category_id: product.category_id,
    unit: product.unit,
    critical_limit: product.critical_limit,
    price: product.price
  }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  editingProduct.value = null
}

const saveProduct = async () => {
  try {
    if (editingProduct.value) {
      await productStore.updateProduct(editingProduct.value.id, formData.value)
    } else {
      await productStore.createProduct(formData.value)
    }
    closeModal()
  } catch (err) {
    alert('Hata: ' + err.message)
  }
}

const confirmDelete = async (product) => {
  if (confirm(`"${product.name}" ürününü silmek istediğinize emin misiniz?`)) {
    try {
      await productStore.deleteProduct(product.id)
    } catch (err) {
      alert('Hata: ' + err.message)
    }
  }
}

const getCategoryName = (categoryId) => {
  const cat = categoryStore.categories.find(c => c.id === categoryId)
  return cat ? cat.name : '-'
}

const getCategoryColor = (categoryId) => {
  const cat = categoryStore.categories.find(c => c.id === categoryId)
  return cat ? cat.color : '#6B7280'
}

const getStockClass = (stock, criticalLimit) => {
  if (stock === 0) return 'text-red-600 dark:text-red-400'
  if (stock <= criticalLimit) return 'text-orange-600 dark:text-orange-400'
  return 'text-green-600 dark:text-green-400'
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('tr-TR', {
    style: 'currency',
    currency: 'TRY'
  }).format(price)
}

onMounted(async () => {
  try {
    await Promise.all([
      productStore.loadProducts(),
      categoryStore.loadCategories()
    ])
  } catch (err) {
    console.error('Failed to load products:', err)
  }
})
</script>
