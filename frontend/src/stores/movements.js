import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { 
  GetAllMovements, 
  GetMovementsByProduct, 
  GetMovementsByType, 
  CreateMovement
} from '@wails/services/MovementService'

export const useMovementStore = defineStore('movements', () => {
  const movements = ref([])
  const isLoading = ref(false)
  const error = ref(null)
  const filterType = ref(null) // 'in', 'out', or null for all
  const filterProductId = ref(null)

  // Computed
  const filteredMovements = computed(() => {
    let result = movements.value

    if (filterType.value) {
      result = result.filter(m => m.type === filterType.value)
    }

    if (filterProductId.value) {
      result = result.filter(m => m.product_id === filterProductId.value)
    }

    return result
  })

  const totalIn = computed(() => {
    return movements.value
      .filter(m => m.type === 'in')
      .reduce((sum, m) => sum + m.quantity, 0)
  })

  const totalOut = computed(() => {
    return movements.value
      .filter(m => m.type === 'out')
      .reduce((sum, m) => sum + m.quantity, 0)
  })

  const recentMovements = computed(() => {
    return [...movements.value]
      .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
      .slice(0, 10)
  })

  // Actions
  async function loadMovements() {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await GetAllMovements()
      movements.value = result || []
    } catch (err) {
      error.value = err.message || 'Hareketler yüklenemedi'
      console.error('Failed to load movements:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function loadMovementsByProduct(productId) {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await GetMovementsByProduct(productId)
      movements.value = result || []
      filterProductId.value = productId
    } catch (err) {
      error.value = err.message || 'Ürün hareketleri yüklenemedi'
      console.error('Failed to load movements by product:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function loadMovementsByType(type) {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await GetMovementsByType(type)
      movements.value = result || []
      filterType.value = type
    } catch (err) {
      error.value = err.message || 'Hareketler yüklenemedi'
      console.error('Failed to load movements by type:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function createMovement(movement) {
    isLoading.value = true
    error.value = null
    
    try {
      await CreateMovement(
        movement.product_id,
        movement.type, // 'in' or 'out'
        movement.quantity,
        movement.notes || ''
      )
      await loadMovements()
    } catch (err) {
      error.value = err.message || 'Hareket oluşturulamadı'
      console.error('Failed to create movement:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  function setFilterType(type) {
    filterType.value = type
  }

  function setFilterProduct(productId) {
    filterProductId.value = productId
  }

  function clearFilters() {
    filterType.value = null
    filterProductId.value = null
  }

  function clearError() {
    error.value = null
  }

  return {
    movements,
    filteredMovements,
    totalIn,
    totalOut,
    recentMovements,
    isLoading,
    error,
    filterType,
    filterProductId,
    loadMovements,
    loadMovementsByProduct,
    loadMovementsByType,
    createMovement,
    setFilterType,
    setFilterProduct,
    clearFilters,
    clearError
  }
})
