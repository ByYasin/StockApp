<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 dark:from-gray-900 dark:to-gray-800">
    <div class="w-full max-w-2xl p-8 mx-4">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-xl p-8">
        <div class="text-center mb-8">
          <h1 class="text-4xl font-bold text-gray-800 dark:text-gray-100 mb-2">Stok Takip Sistemi</h1>
          <p class="text-gray-600 dark:text-gray-400">Veritabanı seçin veya yeni bir tane oluşturun</p>
        </div>

        <div v-if="!isLoading && databases.length > 0" class="mb-6">
          <h2 class="text-lg font-semibold mb-3 text-gray-700 dark:text-gray-300">Mevcut Veritabanları</h2>
          <div class="space-y-2">
            <div
              v-for="db in databases"
              :key="db.path"
              @click="handleSelectDatabase(db.path)"
              class="p-4 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-blue-50 dark:hover:bg-gray-700 hover:border-blue-300 dark:hover:border-blue-500 cursor-pointer transition-all duration-200 group"
            >
              <div class="flex justify-between items-center">
                <div>
                  <p class="font-medium text-gray-800 dark:text-gray-200 group-hover:text-blue-600 dark:group-hover:text-blue-400">{{ db.name }}</p>
                  <p class="text-sm text-gray-500 dark:text-gray-400">{{ db.modified }}</p>
                </div>
                <div class="text-right">
                  <p class="text-sm text-gray-600 dark:text-gray-300">{{ db.size.toFixed(2) }} MB</p>
                  <span class="text-xs text-green-600 dark:text-green-400" v-if="db.is_active">● Aktif</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="!isLoading && databases.length === 0" class="text-center py-8 mb-6">
          <svg class="mx-auto h-16 w-16 text-gray-400 dark:text-gray-500 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
          </svg>
          <p class="text-gray-600 dark:text-gray-400">Henüz hiç veritabanı yok. Yeni bir tane oluşturun!</p>
        </div>

        <div v-if="isLoading" class="text-center py-8 mb-6">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 dark:border-blue-400"></div>
          <p class="mt-2 text-gray-600 dark:text-gray-400">Yükleniyor...</p>
        </div>

        <div v-if="error" class="mb-6 p-4 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg">
          <p class="text-red-800 dark:text-red-200 text-sm">{{ error }}</p>
        </div>

        <button
          @click="showCreateDialog = true"
          class="w-full btn btn-primary py-3 text-lg"
          :disabled="isLoading"
        >
          <span class="flex items-center justify-center">
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
            </svg>
            Yeni Veritabanı Oluştur
          </span>
        </button>
      </div>
    </div>

    <div v-if="showCreateDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 w-full max-w-md mx-4">
        <h3 class="text-xl font-bold mb-4 text-gray-800 dark:text-gray-100">Yeni Veritabanı Oluştur</h3>
        
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Veritabanı Adı</label>
          <input
            v-model="newDbName"
            type="text"
            placeholder="Örn: Depo_A"
            class="input bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100 border-gray-300 dark:border-gray-600"
            @keyup.enter="handleCreateDatabase"
          />
        </div>

        <div class="flex gap-3">
          <button
            @click="handleCreateDatabase"
            class="flex-1 btn btn-primary"
            :disabled="!newDbName.trim() || isCreating"
          >
            {{ isCreating ? 'Oluşturuluyor...' : 'Oluştur' }}
          </button>
          <button
            @click="showCreateDialog = false"
            class="flex-1 btn btn-secondary"
            :disabled="isCreating"
          >
            İptal
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useDatabaseStore } from '@/stores/database'
import { storeToRefs } from 'pinia'

const router = useRouter()
const dbStore = useDatabaseStore()
const { databases, isLoading, error } = storeToRefs(dbStore)

const showCreateDialog = ref(false)
const newDbName = ref('')
const isCreating = ref(false)

onMounted(async () => {
  await dbStore.loadDatabases()
})

async function handleSelectDatabase(dbPath) {
  try {
    await dbStore.selectDatabase(dbPath)
    router.push({ name: 'Dashboard' })
  } catch (err) {
    console.error('Failed to select database:', err)
  }
}

async function handleCreateDatabase() {
  if (!newDbName.value.trim()) return
  
  isCreating.value = true
  dbStore.clearError()
  
  try {
    await dbStore.createDatabase(newDbName.value)
    showCreateDialog.value = false
    newDbName.value = ''
    
    if (databases.value.length > 0) {
      const newDb = databases.value[databases.value.length - 1]
      await handleSelectDatabase(newDb.path)
    }
  } catch (err) {
    console.error('Failed to create database:', err)
    alert(`Veritabanı oluşturulamadı: ${error.value || err.message}`)
  } finally {
    isCreating.value = false
  }
}
</script>
