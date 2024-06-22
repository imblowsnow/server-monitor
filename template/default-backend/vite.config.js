import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        VueDevTools(),
        AutoImport({
            resolvers: [ElementPlusResolver()],
        }),
        Components({
            resolvers: [ElementPlusResolver()],
        }),
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
    // 修改编译后生成的目录
    build: {
        outDir: '../../pkg/server/template/backend/default',
    },
    // 设置代理
    server: {
        proxy: {
            '/admin-api': {
                target: 'http://127.0.0.1:22251',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/admin-api/, '/admin-api')
            },
            '/ws': {
                target: 'ws://127.0.0.1:22251',
                changeOrigin: true,
                ws: true,
                rewrite: (path) => path.replace(/^\/ws/, '/ws')
            }
        }
    }
})
