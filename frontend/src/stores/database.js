import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { ListDatabases, CreateDatabase, SwitchDatabase, IsConnected } from '../../wailsjs/go/app/App'

export const useDatabaseStore = defineStore('database', () => {
  // State
  const databases = ref([])
  const currentDatabase = ref(null)
  const isConnected = ref(false)
  const isLoading = ref(false)
  const error = ref(null)

  // Getters
  const activeDatabaseName = computed(() => {
    return currentDatabase.value?.name || 'Veritabanı Seçilmedi'
  })

  // Actions
  async function loadDatabases() {
    isLoading.value = true
    error.value = null
    
    try {
      const result = await ListDatabases()
      databases.value = result || []
      
      // Check connection status
      isConnected.value = await IsConnected()
    } catch (err) {
      error.value = err.message || 'Veritabanları yüklenemedi'
      console.error('Failed to load databases:', err)
      databases.value = []
    } finally {
      isLoading.value = false
    }
  }

  async function selectDatabase(dbPath) {
    isLoading.value = true
    error.value = null
    
    try {
      await SwitchDatabase(dbPath)
      currentDatabase.value = databases.value.find(db => db.path === dbPath)
      isConnected.value = true
    } catch (err) {
      error.value = err.message || 'Veritabanı seçilemedi'
      console.error('Failed to select database:', err)
      isConnected.value = false
    } finally {
      isLoading.value = false
    }
  }

  async function createDatabase(name) {
      isLoading.value = true
      error.value = null
      
      try {
        console.log('[Store] Creating database:', name)
        await CreateDatabase(name)
        console.log('[Store] Database created successfully')
        await loadDatabases()
      } catch (err) {
        // Get detailed error message
        const errorMsg = err?.message || err?.toString() || 'Bilinmeyen hata'
        error.value = `Veritabanı oluşturulamadı: ${errorMsg}`
        console.error('[Store] Failed to create database:', {
          name,
          error: err,
          message: errorMsg,
          stack: err?.stack
        })
        throw err
      } finally {
        isLoading.value = false
      }
    }

  async function backupDatabase() {
    try {
      // TODO: Implement backup in backend
      return ''
    } catch (err) {
      error.value = err.message
      console.error('Failed to backup database:', err)
      throw err
    }
  }

  function clearError() {
    error.value = null
  }

  function disconnect() {
    // Veritabanı bağlantısını kes
    currentDatabase.value = null
    isConnected.value = false
    error.value = null
  }

  return {
    // State
    databases,
    currentDatabase,
    isConnected,
    isLoading,
    error,
    // Getters
    activeDatabaseName,
    // Actions
    loadDatabases,
    selectDatabase,
    createDatabase,
    backupDatabase,
    clearError,
    disconnect
  }
})
