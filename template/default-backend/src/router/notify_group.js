export default {
    path: 'notify_group',
    children: [
        {
            path: '',
            name: 'notify_group',
            component: () => import('@/views/notify_group/index.vue')
        },
        {
            path: 'add',
            name: 'notify_group_add',
            component: () => import('@/views/notify_group/add.vue')
        },
        {
            path: 'edit/:id',
            name: 'notify_group_edit',
            component: () => import('@/views/notify_group/edit.vue')
        }
    ]
}
