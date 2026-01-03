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
            <h1 class="text-2xl font-bold text-gray-800 dark:text-gray-100">Stok Hareketleri</h1>
          </div>
          <button
            @click="openCreateModal"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            + Yeni Hareket
          </button>
        </div>
      </div>
    </header>
    
    <main class="flex-1 overflow-auto p-6">
      <div class="max-w-7xl mx-auto">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
          <div class="card bg-white dark:bg-gray-800">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400">Toplam Giriş</p>
                <p class="text-2xl font-bold text-green-600 dark:text-green-400">{{ stats.totalIn }}</p>
              </div>
              <div class="p-3 bg-green-100 dark:bg-green-900 rounded-full">
                <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card bg-white dark:bg-gray-800">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400">Toplam Çıkış</p>
                <p class="text-2xl font-bold text-red-600 dark:text-red-400">{{ stats.totalOut }}</p>
              </div>
              <div class="p-3 bg-red-100 dark:bg-red-900 rounded-full">
                <svg class="w-6 h-6 text-red-600 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13l-5 5m0 0l-5-5m5 5V6"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card bg-white dark:bg-gray-800">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400">Toplam Hareket</p>
                <p class="text-2xl font-bold text-gray-800 dark:text-gray-100">{{ movementStore.movements.length }}</p>
              </div>
              <div class="p-3 bg-blue-100 dark:bg-blue-900 rounded-full">
                <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>

        <div class="card bg-white dark:bg-gray-800 mb-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <select
              v-model="filterProduct"
              class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">Tüm Ürünler</option>
              <option v-for="product in productStore.products" :key="product.id" :value="product.id">
                {{ product.name }}
              </option>
            </select>
            
            <select
              v-model="filterType"
              class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">Tüm Hareketler</option>
              <option value="IN">Sadece Giriş</option>
              <option value="OUT">Sadece Çıkış</option>
            </select>
            
            <input
              v-model="filterDate"
              type="date"
              class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>

        <div class="card bg-white dark:bg-gray-800">
          <div v-if="movementStore.loading" class="text-center py-12">
            <p class="text-gray-500 dark:text-gray-400">Yükleniyor...</p>
          </div>
          
          <div v-else-if="filteredMovements.length === 0" class="text-center py-12">
            <p class="text-gray-500 dark:text-gray-400">Hareket kaydı bulunamadı</p>
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
                  <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase">İşlemler</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200 dark:divide-gray-700">
                <tr v-for="movement in filteredMovements" :key="movement.id" class="hover:bg-gray-50 dark:hover:bg-gray-700">
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
                  <td class="px-4 py-3 text-sm font-medium" :class="movement.type === 'IN' ? 'text-green-600 dark:text-green-400' : 'text-red-600 dark:text-red-400'">
                    {{ movement.type === 'IN' ? '+' : '-' }}{{ movement.quantity }}
                  </td>
                  <td class="px-4 py-3 text-sm text-gray-500 dark:text-gray-400">
                    {{ movement.note || '-' }}
                  </td>
                  <td class="px-4 py-3 text-sm">
                    <button
                      @click="confirmDelete(movement)"
                      class="text-red-600 dark:text-red-400 hover:text-red-800 dark:hover:text-red-300"
                      title="Sil"
                    >
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </main>

    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 max-w-md w-full mx-4">
        <h2 class="text-xl font-bold text-gray-800 dark:text-gray-100 mb-4">Yeni Hareket</h2>
        
        <form @submit.prevent="saveMovement" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Ürün *</label>
            <select
              v-model="formData.product_id"
              required
              @change="onProductChange"
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">Ürün Seçin</option>
              <option v-for="product in productStore.products" :key="product.id" :value="product.id">
                {{ product.name }} (Mevcut: {{ product.current_stock }})
              </option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Hareket Tipi *</label>
            <div class="grid grid-cols-2 gap-3">
              <button
                type="button"
                @click="formData.type = 'IN'"
                class="px-4 py-2 rounded-lg border-2 transition-colors"
                :class="formData.type === 'IN' 
                  ? 'border-green-500 bg-green-50 dark:bg-green-900/20 text-green-700 dark:text-green-400 font-medium' 
                  : 'border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:border-green-300 dark:hover:border-green-700'"
              >
                <svg class="w-5 h-5 mx-auto mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"></path>
                </svg>
                Giriş
              </button>
              <button
                type="button"
                @click="formData.type = 'OUT'"
                class="px-4 py-2 rounded-lg border-2 transition-colors"
                :class="formData.type === 'OUT' 
                  ? 'border-red-500 bg-red-50 dark:bg-red-900/20 text-red-700 dark:text-red-400 font-medium' 
                  : 'border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:border-red-300 dark:hover:border-red-700'"
              >
                <svg class="w-5 h-5 mx-auto mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13l-5 5m0 0l-5-5m5 5V6"></path>
                </svg>
                Çıkış
              </button>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Miktar *</label>
            <input
              v-model.number="formData.quantity"
              type="number"
              required
              min="1"
              :max="formData.type === 'OUT' ? selectedProductStock : undefined"
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
            <p v-if="formData.type === 'OUT' && selectedProductStock > 0" class="text-xs text-gray-500 dark:text-gray-400 mt-1">
              Maksimum çıkış: {{ selectedProductStock }}
            </p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Not</label>
            <textarea
              v-model="formData.note"
              rows="3"
              placeholder="Hareket hakkında not..."
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            ></textarea>
          </div>
          
          <div class="flex space-x-3 pt-4">
            <button
              type="submit"
              class="flex-1 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              Kaydet
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
import { ref, computed, onMounted } from 'vue'
import { useMovementStore } from '@/stores/movements'
import { useProductStore } from '@/stores/products'

const movementStore = useMovementStore()
const productStore = useProductStore()

const showModal = ref(false)
const filterProduct = ref('')
const filterType = ref('')
const filterDate = ref('')

const formData = ref({
  product_id: '',
  type: 'IN',
  quantity: 1,
  note: ''
})

const selectedProductStock = ref(0)

const stats = computed(() => {
  const totalIn = movementStore.movements
    .filter(m => m.type === 'IN')
    .reduce((sum, m) => sum + m.quantity, 0)
  
  const totalOut = movementStore.movements
    .filter(m => m.type === 'OUT')
    .reduce((sum, m) => sum + m.quantity, 0)
  
  return { totalIn, totalOut }
})

const filteredMovements = computed(() => {
  let filtered = [...movementStore.movements]
  
  if (filterProduct.value) {
    filtered = filtered.filter(m => m.product_id === filterProduct.value)
  }
  
  if (filterType.value) {
    filtered = filtered.filter(m => m.type === filterType.value)
  }
  
  if (filterDate.value) {
    filtered = filtered.filter(m => {
      const movementDate = new Date(m.created_at).toISOString().split('T')[0]
      return movementDate === filterDate.value
    })
  }
  
  return filtered.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
})

const openCreateModal = () => {
  formData.value = {
    product_id: '',
    type: 'IN',
    quantity: 1,
    note: ''
  }
  selectedProductStock.value = 0
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}

const onProductChange = () => {
  const product = productStore.products.find(p => p.id === formData.value.product_id)
  selectedProductStock.value = product ? product.current_stock : 0
}

const saveMovement = async () => {
  try {
    if (formData.value.type === 'OUT' && formData.value.quantity > selectedProductStock.value) {
      alert(`Yetersiz stok! Mevcut stok: ${selectedProductStock.value}`)
      return
    }
    
    await movementStore.createMovement(formData.value)
    await productStore.loadProducts()
    closeModal()
  } catch (err) {
    alert('Hata: ' + err.message)
  }
}

const confirmDelete = async (movement) => {
  if (confirm('Bu hareketi silmek istediğinize emin misiniz? Bu işlem geri alınamaz.')) {
    try {
      await movementStore.deleteMovement(movement.id)
      await productStore.loadProducts()
    } catch (err) {
      alert('Hata: ' + err.message)
    }
  }
}

const getProductName = (productId) => {
  const product = productStore.products.find(p => p.id === productId)
  return product ? product.name : 'Bilinmeyen Ürün'
}

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

onMounted(async () => {
  try {
    await Promise.all([
      movementStore.loadMovements(),
      productStore.loadProducts()
    ])
  } catch (err) {
    console.error('Failed to load movements:', err)
  }
})
</script>
