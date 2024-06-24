import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('@/views/index.vue')
        },
        {
            path: '/server/:id',
            name: 'server',
            component: () => import('@/views/server.vue')
        }
    ]
})
export default router
