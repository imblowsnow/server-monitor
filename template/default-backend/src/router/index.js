import {createRouter, createWebHistory} from 'vue-router'
import MonitorLayout from "@/layout/monitor-layout.vue";
import serverRouter from './server.js'
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            component: MonitorLayout,
            children: [
                {
                    path: '',
                    name: 'dashboard',
                    component: () => import('@/views/dashboard.vue')
                },
                {
                    path: '/monitor',
                    name: 'monitor',
                    component: () => import('@/views/monitor.vue')
                },
                serverRouter
            ]
        }
    ]
})

export default router
