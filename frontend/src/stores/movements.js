import { defineStore } from 'pinia'
import { 
  GetAllMovements, 
  CreateMovement, 
  DeleteMovement,
  GetMovementsByProduct,
  GetMovementStats
} from '../../wailsjs/go/app/App'

export const useMovementStore = defineStore('movements', {
  state: () => ({
    movements: [],
    stats: null,
    loading: false,
    error: null
  }),

  actions: {
    async loadMovements() {
      this.loading = true
      this.error = null
      try {
        this.movements = await GetAllMovements()
      } catch (err) {
        this.error = err.message || 'Failed to load movements'
        console.error('Error loading movements:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async createMovement(movementData) {
      this.loading = true
      this.error = null
      try {
        const newMovement = await CreateMovement(movementData)
        this.movements.push(newMovement)
        return newMovement
      } catch (err) {
        this.error = err.message || 'Failed to create movement'
        console.error('Error creating movement:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async deleteMovement(id) {
      this.loading = true
      this.error = null
      try {
        await DeleteMovement(id)
        this.movements = this.movements.filter(m => m.id !== id)
      } catch (err) {
        this.error = err.message || 'Failed to delete movement'
        console.error('Error deleting movement:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async loadMovementsByProduct(productId) {
      this.loading = true
      this.error = null
      try {
        return await GetMovementsByProduct(productId)
      } catch (err) {
        this.error = err.message || 'Failed to load product movements'
        console.error('Error loading product movements:', err)
        throw err
      } finally {
        this.loading = false
      }
    },

    async loadStats() {
      this.loading = true
      this.error = null
      try {
        this.stats = await GetMovementStats()
        return this.stats
      } catch (err) {
        this.error = err.message || 'Failed to load movement stats'
        console.error('Error loading movement stats:', err)
        throw err
      } finally {
        this.loading = false
      }
    }
  }
})
