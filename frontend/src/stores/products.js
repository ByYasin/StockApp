import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { 
  GetAllProducts, 
  GetProductsByCategory, 
  SearchProducts, 
  CreateProduct, 
  UpdateProduct, 
  DeleteProduct,
  GetLowStockProducts
} from '@wails/services/ProductService'

export const useProductStore = defineStore('products', () => {
  const products = ref([])
  const isLoading = ref(false)
  const error = ref(null)
  const searchQuery = ref('')
  const selectedCategoryId = ref(null)

  // Computed
  const filteredProducts = computed(() => {
    let result = products.value

    // Category filter
    if (selectedCategoryId.value) {
      result = result.filter(p => p.category_id === selectedCategoryId.value)
    }

    // Search filter
    if (searchQuery.value) {
      const query = searchQuery.value.toLowerCase()
      result = result.filter(p => 
        p.name.toLowerCase().includes(query) ||
        p.code?.toLowerCase().includes(query) ||
        p.description?.toLowerCase().includes(query)
      )
    }

    return result
  })

  const lowStockProducts = computed(() => {
    return products.value.filter(p => p.quantity <= p.minimum_stock)
  })

  // Actions
  async function loadProducts() {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await GetAllProducts()
      products.value = result || []
    } catch (err) {
      error.value = err.message || 'Ürünler yüklenemedi'
      console.error('Failed to load products:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function loadProductsByCategory(categoryId) {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await GetProductsByCategory(categoryId)
      products.value = result || []
      selectedCategoryId.value = categoryId
    } catch (err) {
      error.value = err.message || 'Ürünler yüklenemedi'
      console.error('Failed to load products by category:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function searchProductsAction(query) {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await SearchProducts(query)
      products.value = result || []
      searchQuery.value = query
    } catch (err) {
      error.value = err.message || 'Arama başarısız'
      console.error('Failed to search products:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function createProduct(product) {
    isLoading.value = true
    error.value = null
    
    try {
      await CreateProduct(
        product.code || '',
        product.name,
        product.description || '',
        product.category_id,
        product.quantity,
        product.unit,
        product.minimum_stock,
        product.unit_price
      )
      await loadProducts()
    } catch (err) {
      error.value = err.message || 'Ürün oluşturulamadı'
      console.error('Failed to create product:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function updateProduct(id, product) {
    isLoading.value = true
    error.value = null
    
    try {
      await UpdateProduct(
        id,
        product.code || '',
        product.name,
        product.description || '',
        product.category_id,
        product.unit,
        product.minimum_stock,
        product.unit_price
      )
      await loadProducts()
    } catch (err) {
      error.value = err.message || 'Ürün güncellenemedi'
      console.error('Failed to update product:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function deleteProduct(id) {
    isLoading.value = true
    error.value = null
    
    try {
      await DeleteProduct(id)
      await loadProducts()
    } catch (err) {
      error.value = err.message || 'Ürün silinemedi'
      console.error('Failed to delete product:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function loadLowStockProducts() {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await GetLowStockProducts()
      products.value = result || []
    } catch (err) {
      error.value = err.message || 'Düşük stok ürünleri yüklenemedi'
      console.error('Failed to load low stock products:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  function setSearchQuery(query) {
    searchQuery.value = query
  }

  function setSelectedCategory(categoryId) {
    selectedCategoryId.value = categoryId
  }

  function clearFilters() {
    searchQuery.value = ''
    selectedCategoryId.value = null
  }

  function clearError() {
    error.value = null
  }

  return {
    products,
    filteredProducts,
    lowStockProducts,
    isLoading,
    error,
    searchQuery,
    selectedCategoryId,
    loadProducts,
    loadProductsByCategory,
    searchProductsAction,
    createProduct,
    updateProduct,
    deleteProduct,
    loadLowStockProducts,
    setSearchQuery,
    setSelectedCategory,
    clearFilters,
    clearError
  }
})
