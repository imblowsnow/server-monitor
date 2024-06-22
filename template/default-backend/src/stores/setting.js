import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', {
  state: () => ({ setting: {} }),
  getters: {

  },
  actions: {
    updateSetting(data) {
      this.setting = data
      // 修改标题
      document.title = data.site_name;
    },
  },
})
