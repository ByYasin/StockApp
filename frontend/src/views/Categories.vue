<template>
  <div class="h-screen flex flex-col bg-gray-50">
    <header class="bg-white shadow-sm border-b border-gray-200">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4">
            <button
              @click="$router.push('/')"
              class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
              title="Ana Sayfaya Dön"
            >
              <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
            </button>
            <h1 class="text-2xl font-bold text-gray-800">Kategori Yönetimi</h1>
          </div>
          <button
            @click="openCreateModal"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            + Yeni Kategori
          </button>
        </div>
      </div>
    </header>
    
    <main class="flex-1 overflow-auto p-6">
      <div class="max-w-4xl mx-auto">
        <!-- Categories Grid -->
        <div v-if="categoryStore.loading" class="card text-center py-12">
          <p class="text-gray-500">Yükleniyor...</p>
        </div>
        
        <div v-else-if="categoryStore.categories.length === 0" class="card text-center py-12">
          <p class="text-gray-500">Henüz kategori eklenmemiş</p>
          <button
            @click="openCreateModal"
            class="mt-4 text-blue-600 hover:underline"
          >
            İlk kategoriyi ekle
          </button>
        </div>
        
        <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div
            v-for="category in categoryStore.categories"
            :key="category.id"
            class="card hover:shadow-lg transition-shadow"
          >
            <div class="flex items-start justify-between">
              <div class="flex items-center space-x-3 flex-1">
                <div
                  class="w-12 h-12 rounded-lg flex items-center justify-center"
                  :style="{ backgroundColor: category.color + '20' }"
                >
                  <svg 
                    class="w-6 h-6" 
                    :style="{ color: category.color }"
                    fill="none" 
                    stroke="currentColor" 
                    viewBox="0 0 24 24"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                  </svg>
                </div>
                
                <div class="flex-1">
                  <h3 class="text-lg font-semibold text-gray-800">{{ category.name }}</h3>
                  <p class="text-sm text-gray-500">
                    {{ getProductCount(category.id) }} ürün
                  </p>
                </div>
              </div>
              
              <div class="flex space-x-2">
                <button
                  @click="openEditModal(category)"
                  class="text-blue-600 hover:text-blue-800 p-2"
                  title="Düzenle"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button
                  @click="confirmDelete(category)"
                  class="text-red-600 hover:text-red-800 p-2"
                  title="Sil"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
            
            <div v-if="category.description" class="mt-3 text-sm text-gray-600">
              {{ category.description }}
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Create/Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4">
        <h2 class="text-xl font-bold text-gray-800 mb-4">
          {{ editingCategory ? 'Kategori Düzenle' : 'Yeni Kategori' }}
        </h2>
        
        <form @submit.prevent="saveCategory" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Kategori Adı *</label>
            <input
              v-model="formData.name"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Açıklama</label>
            <textarea
              v-model="formData.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            ></textarea>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Renk *</label>
            <div class="flex items-center space-x-3">
              <input
                v-model="formData.color"
                type="color"
                required
                class="w-16 h-10 border border-gray-300 rounded cursor-pointer"
              />
              <input
                v-model="formData.color"
                type="text"
                required
                pattern="^#[0-9A-Fa-f]{6}$"
                placeholder="#3B82F6"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent font-mono"
              />
            </div>
            <p class="text-xs text-gray-500 mt-1">Hex renk kodu (örn: #3B82F6)</p>
          </div>
          
          <div class="flex space-x-3 pt-4">
            <button
              type="submit"
              class="flex-1 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              {{ editingCategory ? 'Güncelle' : 'Kaydet' }}
            </button>
            <button
              type="button"
              @click="closeModal"
              class="flex-1 px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 transition-colors"
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
import { ref, onMounted } from 'vue'
import { useCategoryStore } from '@/stores/categories'
import { useProductStore } from '@/stores/products'

const categoryStore = useCategoryStore()
const productStore = useProductStore()

const showModal = ref(false)
const editingCategory = ref(null)
const formData = ref({
  name: '',
  description: '',
  color: '#3B82F6'
})

const openCreateModal = () => {
  editingCategory.value = null
  formData.value = {
    name: '',
    description: '',
    color: '#3B82F6'
  }
  showModal.value = true
}

const openEditModal = (category) => {
  editingCategory.value = category
  formData.value = {
    name: category.name,
    description: category.description || '',
    color: category.color
  }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  editingCategory.value = null
}

const saveCategory = async () => {
  try {
    if (editingCategory.value) {
      await categoryStore.updateCategory(editingCategory.value.id, formData.value)
    } else {
      await categoryStore.createCategory(formData.value)
    }
    closeModal()
  } catch (err) {
    alert('Hata: ' + err.message)
  }
}

const confirmDelete = async (category) => {
  const productCount = getProductCount(category.id)
  
  if (productCount > 0) {
    alert(`Bu kategori ${productCount} ürün tarafından kullanılıyor. Önce ürünleri başka kategorilere taşıyın.`)
    return
  }
  
  if (confirm(`"${category.name}" kategorisini silmek istediğinize emin misiniz?`)) {
    try {
      await categoryStore.deleteCategory(category.id)
    } catch (err) {
      alert('Hata: ' + err.message)
    }
  }
}

const getProductCount = (categoryId) => {
  return productStore.products.filter(p => p.category_id === categoryId).length
}

onMounted(async () => {
  try {
    await Promise.all([
      categoryStore.loadCategories(),
      productStore.loadProducts()
    ])
  } catch (err) {
    console.error('Failed to load categories:', err)
  }
})
</script>
