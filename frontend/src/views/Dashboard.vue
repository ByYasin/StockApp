<template>
  <div class="h-screen flex flex-col bg-gray-50">
    <!-- Header / Navigation -->
    <header class="bg-white shadow-sm border-b border-gray-200">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4">
            <h1 class="text-2xl font-bold text-gray-800">Stok Takip Sistemi</h1>
            <span class="px-3 py-1 text-sm bg-blue-100 text-blue-800 rounded-full">
              {{ dbStore.activeDatabaseName }}
            </span>
          </div>
          
          <nav class="flex space-x-1">
            <router-link
              v-for="item in navItems"
              :key="item.name"
              :to="{ name: item.route }"
              class="px-4 py-2 rounded-lg text-sm font-medium transition-colors"
              :class="currentRoute === item.route 
                ? 'bg-blue-600 text-white' 
                : 'text-gray-600 hover:bg-gray-100'"
            >
              {{ item.label }}
            </router-link>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-1 overflow-auto p-6">
      <div class="max-w-7xl mx-auto">
        <!-- Stats Cards -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
          <div class="card">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Toplam Ürün</p>
                <p class="text-3xl font-bold text-gray-800">0</p>
              </div>
              <div class="p-3 bg-blue-100 rounded-full">
                <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Kritik Stok</p>
                <p class="text-3xl font-bold text-red-600">0</p>
              </div>
              <div class="p-3 bg-red-100 rounded-full">
                <svg class="w-8 h-8 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Kategoriler</p>
                <p class="text-3xl font-bold text-gray-800">0</p>
              </div>
              <div class="p-3 bg-green-100 rounded-full">
                <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">Bugünkü Hareketler</p>
                <p class="text-3xl font-bold text-gray-800">0</p>
              </div>
              <div class="p-3 bg-purple-100 rounded-full">
                <svg class="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>

        <!-- Welcome Section -->
        <div class="card mb-6">
          <h2 class="text-2xl font-bold text-gray-800 mb-4">Hoş Geldiniz!</h2>
          <p class="text-gray-600 mb-6">
            Stok takip sisteminize başarıyla bağlandınız. Sol üstteki menüden istediğiniz sayfaya geçebilirsiniz.
          </p>
          
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <router-link
              :to="{ name: 'Products' }"
              class="p-4 border-2 border-gray-200 rounded-lg hover:border-blue-500 hover:bg-blue-50 transition-all cursor-pointer group"
            >
              <div class="flex items-center space-x-3">
                <div class="p-2 bg-blue-100 rounded-lg group-hover:bg-blue-200 transition-colors">
                  <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-800">Ürünler</p>
                  <p class="text-sm text-gray-500">Ürün yönetimi</p>
                </div>
              </div>
            </router-link>

            <router-link
              :to="{ name: 'Movements' }"
              class="p-4 border-2 border-gray-200 rounded-lg hover:border-green-500 hover:bg-green-50 transition-all cursor-pointer group"
            >
              <div class="flex items-center space-x-3">
                <div class="p-2 bg-green-100 rounded-lg group-hover:bg-green-200 transition-colors">
                  <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-800">Stok Hareketi</p>
                  <p class="text-sm text-gray-500">Giriş/Çıkış işlemleri</p>
                </div>
              </div>
            </router-link>

            <router-link
              :to="{ name: 'Categories' }"
              class="p-4 border-2 border-gray-200 rounded-lg hover:border-purple-500 hover:bg-purple-50 transition-all cursor-pointer group"
            >
              <div class="flex items-center space-x-3">
                <div class="p-2 bg-purple-100 rounded-lg group-hover:bg-purple-200 transition-colors">
                  <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                  </svg>
                </div>
                <div>
                  <p class="font-medium text-gray-800">Kategoriler</p>
                  <p class="text-sm text-gray-500">Kategori yönetimi</p>
                </div>
              </div>
            </router-link>
          </div>
        </div>

        <!-- Recent Activities Placeholder -->
        <div class="card">
          <h3 class="text-lg font-semibold mb-4 text-gray-800">Son Hareketler</h3>
          <div class="text-center py-8 text-gray-500">
            <p>Henüz hareket kaydı yok</p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useDatabaseStore } from '@/stores/database'

const route = useRoute()
const dbStore = useDatabaseStore()

const currentRoute = computed(() => route.name)

const navItems = [
  { label: 'Ana Sayfa', route: 'Dashboard' },
  { label: 'Ürünler', route: 'Products' },
  { label: 'Kategoriler', route: 'Categories' },
  { label: 'Hareketler', route: 'Movements' }
]
</script>
