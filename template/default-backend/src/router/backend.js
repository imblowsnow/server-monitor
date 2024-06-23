import {createRouter, createWebHistory} from "vue-router";
import MonitorLayout from "@/layout/monitor-layout.vue";
import serverRouter from "@/router/server.js";
import serverGroupRouter from "@/router/server_group.js";
import notifyGroupRouter from "@/router/notify_group.js";
import notifyLogRouter from "@/router/notify_log.js";
import notifyChannelRouter from "@/router/notify_channel.js";

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
                {
                    path: '/setting',
                    name: 'setting',
                    component: () => import('@/views/setting.vue')
                },
                serverRouter,
                serverGroupRouter,
                notifyGroupRouter,
                notifyLogRouter,
                notifyChannelRouter
            ]
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/login.vue')
        }
    ]
})


export default router
