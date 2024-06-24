import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";
import MonitorLayout from "@/layout/monitor-layout.vue";
import serverRouter from "@/router/admin/server.js";
import serverGroupRouter from "@/router/admin/server_group.js";
import notifyGroupRouter from "@/router/admin/notify_group.js";
import notifyLogRouter from "@/router/admin/notify_log.js";
import notifyChannelRouter from "@/router/admin/notify_channel.js";

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            component: MonitorLayout,
            children: [
                {
                    path: '',
                    name: 'dashboard',
                    component: () => import('@/views/admin/dashboard.vue')
                },
                {
                    path: '/monitor',
                    name: 'monitor',
                    component: () => import('@/views/admin/monitor.vue')
                },
                {
                    path: '/setting',
                    name: 'setting',
                    component: () => import('@/views/admin/setting.vue')
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
            component: () => import('@/views/admin/login.vue')
        }
    ]
})

// 拦截器
router.beforeEach((to, from, next) => {
    if (to.name !== 'login' && !localStorage.getItem('token')) {
        next({name: 'login'})
    } else {
        next()
    }
})

export default router
