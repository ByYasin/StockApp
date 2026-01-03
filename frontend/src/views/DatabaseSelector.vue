<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100">
    <div class="w-full max-w-2xl p-8 mx-4">
      <div class="bg-white rounded-lg shadow-xl p-8">
        <!-- Header -->
        <div class="text-center mb-8">
          <h1 class="text-4xl font-bold text-gray-800 mb-2">Stok Takip Sistemi</h1>
          <p class="text-gray-600">Veritabanı seçin veya yeni bir tane oluşturun</p>
        </div>

        <!-- Database List -->
        <div v-if="!isLoading && databases.length > 0" class="mb-6">
          <h2 class="text-lg font-semibold mb-3 text-gray-700">Mevcut Veritabanları</h2>
          <div class="space-y-2">
            <div
              v-for="db in databases"
              :key="db.path"
              @click="handleSelectDatabase(db.path)"
              class="p-4 border border-gray-200 rounded-lg hover:bg-blue-50 hover:border-blue-300 cursor-pointer transition-all duration-200 group"
            >
              <div class="flex justify-between items-center">
                <div>
                  <p class="font-medium text-gray-800 group-hover:text-blue-600">{{ db.name }}</p>
                  <p class="text-sm text-gray-500">{{ db.modified }}</p>
                </div>
                <div class="text-right">
                  <p class="text-sm text-gray-600">{{ db.size.toFixed(2) }} MB</p>
                  <span class="text-xs text-green-600" v-if="db.is_active">● Aktif</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-if="!isLoading && databases.length === 0" class="text-center py-8 mb-6">
          <svg class="mx-auto h-16 w-16 text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
          </svg>
          <p class="text-gray-600">Henüz hiç veritabanı yok. Yeni bir tane oluşturun!</p>
        </div>

        <!-- Loading State -->
        <div v-if="isLoading" class="text-center py-8 mb-6">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
          <p class="mt-2 text-gray-600">Yükleniyor...</p>
        </div>

        <!-- Error Message -->
        <div v-if="error" class="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg">
          <p class="text-red-800 text-sm">{{ error }}</p>
        </div>

        <!-- Create New Database Button -->
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

    <!-- Create Database Dialog -->
    <div v-if="showCreateDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-md mx-4">
        <h3 class="text-xl font-bold mb-4">Yeni Veritabanı Oluştur</h3>
        
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-2">Veritabanı Adı</label>
          <input
            v-model="newDbName"
            type="text"
            placeholder="Örn: Depo_A"
            class="input"
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
  dbStore.clearError() // Clear previous errors
  
  try {
    await dbStore.createDatabase(newDbName.value)
    showCreateDialog.value = false
    newDbName.value = ''
    
    // Auto-select the newly created database
    if (databases.value.length > 0) {
      const newDb = databases.value[databases.value.length - 1]
      await handleSelectDatabase(newDb.path)
    }
  } catch (err) {
    console.error('Failed to create database:', err)
    // Error is already set in store, just log it
    alert(`Veritabanı oluşturulamadı: ${error.value || err.message}`)
  } finally {
    isCreating.value = false
  }
}
</script>
