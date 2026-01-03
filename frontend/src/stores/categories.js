import { defineStore } from 'pinia'
import { 
  GetAllCategories, 
  CreateCategory, 
  UpdateCategory, 
  DeleteCategory 
} from '../../wailsjs/go/app/App'

export const useCategoryStore = defineStore('categories', {
  state: () => ({
    categories: [],
    loading: false,
    error: null
  }),

  actions: {
    async loadCategories() {
      this.loading = true
      this.error = null
      try {
        this.categories = await GetAllCategories()
      } catch (err) {
        this.error = err.message || 'Failed to load categories'
        console.error('Error loading categories:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async createCategory(categoryData) {
      this.loading = true
      this.error = null
      try {
        const newCategory = await CreateCategory(categoryData)
        this.categories.push(newCategory)
        return newCategory
      } catch (err) {
        this.error = err.message || 'Failed to create category'
        console.error('Error creating category:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async updateCategory(id, categoryData) {
      this.loading = true
      this.error = null
      try {
        const updatedCategory = await UpdateCategory(id, categoryData)
        const index = this.categories.findIndex(c => c.id === id)
        if (index !== -1) {
          this.categories[index] = updatedCategory
        }
        return updatedCategory
      } catch (err) {
        this.error = err.message || 'Failed to update category'
        console.error('Error updating category:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async deleteCategory(id) {
      this.loading = true
      this.error = null
      try {
        await DeleteCategory(id)
        this.categories = this.categories.filter(c => c.id !== id)
      } catch (err) {
        this.error = err.message || 'Failed to delete category'
        console.error('Error deleting category:', err)
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
