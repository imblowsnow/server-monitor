export default {
    path: 'notify_channel',
    children: [
        {
            path: '',
            name: 'notify_channel',
            component: () => import('@/views/notify_channel/index.vue')
        },
        {
            path: 'add',
            name: 'notify_channel_add',
            component: () => import('@/views/notify_channel/add.vue')
        },
        {
            path: 'edit/:id',
            name: 'notify_channel_edit',
            component: () => import('@/views/notify_channel/edit.vue')
        }
    ]
}
