import { defineStore } from 'pinia'
import { ref } from 'vue'
import { GetAllCategories, CreateCategory, UpdateCategory, DeleteCategory } from '@wails/services/CategoryService'

export const useCategoryStore = defineStore('categories', () => {
  const categories = ref([])
  const isLoading = ref(false)
  const error = ref(null)

  async function loadCategories() {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await GetAllCategories()
      categories.value = result || []
    } catch (err) {
      error.value = err.message || 'Kategoriler yüklenemedi'
      console.error('Failed to load categories:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function createCategory(category) {
    isLoading.value = true
    error.value = null
    
    try {
      await CreateCategory(category.name, category.color || '#6B7280')
      await loadCategories()
    } catch (err) {
      error.value = err.message || 'Kategori oluşturulamadı'
      console.error('Failed to create category:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function updateCategory(id, category) {
    isLoading.value = true
    error.value = null
    
    try {
      await UpdateCategory(id, category.name, category.color)
      await loadCategories()
    } catch (err) {
      error.value = err.message || 'Kategori güncellenemedi'
      console.error('Failed to update category:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function deleteCategory(id) {
    isLoading.value = true
    error.value = null
    
    try {
      await DeleteCategory(id)
      await loadCategories()
    } catch (err) {
      error.value = err.message || 'Kategori silinemedi'
      console.error('Failed to delete category:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  function clearError() {
    error.value = null
  }

  return {
    categories,
    isLoading,
    error,
    loadCategories,
    createCategory,
    updateCategory,
    deleteCategory,
    clearError
  }
})
