<template>
  <div class="h-screen flex flex-col bg-gray-50">
    <header class="bg-white shadow-sm border-b border-gray-200">
      <div class="px-6 py-4">
        <h1 class="text-2xl font-bold text-gray-800">Kategori Yönetimi</h1>
      </div>
    </header>
    
    <main class="flex-1 overflow-auto p-6">
      <div class="max-w-7xl mx-auto">
        <div class="card text-center py-12">
          <p class="text-gray-600">Kategori yönetimi sayfası - Yakında gelecek</p>
          <router-link :to="{ name: 'Dashboard' }" class="text-blue-600 hover:underline mt-4 inline-block">
            Ana sayfaya dön
          </router-link>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
// TODO: Implement category management
</script>
