import { defineStore } from 'pinia'
import { 
  GetAllProducts, 
  CreateProduct, 
  UpdateProduct, 
  DeleteProduct,
  GetLowStockProducts
} from '../../wailsjs/go/app/App'

export const useProductStore = defineStore('products', {
  state: () => ({
    products: [],
    loading: false,
    error: null,
    searchQuery: '',
    selectedCategory: '',
    stockFilter: 'all' // 'all', 'low', 'out'
  }),

  getters: {
    filteredProducts: (state) => {
      let filtered = [...state.products]

      // Search filter
      if (state.searchQuery) {
        const query = state.searchQuery.toLowerCase()
        filtered = filtered.filter(p => 
          p.name.toLowerCase().includes(query) || 
          p.code.toLowerCase().includes(query)
        )
      }

      // Category filter
      if (state.selectedCategory) {
        filtered = filtered.filter(p => p.category_id === state.selectedCategory)
      }

      // Stock filter
      if (state.stockFilter === 'low') {
        filtered = filtered.filter(p => p.current_stock <= p.critical_limit && p.current_stock > 0)
      } else if (state.stockFilter === 'out') {
        filtered = filtered.filter(p => p.current_stock === 0)
      }

      return filtered
    },

    lowStockProducts: (state) => {
      return state.products.filter(p => 
        p.current_stock <= p.critical_limit && p.current_stock > 0
      )
    },

    outOfStockProducts: (state) => {
      return state.products.filter(p => p.current_stock === 0)
    }
  },

  actions: {
    async loadProducts() {
      this.loading = true
      this.error = null
      try {
        this.products = await GetAllProducts()
      } catch (err) {
        this.error = err.message || 'Failed to load products'
        console.error('Error loading products:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async createProduct(productData) {
      this.loading = true
      this.error = null
      try {
        const newProduct = await CreateProduct(productData)
        this.products.push(newProduct)
        return newProduct
      } catch (err) {
        this.error = err.message || 'Failed to create product'
        console.error('Error creating product:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async updateProduct(id, productData) {
      this.loading = true
      this.error = null
      try {
        const updatedProduct = await UpdateProduct(id, productData)
        const index = this.products.findIndex(p => p.id === id)
        if (index !== -1) {
          this.products[index] = updatedProduct
        }
        return updatedProduct
      } catch (err) {
        this.error = err.message || 'Failed to update product'
        console.error('Error updating product:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async deleteProduct(id) {
      this.loading = true
      this.error = null
      try {
        await DeleteProduct(id)
        this.products = this.products.filter(p => p.id !== id)
      } catch (err) {
        this.error = err.message || 'Failed to delete product'
        console.error('Error deleting product:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async loadLowStockProducts() {
      this.loading = true
      this.error = null
      try {
        return await GetLowStockProducts()
      } catch (err) {
        this.error = err.message || 'Failed to load low stock products'
        console.error('Error loading low stock products:', err)
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
