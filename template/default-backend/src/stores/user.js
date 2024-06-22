import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({ token: localStorage.getItem('token') }),
  getters: {

  },
  actions: {
    updateToken(token) {
        this.token = token
    },
    isLogin() {
        return this.token !== null
    },
    setToken(token) {
      this.token = token
      localStorage.setItem('token', token)
    },
    logout() {
      this.token = null
      localStorage.setItem('token', '')
    }
  },
})
