export default {
    path: 'notify_log',
    children: [
        {
            path: '',
            name: 'notify_log',
            component: () => import('@/views/admin/notify_log/index.vue')
        }
    ]
}
