import { defineStore } from 'pinia'
import { GetTheme, SetTheme } from '../../wailsjs/go/app/App'

export const useThemeStore = defineStore('theme', {
  state: () => ({
    currentTheme: 'light',
    isDark: false
  }),

  actions: {
    async loadTheme() {
      try {
        this.currentTheme = await GetTheme()
        this.isDark = this.currentTheme === 'dark'
        this.applyTheme()
      } catch (err) {
        console.error('Failed to load theme:', err)
      }
    },

    async toggleTheme() {
      this.isDark = !this.isDark
      this.currentTheme = this.isDark ? 'dark' : 'light'
      
      try {
        await SetTheme(this.currentTheme)
        this.applyTheme()
      } catch (err) {
        console.error('Failed to save theme:', err)
      }
    },

    applyTheme() {
      if (this.isDark) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    }
  }
})
