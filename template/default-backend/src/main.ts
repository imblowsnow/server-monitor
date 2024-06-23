import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import mitt from 'mitt'

import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

const app = createApp(App)

app.use(createPinia())

app.use(router)
app.use(ElementPlus)

// 注册 全局事件总线
app.config.globalProperties.$bus = mitt()

app.mount('#app')
